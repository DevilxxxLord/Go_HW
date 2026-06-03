package bins

import (
	"errors"
	"fmt"
	"time"
)

type Bin struct {
	Id        string    `json:"id"`
	Private   bool      `json:"private"`
	CreatedAt time.Time `json:"createdAt"`
	Name      string    `json:"name"`
}

func checkBin(b *Bin) error {
	if b.Id == "" {
		return errors.New("INVALID_ID")
	}
	if b.Name == "" {
		return errors.New("INVALID_NAME")
	}
	return nil
}

func (b *Bin) CreatedBinList(id string, private bool, name string) (binList []Bin) {
	b.Id = id
	b.Private = private
	b.CreatedAt = time.Now()
	b.Name = name
	err := checkBin(b)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	binList = append(binList, *b)
	return binList
}
