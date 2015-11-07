package gui

import (
	"EnCloud/go-gtk/gtk"
)

func WelcomeFrame() gtk.IWidget {

	page := gtk.NewFrame("Welcome")
	vbox := gtk.NewHBox(false, 1)
	text := gtk.NewLabel("Hey! This a welcome page. This could become a login page.\n\n" +
		"Users could enter an email and password to produce keys!")

	vbox.Add(text)
	page.Add(vbox)

	return page
}
