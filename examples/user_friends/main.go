package main

import (
	"fmt"
	"log"

	"github.com/iwajezhgf/vimeworldgo"
)

func main() {
	client := vimeworldgo.NewClient(vimeworldgo.Options{})

	res, err := client.GetUserFriends(6344094)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
