package main

import (
	"fmt"
	"log"

	"github.com/Fdmitryy/quantumgates"
)

func main() {
	input := []complex128{0.6, 0.8}
	fmt.Println("Input: ", input)

	output, err := quantumgates.NewI().Transform(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("I gate:", output)

	output, err = quantumgates.NewX().Transform(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("X gate:", output)

	output, err = quantumgates.NewS().Transform(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("S gate:", output)

	output, err = quantumgates.NewH().Transform(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("H gate:", output)

	output, err = quantumgates.NewT().Transform(input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("T gate:", output)

	CNOT := quantumgates.NewCNOT()
	controller1, output1, err := CNOT.Transform([]complex128{0, 1}, input)
	if err != nil {
		log.Fatal(err)
	}

	controller2, output2, err := CNOT.Transform([]complex128{1, 0}, input)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf(
		"CNOT gate:\n\tcontroller1: %v, output1: %v\n\tcontroller2: %v, output2: %v\n",
		controller1, output1, controller2, output2)
}
