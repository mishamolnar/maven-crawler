package main

import (
	"fmt"
	"maven_crawler/crawl"
	"os"
)

func main() {
	links, err := crawl.Extract(os.Args[1])
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(links)
}
