package bins

import (
	"errors"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

var BinList = []Bin{}

func createdBin(b *Bin) (c *Bin, err error) {
	if b.Id == "" {
		return nil, errors.New("INVALID_ID")
	}
	if b.Name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	return c, nil
}

// func CreatedBinList([]Bin) {

// }
