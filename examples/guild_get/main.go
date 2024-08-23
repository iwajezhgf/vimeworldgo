package main

import (
	"fmt"
	"log"

	"github.com/iwajezhgf/vimeworldgo"
)

func main() {
	client := vimeworldgo.NewClient(vimeworldgo.Options{})

	// image links are proxied
	// use client.GetGuildById(15048, false) to prevent it
	res, err := client.GetGuildById(10567)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
