package level

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"reflect"
)

type Level struct {
}

type levelFormat struct {
	Name      string
	LevelData []levelLine
}

type levelLine []byte

func Load(name string) error {
	fp, err := filepath.Abs("./levels/" + name + ".json")
	dat, err := ioutil.ReadFile(fp)

	var parsed levelFormat

	err = json.Unmarshal(dat, &parsed)

	// fmt.Println(parsed)
	// return nil

	// breakChars := []byte{10}

	if err != nil {
		return err
	}

	for _, stageLine := range parsed.LevelData {

		// tileID, err := strconv.Atoi(string(tileData))

		if err != nil {
			fmt.Errorf(err.Error())
		}

		fmt.Println(reflect.TypeOf(stageLine[0]))

		fmt.Println(stageLine)

		// Is it line break time?
		// if tileData == 10 {
		// 	fmt.Println()
		// } else {
		// 	fmt.Print(leveldetail.GetTile(tileId).Character)
		// }

	}

	fmt.Println()

	return nil
}
