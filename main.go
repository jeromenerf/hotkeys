package main

import (
	"log"
	"os/exec"

	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
)

func main() {
	X, err := xgbutil.NewConn()
	if err != nil {
		log.Fatal(err)
	}

	keybind.Initialize(X)

	// FIXME: use a json config?
	maps := map[string]string{
		"Mod1-Print":               "import png:-  | tee \"~/tmp/screenshot-$(date +'%Y-%m-%d-%T')\".png | xclip -t image/png -selection c",
		"Mod1-shift-Print":         "import -window \"$(xdotool getwindowfocus)\" png:- | tee \"~/tmp/screenshot-$(date +'%Y-%m-%d-%T')\".png | xclip -t image/png -selection c",
		"Mod1-control-shift-Print": "import -window root png:-  | tee \"~/tmp/screenshot-$(date +'%Y-%m-%d-%T')\".png | xclip -t image/png -selection c",
	}

	for k, v := range maps {
		err = keybind.KeyPressFun(
			func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
				err := exec.Command("/bin/sh", "-c", v).Run()
				if err != nil {
					log.Printf("Command failed with %s\n", err)
				}
			}).Connect(X, X.RootWin(), k, true)
		if err != nil {
			log.Fatal(err)
		}
	}
	xevent.Main(X)
}
