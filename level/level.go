package level

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
)

type Level struct {
}

func Load(name string) error {
	fp, err := filepath.Abs("./levels/" + name + ".level")
	dat, err := ioutil.ReadFile(fp)

	if err != nil {
		return err
	}

	fmt.Println(dat)

	return nil
}
