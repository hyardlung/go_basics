package main

import (
	"3-struct/bins"
	"3-struct/storage"
	"errors"
	"fmt"

	"github.com/fatih/color"
)

func main() {
	fmt.Println("___ Bins manager ___")
	binList := storage.NewBinList()
Main:
	for {
		v, err := getMenuVariant()
		if err != nil {
			color.Red(err.Error())
			continue
		}
		switch v {
		case 1:
			createBin(binList)
		case 2:
			showBins(binList)
		default:
			break Main
		}
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func createBin(binList *bins.BinList) {
	id := promptData("Enter bin's id")
	private := promptData("Does your bin is Private? (y/n)")
	isPrivate := private == "y"
	name := promptData("Enter bin's name")

	newBin, err := bins.NewBin(id, isPrivate, name)
	if err != nil {
		fmt.Println(err)
	}

	if newBin != nil {
		bins.AddNewBin(binList, *newBin)
		fmt.Printf("\nYour bin list now has %d bin(s)\n", len(binList.Bins))
	}
}

func showBins(bins *bins.BinList) {
	for _, b := range bins.Bins {
		b.OutputBin()
	}
}

func getMenuVariant() (uint8, error) {
	var variant uint8
	fmt.Println("1. Create bin")
	fmt.Println("2. Show all bins")
	fmt.Println("3. Exit")
	fmt.Scan(&variant)
	if variant < 1 || variant > 3 {
		return 0, errors.New("choose from menu")
	}
	return variant, nil
}
