package window

import (
	"fmt"

	"fyne.io/fyne/app"
	"fyne.io/fyne/widget"
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
	// largeText := widget.NewMultiLineEntry()

	form := &widget.Form{
		OnCancel: func() {
			w.Close()
		},
		OnSubmit: func() {
			fmt.Println("Form submitted")
			fmt.Println("username:", username.Text)
			fmt.Println("server:", server.Text)
			fmt.Println("channel:", channel.Text)
			fmt.Println("Password:", password.Text)
			// fmt.Println("Message:", largeText.Text)
		},
	}
	form.Append("Username", username)
	form.Append("Server", server)
	form.Append("Channel", channel)

	form.Append("Password", password)
	w.SetContent(form)
	w.ShowAndRun()
}
