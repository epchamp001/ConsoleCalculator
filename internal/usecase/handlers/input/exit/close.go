package exit

import (
	"fmt"
	"github.com/fatih/color"
	"os"
)

func ExitProg(input string) {
	if input == "exit" {
		fmt.Println()
		fmt.Println("____________________________________________________")
		color.New(color.FgHiMagenta, color.Bold).Println("Bye! Come again! 😊")
		fmt.Println("____________________________________________________\n")

		os.Exit(0)
	}
}
