## About GGIpkg
GGIpkg(GoGetInfo Package) is a package for Golang language that allows you to use it to get the information of your desired site

## Install and Run
```bash
go mod init theFr3Y
go get -u https://github.com/theFr3Y/GGIpkg.git
````

## How to Use
```
package main

import (
  "fmt"
  "github.com/theFr3Y/GGIpkg"
)

func main() {
  GGIpkg.All("google.com")
}
````

## All of methods:
- All
-Country
-Org
-CountryCode
-Ration
-RationName
-Isp
-As
-City
-TimeZone
-Query
-Zip
-Lat
-Lon
-Status
