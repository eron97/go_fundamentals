package main

import "fmt"

func main() {

	animals := struct {
		Mammals struct {
			Race  string
			Color string
			Size  string
			/// ...
		}
		Amphibians struct{}
		/// ...
	}{
		/// ...
	}

	fmt.Println(animals)

}
