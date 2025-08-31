package bins

import (
	"3-struct/file"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/fatih/color"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func (bin *Bin) OutputBin() {
	fmt.Printf("id: %s, private: %t, name: %s\n", bin.Id, bin.Private, bin.Name)
}

type BinList struct {
	Bins []Bin `json:"bins"`
}

func (bins *BinList) ToBytes() ([]byte, error) {
	return json.Marshal(bins)
}

func (bins *BinList) save() {
	data, err := bins.ToBytes()
	if err != nil {
		color.Red("Something went wrong... :(")
		return
	}
	file.WriteFile(data, "data.json")
}

func NewBin(id string, private bool, name string) (*Bin, error) {
	if name == "" {
		return nil, errors.New("name can't be empty")
	}
	newBin := &Bin{
		Id:        id,
		Private:   private || false,
		CreatedAt: time.Now(),
		Name:      name,
	}

	return newBin, nil
}

func AddNewBin(binList *BinList, bin Bin) {
	binList.Bins = append(binList.Bins, bin)
	binList.save()
}