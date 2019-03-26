package window

import (
	"crypto/tls"
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
	irc "github.com/thoj/go-ircevent"
)

func New() {
	// create a new application
	app := app.New()
	// create a new window
	w := app.NewWindow("Form")

	username := widget.NewEntry()
	username.SetPlaceHolder("cooldude123")
	server := widget.NewEntry()
	server.SetPlaceHolder("test@example.com")
	channel := widget.NewEntry()
	channel.SetPlaceHolder("#golang-rocks")
	password := widget.NewPasswordEntry()

	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: func() {
			// fmt.Println("Form submitted")
			// fmt.Println("username:", username.Text)
			// fmt.Println("server:", server.Text)
			// fmt.Println("channel:", channel.Text)
			// fmt.Println("Password:", password.Text)

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

		},
	}
	form.Append("Username", username)
	form.Append("Password", password)
	form.Append("Server", server)
	form.Append("Channel", channel)

	w.SetContent(form)
	w.ShowAndRun()
}
