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

// var (
// 	Red   = Color("\033[1;31m%s\033[0m")
// 	Green = Color("\033[1;32m%s\033[0m")
// 	Blue  = Color("\033[1;34m%s\033[0m")
// 	Cyan  = Color("\033[1;36m%s\033[0m")
// )

// func Color(colorString string) func(...interface{}) string {
// 	sprint := func(args ...interface{}) string {
// 		return fmt.Sprintf(colorString,
// 			fmt.Sprint(args...))
// 	}
// 	return sprint
// }

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
			result = data[i].Name[i]
		}
	}
	return result
}
