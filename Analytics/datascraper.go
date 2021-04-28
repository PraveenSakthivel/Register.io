package main

import (
	"fmt"
	"io/ioutil"
	"main/models"
	"net/http"

	"github.com/tidwall/gjson"
)

// ScrapeData => Collect place data and write to DB
func ScrapeData() {
	models.ConnectDB()

	url := "https://rumobile.rutgers.edu/2/places.txt"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	jsonString := string(body)
	buildings := gjson.Get(jsonString, "all")
	buildings.ForEach(func(key gjson.Result, value gjson.Result) bool {
		if value.Get("building_code").String() != "" {
			entry := models.BuildingLookup{}
			entry.BldgNum = value.Get("building_number").String()
			entry.Code = value.Get("building_code").String()
			entry.Name = value.Get("title").String()
			entry.Lat = value.Get("location.latitude").String()
			entry.Long = value.Get("location.longitude").String()
			models.DB.Create(&entry)
		}
		return true
	})
}
