package main

import (
	"flag"
	"fmt"

	g "github.com/daffainfo/apiguesser/guesser"
)

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

func main() {
	show_banner()
	api := flag.String("api", "", "An API Key. Example: tue3sv9hzsey1me4l7fwq3t46u5k8wag")
	path := flag.String("path", "", "A file with API Key. Example: daffainfo.txt")
	flag.Parse()

	if *api != "" && *path == "" && len(*api) > 3 {
		fmt.Println("[!] Possible Key:")
		fmt.Println(g.Regex_api(*api))
	} else if *api == "" && *path != "" {
		fmt.Println("[!] Possible Key:")
		g.Regex_api_file(*path)
	} else if *api != "" || *path != "" {
		fmt.Println("[X] Can't call 2 arguments at once")
	}
}
