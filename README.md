# ipinfo

Get information about an IP address like the PTR record, geo information, AS-number and WHOIS.

## Work in progress

This is a "work in progress" project at the moment. Things can chang.

## GeoLite2

Get the `GeoLite2-ASN.mmdb` and `GeoLite2-City.mmdb` from [MaxMind](https://dev.maxmind.com/geoip/geoip2/geolite2/) and place them in the directory of your executable. Path's can change in the future.

Make sure that you give [MaxMind](https://www.maxmind.com) the credits of the information in the databases.

```html
This product includes GeoLite2 data created by MaxMind, available from
<a href="http://www.maxmind.com">http://www.maxmind.com</a>.
```

## Used libs

I used [oschwald's](https://github.com/oschwald) [geoip2-golang](https://github.com/oschwald/geoip2-golang) library to read the MaxMind databases.

```bash
go get github.com/oschwald/geoip2-golang
```

## Example CLI

Output of the example CLI.

```bash
$ ./cli -ip=1.1.1.1 -geo=true
# Information on 1.1.1.1

     IP address : 1.1.1.1
     IP decimal : 16843009
     PTR fir IP : 1dot1dot1dot1.cloudflare-dns.com

## Geo information

           City : Research
    Subdivision : Victoria
        Country : Australia
   Country Code : AU
       TimeZone : Australia/Melbourne
       Latitude : -37.7
      Longitude : 145.1833
Accuracy Radius : 1000

## AS information

      AS Number : 13335
   Organization : Cloudflare Inc

## WHOIS

% This is the RIPE Database query service.
% The objects are in RPSL format.
%
% The RIPE Database is subject to Terms and Conditions.
% See http://www.ripe.net/db/support/db-terms-conditions.pdf

% Information related to '1.1.1.0 - 1.1.1.255'

inetnum:        1.1.1.0 - 1.1.1.255
netname:        APNIC-LABS
descr:          APNIC and Cloudflare DNS Resolver project
descr:          Routed globally by AS13335/Cloudflare
descr:          Research prefix for APNIC Labs
country:        AU
org:            ORG-ARAD1-AP
admin-c:        DUMY-RIPE
tech-c:         DUMY-RIPE
mnt-by:         APNIC-HM
mnt-routes:     MAINT-AU-APNIC-GM85-AP
mnt-irt:        IRT-APNICRANDNET-AU
status:         ASSIGNED PORTABLE
remarks:        ---------------
remarks:        All Cloudflare abuse reporting can be done via
remarks:        resolver-abuse@cloudflare.com
remarks:        ---------------
last-modified:  2018-03-30T01:51:28Z
source:         APNIC-GRS
remarks:        ****************************
remarks:        * THIS OBJECT IS MODIFIED
remarks:        * Please note that all data that is generally regarded as personal
remarks:        * data has been removed from this object.
remarks:        * To view the original object, please query the APNIC Database at:
remarks:        * http://www.apnic.net/
remarks:        ****************************

% Information related to '1.1.1.0/24AS13335'

route:          1.1.1.0/24
origin:         AS13335
descr:          APNIC Research and Development
                6 Cordelia St
mnt-by:         MAINT-AU-APNIC-GM85-AP
last-modified:  2018-03-16T16:58:06Z
source:         APNIC-GRS
remarks:        ****************************
remarks:        * THIS OBJECT IS MODIFIED
remarks:        * Please note that all data that is generally regarded as personal
remarks:        * data has been removed from this object.
remarks:        * To view the original object, please query the APNIC Database at:
remarks:        * http://www.apnic.net/
remarks:        ****************************

% This query was served by the RIPE Database Query Service version 1.91.2 (BLAARKOP)
```
