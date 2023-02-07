package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/kravetsd/chose-your-own-adventure/cyoa"
	"github.com/kravetsd/chose-your-own-adventure/cyoaweb"
)

func main() {
	fmt.Println("Hello, cyoa!")
	fl, err := os.Open("gopher.json")
	if err != nil {
		log.Fatal("Openning file:", err)
	}
	defer fl.Close()

	jsondec := json.NewDecoder(fl)

	var bk cyoa.Story

	err = jsondec.Decode(&bk)

	fmt.Printf("%+v", bk["intro"])

	if err != nil {
		log.Fatal("Error decoding json:", err)
	}

	cyoaweb.RunStoryWeb()

}
