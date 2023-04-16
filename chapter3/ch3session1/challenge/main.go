package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
)

var (
	url    = "https://jsonplaceholder.typicode.com/posts"
	method = "POST"
	loop   = true
	wind   = 0
	water  = 0
)

func main() {

	for true {
		randomizer()
		post()
		time.Sleep(15 * time.Second)
	}

}

func randomizer() {
	min := 1
	max := 100

	wind = rand.Intn(max-min) + min
	water = rand.Intn(max-min) + min
}

func post() {

	data := fmt.Sprintf(`{
		"water": %d,
		"wind": %d
	}`, water, wind)

	payload := strings.NewReader(data)

	client := &http.Client{}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("Content-Type", "application/json")

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
	fmt.Println(string(body))
	statusWater()
	statusWind()
}

func statusWater() {
	var waterStatus string

	if water < 5 {
		waterStatus = "aman"
	} else if water >= 5 && water <= 8 {
		waterStatus = "siaga"
	} else {
		waterStatus = "bahaya"
	}

	fmt.Printf("status water : %s\n", waterStatus)
}

func statusWind() {
	var windStatus string

	if wind < 6 {
		windStatus = "aman"
	} else if water >= 7 && water <= 15 {
		windStatus = "siaga"
	} else {
		windStatus = "bahaya"
	}

	fmt.Printf("status wind : %s\n", windStatus)
}
