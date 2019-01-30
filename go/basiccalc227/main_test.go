package main

import "testing"

func TestCalc(t *testing.T) {
	input := []string{
		"1 + 1",
		"2 + 2",
		"2 + 2 /2",
		"2 +2 *2",
		"10 - 5 + 2 * 2",
		"2 + (3 + 4)",
		"(1+5) - 4 - 7",
		"(1+5) - (4 - 7)",
		"((3+5) - 4) - 7",
		"((3+5)) - 4 - 7",
		"3+5 - (4 - 7)",
		"3+5 - ((4 - 7))",
		"(3+5) - ((4 - 7))",
		"(2*2) - (4/2)",
		"(2*2) * (4/2)",
	}
	expected := []float64{
		2.0,
		4.0,
		3.0,
		6.0,
		9.0,
		9.0,
		-5.0,
		9.0,
		-3.0,
		-3.0, // double parens on right
		11.0,
		11.0,
		11.0, // double parens on left
		2.0,
		8.0,
	}

	for i := 0; i < len(input); i++ {
		sol := Calculate(input[i])

		if sol != expected[i] {
			t.Errorf(
				"input %v; output %v; expected %v",
				input[i],
				sol,
				expected[i],
			)

		}

	}

}
