package main

import (
	"fmt"
	"3-struct/bins"
	// "3-struct/api"
	// "3-struct/file"
	// "3-struct/storage"
)



func main() {
	id := promptData("Enter bin's id")
	private := promptData("Does your bin is Private? (y/n)")
	isPrivate := private == "y"
	name := promptData("Enter bin's name")

	newBin, err := bins.NewBin(id, isPrivate, name)
	if err != nil {
		fmt.Println(err)
	}

	binList := bins.BinList{}
	if newBin != nil {
		bins.AddNewBin(&binList, *newBin)
		fmt.Println(newBin)
		fmt.Printf("\nYour bin list now has %d bin(s)", len(binList.Bins))
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}
