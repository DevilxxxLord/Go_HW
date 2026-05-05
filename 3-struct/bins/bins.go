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

var BinList = []Bin{}

func createdBin(b *Bin) (c *Bin, err error) {
	if b.id == "" {
		return nil, errors.New("INVALID_ID")
	}
	if b.name == "" {
		return nil, errors.New("INVALID_NAME")
	}
	return c, nil
}

func createdBinList([]Bin) {

}
