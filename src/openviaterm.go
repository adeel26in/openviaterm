package main

import (
	"fmt"

	"github.com/cli/browser"
)

func main() {

	var input_first string

	fmt.Println("Welcome to openviaterm!")

	fmt.Print("What would you like to open? File or URL: ")

	fmt.Scan(&input_first)

	switch input_first {
	case "File", "file":
		var file_path string

		fmt.Print("Enter the file path: ")
		fmt.Println("USE THIS AT YOUR OWN RISK! CAN BE BROKEN OR CAN GLITCH!")
		fmt.Scan(&file_path)

		err := browser.OpenFile(file_path)

		if err != nil {

			fmt.Println("Error opening file:", err)
		}
	case "URL", "url":
		var url string

		fmt.Print("Enter the URL: ")
		fmt.Scan(&url)

		err := browser.OpenURL(url)

		if err != nil {

			fmt.Println("Error opening URL:", err)
		}
	}

}
