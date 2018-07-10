package main

import (
	"flag"
	"fmt"

	"github.com/binaryfigments/ipinfo"
)

func main() {

	var ip string
	var geo bool

	flag.StringVar(&ip, "ip", "84.26.250.163", "ip address")
	flag.BoolVar(&geo, "geo", true, "ip address")
	flag.Parse()

	data := ipinfo.Get(ip, geo)
	/*
		json, err := json.MarshalIndent(data, "", "  ")
		if err != nil {
			fmt.Println(err)
		}
		fmt.Printf("%s\n", json)
	*/
	fmt.Printf("# Information on %s\n", data.IPAddress)
	fmt.Printf("\n")
	fmt.Printf("     IP address : %s\n", data.IPAddress)
	fmt.Printf("     IP decimal : %v\n", data.IPDecimal)
	fmt.Printf("     PTR fir IP : %s\n", data.PTR)
	fmt.Printf("\n")
	fmt.Printf("## Geo information\n")
	fmt.Printf("\n")
	fmt.Printf("           City : %s\n", data.GeoInfo.City)
	fmt.Printf("    Subdivision : %s\n", data.GeoInfo.Subdivision)
	fmt.Printf("        Country : %s\n", data.GeoInfo.Country)
	fmt.Printf("   Country Code : %s\n", data.GeoInfo.CountryCode)
	fmt.Printf("       TimeZone : %s\n", data.GeoInfo.TimeZone)
	fmt.Printf("       Latitude : %v\n", data.GeoInfo.Latitude)
	fmt.Printf("      Longitude : %v\n", data.GeoInfo.Longitude)
	fmt.Printf("Accuracy Radius : %v\n", data.GeoInfo.AccuracyRadius)
	fmt.Printf("\n")
	fmt.Printf("## AS information\n")
	fmt.Printf("\n")
	fmt.Printf("      AS Number : %v\n", data.ASInfo.ASNumber)
	fmt.Printf("   Organization : %v\n", data.ASInfo.ASOrganization)
	fmt.Printf("\n")
	fmt.Printf("## WHOIS\n")
	fmt.Printf("\n")
	fmt.Printf("%s\n", data.WhoisRAW)
}
