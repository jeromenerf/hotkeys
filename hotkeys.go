package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os/exec"
	"os/user"

	"golang.org/x/exp/inotify"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

type Hotkey struct {
	Key  string `json:key`
	Desc string `json:desc`
	Cmd  string `json:cmd`
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

	watcher, err := inotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	err = watcher.AddWatch(configfile, inotify.IN_CLOSE_WRITE)
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		for {
			select {
			case ev := <-watcher.Event:
				log.Println(ev)
				err := bindall(configfile, X)
				if err != nil {
					log.Println(err)
					continue
				}

			case err := <-watcher.Error:
				log.Println("error:", err)
			}
		}
	}()
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
