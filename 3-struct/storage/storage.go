package storage

import (
	"3-struct/bins"
	"3-struct/file"
	"encoding/json"

	"github.com/fatih/color"
)

func NewBinList() *bins.BinList {
	file, err := file.ReadFile("data.json")
	if err != nil {
		return &bins.BinList{
			Bins: []bins.Bin{},
		}
	}
	var binList bins.BinList
	err = json.Unmarshal(file, &binList)
	if err != nil {
		color.Red("Can't parse bins")
		return &bins.BinList{
			Bins: []bins.Bin{},
		}
	}
	return &binList
}

// func SaveBin(id string, isPrivate bool, name string) {
// 	bin, err := bins.NewBin(id, isPrivate, name)
// }