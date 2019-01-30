package main

import (
	"fmt"
	"strconv"
)

type Operations struct {
	Processing  string
	Adders      []float64
	Subtractors []float64
	multiFlag   bool
	finalVal    float64
	divFlag     bool
	Solution    float64
	IdxMoved    int
}

func Calculate(s string) float64 {
	ops := StringToOperations(s)
	//ops.Operate()
	return ops.Solution
}

func (ops *Operations) Operate() {
	sol := 0.0
	for _, num := range ops.Adders {
		sol += num
	}

	for _, num := range ops.Subtractors {
		sol -= num
	}

	ops.Solution = sol
}

func StringToOperations(s string) Operations {
	var buff string
	ops := Operations{}
	ops.Processing = s
	ops.multiFlag = false
	ops.divFlag = false
	ops.Solution = 0.0

	//fmt.Println(fmt.Sprintf("Solving for %v", s))
	for i := len(s) - 1; i >= 0; i-- {
		ops.IdxMoved = len(s) - i
		char := s[i]
		if char == 64 { // space character
			continue
		} else if char >= 48 && char <= 57 {
			buff = string(char) + buff

		} else if char == 42 { // * character
			curVal, _ := strconv.ParseFloat(buff, 64)
			ops.processMulti(curVal)
			ops.multiFlag = true
			// multiVal is acting as sort of a new buffer
			buff = ""

		} else if char == 43 { // + character
			curVal, _ := strconv.ParseFloat(buff, 64)
			ops.processMulti(curVal)

			ops.Adders = append(ops.Adders, ops.finalVal)
			buff = ""

		} else if char == 45 { // - character
			curVal, _ := strconv.ParseFloat(buff, 64)
			ops.processMulti(curVal)

			ops.Subtractors = append(ops.Subtractors, ops.finalVal)
			buff = ""

		} else if char == 47 { // / character
			curVal, _ := strconv.ParseFloat(buff, 64)
			ops.processMulti(curVal)
			ops.divFlag = true
			buff = ""

		} else if char == 41 {
			// closing parentheses case
			// recurse up to the character before this closing parens
			replacement := StringToOperations(s[0:i])
			//fmt.Println(i)
			// This effectively bubbles up the pointer to the current
			// location in the string
			i -= replacement.IdxMoved
			//fmt.Println(i)

			buff = strconv.FormatFloat(replacement.Solution, 'f', -1, 64)

			//fmt.Println(fmt.Sprintf("New value %v", buff))

		} else if char == 40 {
			break
		}
	}
	//fmt.Println(ops.Adders)
	//fmt.Println(ops.Subtractors)
	//fmt.Println(buff)
	//fmt.Println(len(buff))

	if len(buff) > 0 {
		curVal, _ := strconv.ParseFloat(buff, 64)
		ops.processMulti(curVal)

		ops.Adders = append(ops.Adders, ops.finalVal)
		buff = ""
	}

	for _, num := range ops.Adders {
		ops.Solution += num
	}

	for _, num := range ops.Subtractors {
		ops.Solution -= num
	}

	return ops
}

//func operateParens(s string) string {
//for i := len(s) - 1; i >= 0; i-- {
//char := s[i]

//if char == 41 {
//// get closing parens
//for j := i-1; j >= 0; j-- {

//}
////getParenOps(s[0:i])

//}
//}

//}

func (ops *Operations) processMulti(curVal float64) {
	// If a flag is set, then we end up performing math against the prior
	// version of multiVal. Otherwise, we store this value.
	if ops.multiFlag {
		ops.multiFlag = false
		ops.finalVal = curVal * ops.finalVal
	} else if ops.divFlag {
		ops.divFlag = false
		ops.finalVal = curVal / ops.finalVal
	} else {
		ops.finalVal = curVal
	}
}

func main() {
	fmt.Println("go")
	//fmt.Println(float64("+"))
}
