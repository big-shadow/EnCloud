package gui

import (
	"EnCloud/go-gtk/gtk"
	"EnCloud/keygen"
	"EnCloud/validator"
)

func WelcomeFrame() gtk.IWidget {

	page := gtk.NewFrame("Welcome")
	vbox := gtk.NewHBox(false, 1)
	table := gtk.NewTable(4, 1, true)

	txtMessage := gtk.NewLabel("We need some info!")
	table.Attach(txtMessage, 1, 2, 0, 1, gtk.FILL, gtk.FILL, 0, 0)

	txtEmail := gtk.NewEntry()
	txtEmail.SetEditable(true)

	// table.Attach: Left, right, top, bottom
	table.Attach(gtk.NewLabel("Email Address"), 1, 2, 2, 3, gtk.FILL, gtk.FILL, 0, 0)
	table.Attach(txtEmail, 1, 2, 3, 4, gtk.FILL, gtk.FILL, 5, 0)

	txtPassword := gtk.NewEntry()
	txtPassword.SetEditable(true)

	table.Attach(gtk.NewLabel("Password"), 1, 2, 4, 5, gtk.FILL, gtk.FILL, 0, 0)
	table.Attach(txtPassword, 1, 2, 5, 6, gtk.FILL, gtk.FILL, 5, 0)

	button := gtk.NewButtonWithLabel("Go!")

	button.Clicked(func() {
		if validator.ValidateEmailAddress(txtEmail.GetText()) {
			if keygen.Generate() {
				txtMessage.SetText("Success! :)")
			} else {
				txtMessage.SetText("No Dice. :(")
			}
		} else {
			txtMessage.SetText("Invalid E-mail Address.")
		}
	})

	table.Attach(button, 1, 2, 7, 8, gtk.FILL, gtk.FILL, 5, 0)

	vbox.Add(table)
	page.Add(vbox)

	return page
}
