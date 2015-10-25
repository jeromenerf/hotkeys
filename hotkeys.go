package main

import (
	"log"
	"os/exec"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

type Hotkey struct {
	key  string
	desc string
	cmd  string
}

func main() {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	keybind.Initialize(X)

	// FIXME: use a json config?
	hotkeys := []Hotkey{
		Hotkey{
			key:  "Print",
			desc: "Take a screenshot of the active window",
			cmd:  "import -window \"$(xdotool getwindowfocus)\" png:- | tee ~/tmp/screenshot-$(date +\"%Y-%m-%d-%T\").png | xclip -t image/png -selection c",
		},
		Hotkey{
			key:  "Mod1-Print",
			desc: "Take a screenshot of a user selected area",
			cmd:  "import png:- | tee ~/tmp/screenshot-$(date +'%Y-%m-%d-%T').png | xclip -t image/png -selection c",
		},
		Hotkey{
			key:  "Mod1-Shift-Print",
			desc: "Take a screenshot of the root window",
			cmd:  "import -window root png:-  | tee ~/tmp/screenshot-$(date +'%Y-%m-%d-%T').png | xclip -t image/png -selection c",
		},
	}

	for _, hotkey := range hotkeys {
		hotkey.attach(X)
	}
	xevent.Main(X)
}

func (hk Hotkey) attach(X *xgbutil.XUtil) {
	err := keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			log.Println(hk.key, hk.cmd)
			go exec.Command("/bin/sh", "-c", hk.cmd).Run()
		}).Connect(X, X.RootWin(), hk.key, true)
	if err != nil {
		log.Fatal(err)
	}
}
