`active-window-screenshot` is a simple screenshot tool for X11.

It captures the active focused window and copy the PNG to X11 clipboard, to be
pasted somewhere else. It doesn't generate any temporary image.

It consists of a simple go wrapper around `xdotool`, `xclip` and `imagemagick`.

Basically, it calls this one liner:
			
```sh
import -window "$(xdotool getwindowfocus)" png:- | xclip -t image/png -selection c
```

so I guess you could also bind this in your WM shortcuts or run this program
and press `Alt-PrintScreen` to trigger the capture.
