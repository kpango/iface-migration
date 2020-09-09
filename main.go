package main

import "log"

func main() {
	a, err := New()

	log.Printf("%#v", a)
	if err != nil {
		log.Fatalln(err)
	}
}
