package main

import (
	"fmt"
	"log"

	"github.com/Fdmitryy/quantumgates"
)

func main() {
	input := []complex128{0.6, 0.8}
	fmt.Println("Input: ", input)

	I := quantumgates.NewI()
	output, err := I.Transform(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("I gate:", output)

	X := quantumgates.NewX()
	output, err = X.Transform(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("X gate:", output)

	S := quantumgates.NewS()
	output, err = S.Transform(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("S gate:", output)

	H := quantumgates.NewH()
	output, err = H.Transform(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("H gate:", output)

	T := quantumgates.NewT()
	output, err = T.Transform(input)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("T gate:", output)
}
