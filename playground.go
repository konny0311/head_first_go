package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

type Remo struct {
	Name         string `json:"name"`
	NewestEvents struct {
		Te struct {
			Val float64 `json:"val"`
		}
	} `json:"newest_events"`
}

func main() {
	url := "https://api.nature.global/1/devices"
	req, _ := http.NewRequest("GET", url, nil)
	token := fmt.Sprintf("Bearer %s", os.Getenv("REMO_TOKEN"))
	req.Header.Set("Authorization", token)

	client := new(http.Client)
	res, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	body, err := ioutil.ReadAll(res.Body)

	var remos []Remo
	err = json.Unmarshal(body, &remos)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(remos[0].NewestEvents.Te.Val)
}
