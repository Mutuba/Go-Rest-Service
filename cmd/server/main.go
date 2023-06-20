package main

import "fmt"

// responsible for instantiation and startuo of the go app
func run() error {
	fmt.Println("Starting up our application")
	return nil
}

func main() {

	fmt.Println("Mutuba is attempting Golang")

	if err := run(); err != nil {
		panic(err)
		fmt.Println(err)
	}
}
