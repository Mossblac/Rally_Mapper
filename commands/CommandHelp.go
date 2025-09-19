package commands

import "fmt"

func CommandHelp(string) error {
	fmt.Println("Enter number of punches")
	fmt.Println("Enter physical equipment avalable")
	fmt.Println("Enter number of desired obstacles")
	fmt.Println("Enter 0-4 to reprisent percentage of required obstacles (0 = 0%, 2 = 50%, 4 = 100%)")
	fmt.Println("Enter maximum size of area in square feet - default is unlimited")
	fmt.Println("verify info and confirm")

	return nil
}

func Wid_Help_Button() {
	fmt.Println("Enter number of punches")
	fmt.Println("Enter physical equipment avalable")
	fmt.Println("Enter number of desired obstacles")
	fmt.Println("Enter 0-4 to reprisent percentage of required obstacles (0 = 0%, 2 = 50%, 4 = 100%)")
	fmt.Println("Enter maximum size of area in square feet - default is unlimited")
	fmt.Println("verify info and confirm")
}
