package adapted

import (
	"main/bins"
	"main/file"
	"main/storage"
)

type All interface {
	SaveStorage(data []byte, name string)
	ReadFiles(name string) ([]byte, bool, error)
	ReadJson(name string) ([]bins.Bin, error)
	CreatedBinList(id string, priv bool, name string)
}

type Adapted struct {
}

func (a *Adapted) SaveStorage(data []byte, name string) {
	storage.SaveStorage(data, name)
}

func (a *Adapted) ReadFiles(name string) ([]byte, bool, error) {
	return file.ReadFiles(name)
}

func (a *Adapted) ReadJson(name string) ([]bins.Bin, error) {
	return storage.ReadJson(name)
}

func (a *Adapted) CreatedBinList(id string, priv bool, name string) {
	b := &bins.Bin{}
	b.CreatedBinList(id, priv, name)
}
