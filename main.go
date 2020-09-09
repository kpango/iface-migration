package main

import "log"

func main() {
	a, err := New()
	if err != nil {
		log.Fatalln(err)
	}
	log.Printf("%#v\n", a)
}
