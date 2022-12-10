package GGIpkg

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func All(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	for i := range result {
		fmt.Printf("%s", result[i])
	}
}

func Country(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["country"])
}

func Org(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["org"])
}

func CountryCode(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["countryCode"])
}

func Ration(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["ragionName"])
}

func Isp(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["isp"])
}

func As(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["as"])
}

func Region(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["ragion"])
}

func City(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["city"])
}

func TimeZone(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["timezone"])
}

func Query(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["query"])
}

func Zip(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["zip"])
}

func Lat(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["lat"])
}

func Lon(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["lon"])
}

func Status(addr string) {
	resp, _ := http.Get("http://ip-api.com/json/" + addr)

	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	result := map[string]string{}
	json.Unmarshal(body, &result)
	fmt.Printf("%s", result["status"])
}
