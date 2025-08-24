package main

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	bins []Bin
}

func main() {
	id := promptData("Enter bin's id")
	private := promptData("Does your bin is Private? (y/n)")
	isPrivate := private == "y"
	name := promptData("Enter bin's name")

	newBin, err := newBin(id, isPrivate, name)
	if err != nil {
		fmt.Println(err)
	}

	binList := BinList{}
	if newBin != nil {
		addNewBin(&binList, *newBin)
		fmt.Printf("\nYour bin list now has %d bin(s)", len(binList.bins))
	}
}

func promptData(prompt string) string {
	fmt.Print(prompt + ": ")
	var res string
	fmt.Scanln(&res)
	return res
}

func newBin(id string, private bool, name string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("name can't be empty")
	}
	newBin := &Bin{
		id: id,
		private: private || false,
		createdAt: time.Now(),
		name: name,
	}

	return newBin, nil
}

func addNewBin(binList *BinList, bin Bin) {
	binList.bins = append(binList.bins, bin)
}
