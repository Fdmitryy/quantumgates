package quantumgates

import (
	"fmt"
	"math"
	"math/cmplx"
)

// I type represents I gate
type I struct {
	matrix [][]complex128
}

// X type represents X gate
type X struct {
	matrix [][]complex128
}

// S type represents S gate
type S struct {
	matrix [][]complex128
}

// H type represents H gate
type H struct {
	matrix [][]complex128
}

// T type represents T gate
type T struct {
	matrix [][]complex128
}

// CNOT type represents CNOT gate
type CNOT struct {
	matrix [][]complex128
}

// NewI method creates new I gate instance
func NewI() *I {
	return &I{
		matrix: [][]complex128{
			[]complex128{1, 0},
			[]complex128{0, 1},
		},
	}
}

// NewX method creates new X gate instance
func NewX() *X {
	return &X{
		matrix: [][]complex128{
			[]complex128{0, 1},
			[]complex128{1, 0},
		},
	}
}

// NewS method creates new S gate instance
func NewS() *S {
	return &S{
		matrix: [][]complex128{
			[]complex128{1, 0},
			[]complex128{0, 1i},
		},
	}
}

// NewH method creates new H gate instance
func NewH() *H {
	var coef complex128 = 1 / math.Sqrt2

	return &H{
		matrix: [][]complex128{
			[]complex128{coef, coef},
			[]complex128{coef, -coef},
		},
	}
}

// NewT method creates new T gate instance
func NewT() *T {
	return &T{
		matrix: [][]complex128{
			[]complex128{1, 0},
			[]complex128{0, cmplx.Sqrt(1i)},
		},
	}
}

// NewCNOT method creates new CNOT gate instance
func NewCNOT() *CNOT {
	return &CNOT{
		matrix: [][]complex128{
			[]complex128{1, 0, 0, 0},
			[]complex128{0, 1, 0, 0},
			[]complex128{0, 0, 0, 1},
			[]complex128{0, 0, 1, 0},
		},
	}
}

// Transform method transforms input qubit
func (i *I) Transform(input []complex128) ([]complex128, error) {
	if len(input) != 2 {
		return nil, fmt.Errorf("input array size must be equal 2")
	}

	res := make([]complex128, 2)

	for index := range res {
		column := i.matrix[index]
		res[index] = input[0]*column[0] + input[1]*column[1]
	}

	return res, nil
}

// Transform method transforms input qubit
func (x *X) Transform(input []complex128) ([]complex128, error) {
	if len(input) != 2 {
		return nil, fmt.Errorf("input array size must be equal 2")
	}

	res := make([]complex128, 2)

	for index := range res {
		column := x.matrix[index]
		res[index] = input[0]*column[0] + input[1]*column[1]
	}

	return res, nil
}

// Transform method transforms input qubit
func (s *S) Transform(input []complex128) ([]complex128, error) {
	if len(input) != 2 {
		return nil, fmt.Errorf("input array size must be equal 2")
	}

	res := make([]complex128, 2)

	for index := range res {
		column := s.matrix[index]
		res[index] = input[0]*column[0] + input[1]*column[1]
	}

	return res, nil
}

// Transform method transforms input qubit
func (h *H) Transform(input []complex128) ([]complex128, error) {
	if len(input) != 2 {
		return nil, fmt.Errorf("input array size must be equal 2")
	}

	res := make([]complex128, 2)

	for index := range res {
		column := h.matrix[index]
		res[index] = input[0]*column[0] + input[1]*column[1]
	}

	return res, nil
}

// Transform method transforms input qubit
func (t *T) Transform(input []complex128) ([]complex128, error) {
	if len(input) != 2 {
		return nil, fmt.Errorf("input array size must be equal 2")
	}

	res := make([]complex128, 2)

	for index := range res {
		column := t.matrix[index]
		res[index] = input[0]*column[0] + input[1]*column[1]
	}

	return res, nil
}

// Transform method transforms input qubit
func (cn *CNOT) Transform(controller,
	input []complex128) ([]complex128, []complex128, error) {

	if len(controller) != 2 || len(input) != 2 {
		return nil, nil, fmt.Errorf("controller and input arrays size must be equal 2")
	}

	united := []complex128{
		controller[0] * input[0],
		controller[0] * input[1],
		controller[1] * input[0],
		controller[1] * input[1],
	}

	unitedRes := make([]complex128, 4)

	for index := range unitedRes {
		column := cn.matrix[index]
		unitedRes[index] = united[0]*column[0] + united[1]*column[1] +
			united[2]*column[2] + united[3]*column[3]
	}

	a := cmplx.Sqrt(cmplx.Pow(unitedRes[0], 2) + cmplx.Pow(unitedRes[1], 2))
	b := cmplx.Sqrt(cmplx.Pow(unitedRes[2], 2) + cmplx.Pow(unitedRes[3], 2))

	var c, d complex128

	if b != 0 {
		c = unitedRes[2] / b
		d = unitedRes[3] / b
	} else if a != 0 {
		c = unitedRes[0] / a
		d = unitedRes[1] / a
	} else {
		if controller[0] != 1 {
			c = input[1]
			d = input[0]
		} else {
			c = input[0]
			d = input[1]
		}
	}

	res := []complex128{c, d}

	return controller, res, nil
}
