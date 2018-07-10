package ipinfo

import (
	"io/ioutil"
	"math/big"
	"net"
	"strings"
	"time"

	"github.com/oschwald/geoip2-golang"
)

// Data struct is the main struct
type Data struct {
	IPAddress    string     `json:"ipaddress"`
	IPDecimal    uint64     `json:"ipdecimal,omitempty"`
	PTR          string     `json:"ptr,omitempty"`
	GeoInfo      *geoip     `json:"geo,omitempty"`
	ASInfo       *asnip     `json:"asn,omitempty"`
	WhoisRAW     string     `json:"whois_raw,omitempty"`
	Controls     []*control `json:"control,omitempty"`
	Error        bool       `json:"error,omitempty"`
	ErrorMessage string     `json:"errormessage,omitempty"`
}

type asnip struct {
	ASNumber       uint   `json:"as_number,omitempty"`
	ASOrganization string `json:"as_organization,omitempty"`
}

type geoip struct {
	Country        string  `json:"country,omitempty"`
	Subdivision    string  `json:"subdivision,omitempty"`
	CountryCode    string  `json:"country_code,omitempty"`
	City           string  `json:"city,omitempty"`
	TimeZone       string  `json:"time_zone,omitempty"`
	Latitude       float64 `json:"latitude,omitempty"`
	Longitude      float64 `json:"longitude,omitempty"`
	AccuracyRadius uint16  `json:"accuracy_radius,omitempty"`
}

type control struct {
	Message  string `json:"message,omitempty"`
	Blocking bool   `json:"blocking,omitempty"`
}

// Version v0.0.1

// Get function, main function of this module.
func Get(ipaddress string, geo bool) *Data {
	response := new(Data)
	response.IPAddress = ipaddress

	ip := net.ParseIP(ipaddress)
	if ip == nil {
		response.Error = true
		response.ErrorMessage = "not a valid ip address"
		return response
	}
	ptrrecord, err := lookupAddr(ip)
	if err != nil {
		contrl := new(control)
		contrl.Message = "No PTR record found for this IP."
		contrl.Blocking = false
		response.Controls = append(response.Controls, contrl)
	}
	response.PTR = ptrrecord

	if geo == true {
		contrl := new(control)
		contrl.Message = "This product includes GeoLite2 data created by MaxMind, available from http://www.maxmind.com."
		contrl.Blocking = false
		response.Controls = append(response.Controls, contrl)
		geodata, err := getGeoInfo(ipaddress)
		if err != nil {
			contrl := new(control)
			contrl.Message = "GeoLocation lookup failed."
			contrl.Blocking = false
			response.Controls = append(response.Controls, contrl)
		}
		response.GeoInfo = geodata

		asndata, err := getASNInfo(ipaddress)
		if err != nil {
			contrl := new(control)
			contrl.Message = "AS lookup failed."
			contrl.Blocking = false
			response.Controls = append(response.Controls, contrl)
		}
		response.ASInfo = asndata
	}

	response.IPDecimal = toDecimal(ip)

	// whois -h whois.ripe.net -B --resource 1.1.1.1
	whois, err := getWhois("-B --resource "+ip.String(), "whois.ripe.net")
	response.WhoisRAW = whois
	// fmt.Println(whois)

	return response
}

func lookupAddr(ip net.IP) (string, error) {
	names, err := net.LookupAddr(ip.String())
	if err != nil || len(names) == 0 {
		return "", err
	}
	return strings.TrimRight(names[0], "."), nil
}

func toDecimal(ip net.IP) uint64 {
	i := big.NewInt(0)
	if to4 := ip.To4(); to4 != nil {
		i.SetBytes(to4)
	} else {
		i.SetBytes(ip)
	}
	return i.Uint64()
}

// --------------------

func getGeoInfo(ipaddress string) (*geoip, error) {
	response := new(geoip)
	// db, err := geoip2.Open("GeoIP2-City.mmdb")
	db, err := geoip2.Open("GeoLite2-City.mmdb")
	if err != nil {
		return response, err
	}
	defer db.Close()

	ip := net.ParseIP(ipaddress)
	record, err := db.City(ip)
	if err != nil {
		return response, err
	}

	response.City = record.City.Names["en"]
	if len(record.Subdivisions) > 0 {
		response.Subdivision = record.Subdivisions[0].Names["en"]
	}
	response.CountryCode = record.Country.IsoCode
	response.Country = record.Country.Names["en"]
	response.TimeZone = record.Location.TimeZone
	response.Latitude = record.Location.Latitude
	response.Longitude = record.Location.Longitude
	response.AccuracyRadius = record.Location.AccuracyRadius

	return response, nil
}

// --------------------

func getASNInfo(ipaddress string) (*asnip, error) {
	response := new(asnip)
	db, err := geoip2.Open("GeoLite2-ASN.mmdb")
	if err != nil {
		return response, err
	}
	defer db.Close()

	ip := net.ParseIP(ipaddress)
	record, err := db.ASN(ip)
	if err != nil {
		return response, err
	}

	response.ASNumber = record.AutonomousSystemNumber
	response.ASOrganization = record.AutonomousSystemOrganization

	return response, nil
}

// ---------------------------

// whois -h whois.ripe.net -B --resource 1.1.1.1
func getWhois(search string, host string) (s string, err error) {

	conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, "43"), time.Second*28)
	defer conn.Close()

	if err != nil {
		return s, err
	}

	conn.Write([]byte(search + "\r\n"))

	buffer, err := ioutil.ReadAll(conn)

	if err != nil {
		return s, err
	}

	s = string(buffer[:])

	return s, err
}
