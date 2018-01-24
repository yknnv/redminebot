package main

import (
    "log"
    "net/http"
    "io/ioutil"
    "encoding/json"
)


func main() {
    client := &http.Client{}
    url := "http://redmine.aeroidea.ru/projects/270.json"
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("Authorization", "Basic ")
    req.Header.Set("Content-Type", "application/json")
    res, _ := client.Do(req)
    defer res.Body.Close()

    data, _ := ioutil.ReadAll(res.Body)
    //var result myXMLstruct
    //xml.Unmarshal(xmlBytes, &result)
    
    decoder := json.NewDecoder(data)
    var configuration string
    err := decoder.Decode(&configuration)

    log.Println(configuration)
    
}
