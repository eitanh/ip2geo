package main

import (
    "fmt"
    "github.com/oschwald/geoip2-golang"
    "log"
    "net"
    "net/http"
    "encoding/json"
)


type whois struct{
         City string
         Country string

}
func listen(){
http.HandleFunc("/", ip2geo)
  if err := http.ListenAndServe(":8080", nil); err != nil {
    log.Fatal(err)
  }

}

func ip2geo(w http.ResponseWriter, r *http.Request){
    queryValues := r.URL.Query()
    //fmt.Printf(string(myip[0]))
    ipParam:=string(queryValues.Get("ip"));
    //fmt.Fprintf(w, "hello, %s!\n", ipParam) 
    db, err := geoip2.Open("GeoLite2-City.mmdb")
    if err != nil {
            log.Fatal(err)
    }
    defer db.Close()
    // If you are using strings that may be invalid, check that ip is not nil
    if (ipParam!=""){
        ip := net.ParseIP(ipParam)
        record, err := db.City(ip)

        if err != nil {
                log.Fatal(err)
        }
        r:=whois{record.City.Names["pt-BR"],record.Country.Names["en"]}
        jdata, err := json.Marshal(r)
        fmt.Printf("%+v\n", r)
        w.Header().Set("Content-Type", "application/json")
        w.Write(jdata)
        //fmt.Printf("Portuguese (BR) city name: %v\n", record.City.Names["pt-BR"])
        //fmt.Printf("English subdivision name: %v\n", record.Subdivisions[0].Names["en"])
        //fmt.Printf("Russian country name: %v\n", record.Country.Names["ru"])
        //fmt.Printf("ISO country code: %v\n", record.Country.IsoCode)
        //fmt.Printf("Time zone: %v\n", record.Location.TimeZone)
        //fmt.Printf("Coordinates: %v, %v\n", record.Location.Latitude, record.Location.Longitude)
    }
}

func main() {
    listen();
}
