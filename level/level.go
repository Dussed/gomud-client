package level

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strconv"

	"github.com/dussed/gomud-client/levels/leveldetail"
)

type Level struct {
}

func Load(name string) error {
	fp, err := filepath.Abs("./levels/" + name + ".level")
	dat, err := ioutil.ReadFile(fp)

	// breakChars := []byte{10}

	if err != nil {
		return err
	}

	for _, tileData := range dat {

		tileId, err := strconv.Atoi(string(tileData))

		if err != nil {
			fmt.Errorf(err.Error())
		}

		// Is it line break time?
		if tileData == 10 {
			fmt.Println()
		} else {
			fmt.Print(leveldetail.GetTile(tileId).Character)
		}

	}

	fmt.Println()

	return nil
}
