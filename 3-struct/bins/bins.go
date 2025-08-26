package bins

import (
	"errors"
	"time"
)

type Bin struct {
	id        string
	private   bool
	createdAt time.Time
	name      string
}

type BinList struct {
	Bins []Bin
}

func NewBin(id string, private bool, name string) (*Bin, error) {
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

func AddNewBin(binList *BinList, bin Bin) {
	binList.Bins = append(binList.Bins, bin)
}