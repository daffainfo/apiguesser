package main

import (
	"flag"
	"fmt"
	// g "github.com/daffainfo/apiguesser/guesser"
)

var (
	Red   = Color("\033[1;31m%s\033[0m")
	Green = Color("\033[1;32m%s\033[0m")
	Blue  = Color("\033[1;34m%s\033[0m")
	Cyan  = Color("\033[1;36m%s\033[0m")
)

func Color(colorString string) func(...interface{}) string {
	sprint := func(args ...interface{}) string {
		return fmt.Sprintf(colorString,
			fmt.Sprint(args...))
	}
	return sprint
}

func show_banner() {
	fmt.Println(Blue(`                                          
         _                                 
 ___ ___|_|___ ___ _ _ ___ ___ ___ ___ ___ 
| .'| . | |___| . | | | -_|_ -|_ -| -_|  _|
|__,|  _|_|   |_  |___|___|___|___|___|_|  
    |_|       |___|                        

Author: Muhammad Daffa
Version: 1.0

Starting...
`))
}

func main() {
	show_banner()
	api := flag.String("api", "", "An API Key. Example: tue3sv9hzsey1me4l7fwq3t46u5k8wag")
	path := flag.String("path", "", "A file with API Key. Example: daffainfo.txt")
	flag.Parse()

	if *api != "" && *path == "" && len(*api) > 3 {
		fmt.Println(Cyan(*api))
		if Regex_api(*api) != "" {
			fmt.Println(Green(Regex_api(*api)))
		} else {
			fmt.Println(Red("Not Match"))
		}

	} else if *api == "" && *path != "" {
		Regex_api_file(*path)
	} else if *api != "" || *path != "" {
		fmt.Println(Red("Can't call 2 arguments at once"))
	}
}
