package main

import (
	"flag"
	"log"

	"github.com/140proof/gofc"
)

var email = flag.String("email", "", "Email address")
var apikey = flag.String("apikey", "", "FullContact API Key")

func main() {
	log.Println("Gofc")
	flag.Parse()

	fc := gofc.NewClient(*apikey)

	res, err := fc.Person().GetByEmail(*email)
	if err != nil {
		log.Println("Error:", err)
	}

	log.Println("Response:", res)

}
