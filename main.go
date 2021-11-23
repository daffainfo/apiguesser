package main

import (
	"bufio"
	"encoding/json"
	"flag"
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

func show_banner() {
	fmt.Println(`                                          
         _                                 
 ___ ___|_|___ ___ _ _ ___ ___ ___ ___ ___ 
| .'| . | |___| . | | | -_|_ -|_ -| -_|  _|
|__,|  _|_|   |_  |___|___|___|___|___|_|  
    |_|       |___|                        

Author: Muhammad Daffa
Version: 1.0`)
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
		Regex_api(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func Regex_api(contents string) {
	var data []regex_data

	resp, err := http.Get("https://raw.githubusercontent.com/daffainfo/ApiGuesser/main/db.json")
	if err != nil {
		fmt.Println("No response from request")
	}
	defer resp.Body.Close()

	byteValue, _ := ioutil.ReadAll(resp.Body)

	var errjson = json.Unmarshal([]byte(byteValue), &data)
	if errjson != nil {
		fmt.Println(err.Error())
		return
	}

	for i := range data {
		re := regexp.MustCompile(data[i].Regex)
		if re.MatchString(contents) {
			res1 := re.FindAllString(contents, 1)
			fmt.Println(data[i].Name, res1)
		}
	}
}

func main() {
	show_banner()
	api := flag.String("api", "", "An API Key. Example: tue3sv9hzsey1me4l7fwq3t46u5k8wag")
	path := flag.String("path", "", "A file with API Key. Example: daffainfo.txt")
	flag.Parse()

	if *api != "" && *path == "" && len(*api) > 3 {
		fmt.Println("Possible API Key:")
		Regex_api(*api)
	} else if *api == "" && *path != "" {
		Regex_api_file(*path)
	} else if *api != "" || *path != "" {
		fmt.Println("Can't call 2 arguments at once")
	}
}
