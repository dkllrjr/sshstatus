package main

import (
	"log"
)

func LogError(err error) {

	if err != nil {
		log.Fatal(err)
	}

}

func LogNoError(err error) {
	
	if err != nil {
		log.Print(err)
	}

}

func LogNotError(note string) {
	
	log.Print(note)

}
