package exercise5

import (
	"bufio"
	"exercises-in-programming/util"
	"fmt"
	"log"
	"os"
	"strconv"
)

// Simple math
func Run() {

	r := bufio.NewReader(os.Stdin)
	xS, err := util.GetInput(r, "What is the first number? ")

	if err != nil {
		log.Fatalln(err)
	}

	yS, err := util.GetInput(r, "What is the second number? ")

	if err != nil {
		log.Fatalln(err)
	}

	x, err := strToInt(xS)

	if err != nil {
		log.Fatalln(err)
	}

	y, err := strToInt(yS)

	if err != nil {
		log.Fatalln(err)
	}

	operations := [4]string{"+", "-", "*", "/"}

	for _, op := range operations {
		out, err := operation(x, y, op)

		if err != nil {
			log.Fatalln(err)
		}

		printOut(x, y, out, op)
	}

}

func strToInt(s string) (int, error) {
	x, err := strconv.Atoi(s)

	if err != nil {
		return 0, err
	}

	if x < 0 {
		return 0, fmt.Errorf("Input %d can't below 0", x)
	}

	return x, nil
}

func operation(x int, y int, op string) (int, error) {
	switch op {
	case "+":
		return x + y, nil
	case "-":
		return x - y, nil
	case "*":
		return x * y, nil
	case "/":
		return x / y, nil
	default:
		return 0, fmt.Errorf("operation %s not supported", op)
	}
}

func printOut(x int, y int, out int, op string) {
	fmt.Printf("%d %s %d = %d\n", x, op, y, out)
}
