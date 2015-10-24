`active-window-screenshot` is a simple screenshot tool for X11.

It captures the active focused window, the root window or a user defined area
and copy the PNG to X11 clipboard, to be pasted somewhere else and also creates
the PNG in `~/tmp/`.

- `mod1-Print`: active window
- `mod1-shift-Print`: user defined area
- `mod1-control-shift-Print`: root window

It consists of a simple go wrapper around `xdotool`, `xclip` and `imagemagick`.

Basically, it calls this one liner:
			
```sh
import -window "$(xdotool getwindowfocus)" png:- | xclip -t image/png -selection c
```

so I guess you could also bind this in your WM shortcuts or run this program
and press `Mod1-Print` to trigger the capture. However, `dwm` making it a real
PITA to spawn complex commands, here it goes.
