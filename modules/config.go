package modules

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"

	"github.com/Schmenn/hub/modules/debug"
	"github.com/Schmenn/hub/structs"
)

var config structs.Config

// GetConfig Reads config.json and returns the settings or creates a new
func GetConfig() structs.Config {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	debug.Check(err)
	fmt.Println(dir)
	f, err := ioutil.ReadFile(path.Join(dir, "config.json"))
	if err != nil {
		fmt.Println("File could not be opened, either it does not exist or the program does not have to permissions to read it\nCreating a new file for you")
		placeHolder, err := json.MarshalIndent(structs.Config{StartUpSound: "REPLACE-WITH-PATH-TO-AUDIO-FILE", Name: "YOUR-NAME"}, "", "    ")
		debug.Check(err)

		//fmt.Println(placeHolder)
		ioutil.WriteFile(path.Join(dir, "config.json"), placeHolder,  0777)

		os.Exit(1)
	} else {
		json.Unmarshal(f, &config)
	}
	

	return config
}
