package main

import (
	"fmt"
	"log"
	"quantumgates"
)

func main() {
	a := []complex128{1 + 2i, 2 - 3i}

	i := quantumgates.NewI()
	res, err := i.Transform(a)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(res)
}
