package main

import (
	"fmt"
	"log"

	"github.com/iwajezhgf/vimeworldgo"
)

func main() {
	client := vimeworldgo.NewClient(vimeworldgo.Options{})

	usernames := []string{"3APADD", "DontCry__", "MomentoMorry", "s0ullesss", "xtrafrancyz"}

	res, err := client.GetUsersByNames(usernames...)
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(res)
}
