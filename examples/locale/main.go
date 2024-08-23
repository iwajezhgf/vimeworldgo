package main

import (
	"fmt"
	"log"

	"github.com/iwajezhgf/vimeworldgo"
)

func main() {
	client := vimeworldgo.NewClient(vimeworldgo.Options{})

	res, err := client.GetLocale("en", "ranks")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
