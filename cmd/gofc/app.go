package main

import (
	"flag"
	"log"

	"github.com/140proof/gofc"
)

var apikey = flag.String("apikey", "", "FullContact API Key")

var method = flag.String("method", "email", "Lookup Method (email|twitter|fbusername|fbid|phone)")
var value = flag.String("value", "", "Value to lookup")

func main() {
	log.Println("Gofc")
	flag.Parse()

	fc := gofc.NewClient(*apikey)

	var res *gofc.PersonResponse
	var err error

	switch *method {
	case "email":
		res, err = fc.Person().GetByEmail(*value)
	case "twitter":
		res, err = fc.Person().GetByTwitter(*value)
	case "phone":
		res, err = fc.Person().GetByPhone(*value)
	case "fbusername":
		res, err = fc.Person().GetByFacebookUsername(*value)
	case "fbid":
		res, err = fc.Person().GetByFacebookID(*value)
	}

	if err != nil {
		log.Println("Error:", err)
	}

	log.Println("Response:", res)

}
