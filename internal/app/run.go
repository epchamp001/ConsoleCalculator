package app

import (
	"Calculator/internal/usecase/handlers/input/exit"
	"Calculator/internal/usecase/handlers/input/validator"
	"Calculator/internal/usecase/operation"
	"bufio"
	"fmt"
	"math"
	"os"
	"strings"

	"github.com/fatih/color"
)

func Run() {
	fmt.Println("\n____________________________________________________")
	color.New(color.FgHiBlue, color.Bold).Println("Welcome to the Console Calculator!")
	color.New(color.FgHiCyan).Println("To end the program, enter \"exit\"")
	fmt.Println("____________________________________________________\n")

	reader := bufio.NewReader(os.Stdin)

	for {
		color.New(color.FgHiGreen).Printf("Input: ")

		input, err := reader.ReadString('\n')
		if err != nil {
			color.New(color.FgHiRed).Println("ERROR: Failed to read input:", err)
			continue
		}

		input = strings.TrimSpace(input)

		exit.ExitProg(input)

		input, err = validator.ValidateExpression(input)
		if err != nil {
			color.New(color.FgHiRed).Println("ERROR: ", err)
			continue
		}

		res, err := operation.CalculateExpression(input)
		if err != nil {
			color.New(color.FgHiRed).Println("ERROR: ", err)
			continue
		}

		color.New(color.FgHiYellow, color.Bold).Printf("\nResult: %.2f\n", math.Round(res*100)/100)
		fmt.Println("____________________________________________________\n")
	}
}
