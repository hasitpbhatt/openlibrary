package main

import (
	"fmt"
	"log"

	"github.com/hasitpbhatt/openlibrary/olclient"
)

func main() {
	// credentials are saved in ".credentials" file in current folder
	//
	// -- .credentials --
	// {"email": "root","password": "root"}

	ol, err := olclient.NewOpenLibrary(false)
	if err != nil {
		log.Fatal(err)
	}

	if w, err := ol.Get("OL2163622M"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(w.Title, w.Authors, w.Description)
	}

	if w, err := ol.GetByISBN("9780395595114"); err != nil {
		log.Fatal(err)
	} else {
		fmt.Println(w.Title, w.Authors, w.Description)
	}
}
