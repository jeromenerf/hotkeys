`hotkeys.go` is a basic hotkeys tool for X11, born from the need to take
screenshots quickly while dwm made it annoying to tee and pipe, doesn't like
shell aliases, etc.


## installation

Providing a sane go workspace:

`go get github.com/jeromenerf/hotkeys`

## usage

Just run `hotkeys` and leave it watch your config file.

Beware, this quick hack uses `inotify` since the cross platform `fsnotify`
package has been left behind for a while now. Don't expect it to work on
anything else than linux :/

## configuration

The configuration file lives at `~/.config/hotkeys.conf.json`:

```json

[
  {
    "key":  "Print",
    "desc": "Take a screenshot of the active window",
    "cmd":  "import -window \"$(xdotool getwindowfocus)\" png:- | tee ~/tmp/screenshot-$(date +'%Y-%m-%d-%T').png | xclip -t image/png -selection c"
  },
  {
    "key":  "Mod1-Print",
    "desc": "Take a screenshot of a user selected area",
    "cmd":  "import png:- | tee ~/tmp/screenshot-$(date +'%Y-%m-%d-%T').png | xclip -t image/png -selection c"
  },
  {
    "key":  "Mod1-Shift-Print",
    "desc": "Take a screenshot of the root window",
    "cmd":  "import -window root png:-  | tee ~/tmp/screenshot-$(date +'%Y-%m-%d-%T').png | xclip -t image/png -selection c"
  },
  {
    "key": "XF86AudioPlay",
    "desc": "Play",
    "cmd": "for player in vlc spotify; do dbus-send --print-reply --dest=org.mpris.MediaPlayer2.$player /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.PlayPause; done"
  },
  {
    "key": "XF86AudioStop",
    "desc": "Stop",
    "cmd": "for player in vlc spotify; do dbus-send --print-reply --dest=org.mpris.MediaPlayer2.$player /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.Stop; done"
  },
  {
    "key": "XF86AudioNext",
    "desc": "Next",
    "cmd": "for player in vlc spotify; do dbus-send --print-reply --dest=org.mpris.MediaPlayer2.$player /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.Next; done"
  },
  {
    "key": "XF86AudioPrev",
    "desc": "Previous",
    "cmd": "for player in vlc spotify; do dbus-send --print-reply --dest=org.mpris.MediaPlayer2.$player /org/mpris/MediaPlayer2 org.mpris.MediaPlayer2.Player.Previous; done"
  },
  {
    "key": "XF86Display",
    "desc": "Switch displays auto magically",
    "cmd": "xrandr --auto"
  },
  {
    "key": "XF86ScreenSaver",
    "desc": "Lock screen",
    "cmd": "slock"
  }
]
```

This sample uses lots of external tools :

- `xdotool` to get various informations for x11 clients
- (dbus) to send messages to audio players
- `import` from `imagemagick`
- `xclip` to copy paste from X11 clipboard
- `slock` to lock screen
