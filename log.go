package main

import (
	"log"
)

func LogError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

func LogNotError(err error) {
	
	if err != nil {
		log.Print(err)
	}

}
