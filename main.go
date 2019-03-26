package main

import (
	"crypto/tls"
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	irc "github.com/thoj/go-ircevent"
)

// const channel = "test"
// const serverssl = "irc.freenode.net:7000"

func main() {
	app := app.New()
	// create a new window
	w := app.NewWindow("Form")

	username := widget.NewEntry()
	username.SetPlaceHolder("cooldude123999")
	server := widget.NewEntry()
	server.SetPlaceHolder("irc.freenode.net:7000")
	channel := widget.NewEntry()
	channel.SetPlaceHolder("#golang-rocks")
	password := widget.NewPasswordEntry()

	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: func() {
			ircnick1 := username.Text
			irccon := irc.IRC(ircnick1, "IRCTestSSL")
			irccon.VerboseCallbackHandler = true
			irccon.Debug = true
			irccon.UseTLS = true
			irccon.TLSConfig = &tls.Config{InsecureSkipVerify: true}
			irccon.AddCallback("001", func(e *irc.Event) { irccon.Join(channel.Text) })
			irccon.AddCallback("366", func(e *irc.Event) {})
			err := irccon.Connect(server.Text)
			if err != nil {
				fmt.Printf("Err %s", err)
				return
			}
			irccon.Loop()
		},
	}
	form.Append("Username", username)
	form.Append("Password", password)
	form.Append("Server", server)
	form.Append("Channel", channel)

	w.SetContent(form)
	w.ShowAndRun()
	// window.New()
	// irc.New()
}
