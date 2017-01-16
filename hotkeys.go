package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"os/user"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

// Hotkey maps a x11 `Key` with a given `Desc` description to an action `Cmd`, which is passed to a shell, so shellisms are accepted
type Hotkey struct {
	Key  string `json:"key'`
	Desc string `json:"desc"`
	Cmd  string `json:"cmd"`
}

func main() {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}
	keybind.Initialize(X)

	currentuser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}
	configfile := currentuser.HomeDir + "/.config/hotkeys.conf.json"

	err = bindall(configfile, X)
	if err != nil {
		log.Println(err)
	}

	err = bindall(configfile, X)
	if err != nil {
		log.Panicln(err)
	}
	xevent.Main(X)
}

func bindall(configfile string, X *xgbutil.XUtil) (err error) {
	config, err := ioutil.ReadFile(configfile)
	if err != nil {
		log.Fatal("Could not find config file: ", err)
		return
	}
	hotkeys := []Hotkey{}
	err = json.Unmarshal(config, &hotkeys)
	if err != nil {
		log.Fatal("Could not parse config file:", err)
		return
	}
	keybind.Detach(X, X.RootWin())
	for _, hotkey := range hotkeys {
		hotkey.attach(X)
	}
	return
}

func (hotkey Hotkey) attach(X *xgbutil.XUtil) {
	log.Println(hotkey.Key)
	err := keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			go exec.Command("/bin/sh", "-c", hotkey.Cmd).Run()
		}).Connect(X, X.RootWin(), hotkey.Key, true)
	if err != nil {
		log.Fatalf("Could not bind %s: %s", hotkey.Key, err.Error())
	}
}
