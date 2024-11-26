package exit

import (
	"fmt"
	"github.com/fatih/color"
	"os"
	"strings"
)

func ExitProg(input string) {
	if strings.ToLower(strings.TrimSpace(input)) == "exit" {
		fmt.Println()
		fmt.Println("____________________________________________________")
		color.New(color.FgHiMagenta, color.Bold).Println("Bye! Come again! 😊")
		fmt.Println("____________________________________________________\n")

		os.Exit(0)
	}
}
