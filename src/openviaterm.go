package main

import (
	"fmt"

	"github.com/cli/browser"
)

func main() {

	var input_url string

	fmt.Println("Welcome to openviaterm!")
	fmt.Println("WARNING: THIS IS NOT A TERMINAL BROWSER! IT OPENS A NEW BROWSER WINDOW!")

	fmt.Print("What URL would you like to open?: ")

	fmt.Scan(&input_url)

	err := browser.OpenURL(input_url)

	if err != nil {

		fmt.Println("Error opening URL:", err)
		return
	}

}
