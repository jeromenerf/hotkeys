// Example window-names fetches a list of all top-level client windows managed
// by the currently running window manager, and prints the name and size
// of each window.
//
// This example demonstrates how to use some aspects of the ewmh and icccm
// packages. It also shows how to use the xwindow package to find the
// geometry of a client window. In particular, finding the geometry is
// intelligent, as it includes the geometry of the decorations if they exist.
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

	err = keybind.KeyPressFun(
		func(X *xgbutil.XUtil, e xevent.KeyPressEvent) {
			err := exec.Command("/bin/sh", "-c", "import -window \"$(xdotool getwindowfocus)\" png:- | xclip -t image/png -selection c").Run()
			if err != nil {
				log.Printf(".. Screenshot failed with %s\n", err)
			}
			log.Printf(".. Screenshot done!")
		}).Connect(X, X.RootWin(), "Mod1-Print", true)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("Program initialized. Start pressing keys!")
	xevent.Main(X)
}
