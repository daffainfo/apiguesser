package guesser

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
)

type regex_data struct {
	Name  string `json:"Name"`
	Regex string `json:"Regex"`
}

func Regex_api_file(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
		fmt.Println(Regex_api(scanner.Text()))
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Regex_api(contents string) string {
	var data []regex_data
	var result string

	resp, err := http.Get("https://raw.githubusercontent.com/daffainfo/apiguesser/main/db.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var errjson = json.Unmarshal(body, &data)
	if errjson != nil {
		fmt.Println(err.Error())
	}

	for i := range data {
		re := regexp.MustCompile(data[i].Regex)
		if re.MatchString(contents) {
			result = data[i].Name
		}
	}
	return result
}
