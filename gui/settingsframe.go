package gui

import (
	"EnCloud/go-gtk/gtk"
)

func SettingsFrame() gtk.IWidget {

	page := gtk.NewFrame("Settings")
	vbox := gtk.NewHBox(false, 1)

	table := gtk.NewTable(2, 2, true)

	for i := uint(0); i < 3; i++ {
		label := gtk.NewLabel("Setting Name:")

		table.Attach(label, 0, 1, i, i+1, gtk.FILL, gtk.FILL, 5, 5)
	}

	for i := uint(0); i < 3; i++ {
		button := gtk.NewButtonWithLabel("Button")

		button.Clicked(func() {
			// Do stuff.
		})

		table.Attach(button, 1, 2, i, i+1, gtk.FILL, gtk.FILL, 5, 5)
	}

	vbox.Add(table)
	page.Add(vbox)

	return page
}
