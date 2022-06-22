package main

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
	Name  []string `json:"Name"`
	Regex string   `json:"Regex"`
}

func Regex_api_file(path string) {
	file, err := os.Open(path)
	if err != nil {
		log.Fatal(Red(err))
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(Cyan(scanner.Text()))
		if Regex_api(scanner.Text()) != "" {
			fmt.Println(Green(Regex_api(scanner.Text())))
		} else {
			fmt.Println(Red("Not match"))
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(Red(err))
	}
}

func Regex_api(contents string) string {
	var data []regex_data
	var result string

	resp, err := http.Get("https://pastebin.com/raw/9BuLKBUG")
	if err != nil {
		fmt.Println(Red("No response from request"))
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	var errjson = json.Unmarshal(body, &data)
	if errjson != nil {
		fmt.Println(Red(err.Error()))
	}

	for i := range data {
		re := regexp.MustCompile(data[i].Regex)
		if re.MatchString(contents) {
			for _, str := range data[i].Name {
				result += str + "\n"
			}
		}
	}
	return result
}
