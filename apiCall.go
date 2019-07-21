package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {

	type APIReturn struct {
		Quote       string
		ID          int
		Designation string
	}

	var APIBody APIReturn

	resp, err := http.Get("https://cascades.dev/api")

	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	stringified := string(body)

	json.Unmarshal([]byte(stringified), &APIBody)

	fmt.Println(APIBody.Quote)

}
