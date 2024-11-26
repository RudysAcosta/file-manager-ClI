package menu

import (
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func MainMenu() (int, error) {

	color.Red("Select an option:")
	fmt.Println("1. Command Line")
	fmt.Println("2. Interactive")
	fmt.Println("3. Exit")

	var option int
	fmt.Scanln(&option)

	if option < 1 || option > 3 {
		return option, errors.New("invalid option")
	}

	return option, nil
}

func InteractiveMenu() {
	color.Blue("Interactive")
}
