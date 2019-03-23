package main

import (
	"crypto/tls"
	"fmt"

	irc "github.com/thoj/go-ircevent"
)

const channel = "#go-eventirc-test"
const serverssl = "irc.freenode.net:7000"

func main() {
	ircnick1 := "blatiblat"
	irccon := irc.IRC(ircnick1, "IRCTestSSL")
	irccon.VerboseCallbackHandler = true
	irccon.Debug = true
	irccon.UseTLS = true
	irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
	irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel) })
	irccon.AddCallback("366", func(e *irc.Event) {})
	err := irccon.Connect(serverssl)
	if err != nil {
		fmt.Printf("Err %s", err)
		return
	}
	irccon.Loop()
}

// package main

// import (
// 	"fyne.io/fyne/app"
// 	"fyne.io/fyne/widget"
// )

// func main() {
// 	app := app.New()

// 	w := app.NewWindow("Hello")
// 	w.SetContent(widget.NewVBox(
// 		widget.NewLabel("Hello Fyne!"),
// 		widget.NewButton("Quit", func() {
// 			app.Quit()
// 		}),
// 	))

// 	w.ShowAndRun()
// }
