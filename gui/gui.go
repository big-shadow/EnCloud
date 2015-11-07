package gui

import (
	"EnCloud/go-gtk/gtk"
	"os"
)

func BuildForm() {
	gtk.Init(&os.Args)
	window := gtk.NewWindow(gtk.WINDOW_TOPLEVEL)
	window.SetTitle("EnCloud - Cloud Encryption")
	window.Connect("destroy", gtk.MainQuit)

	notebook := gtk.NewNotebook()

	notebook.AppendPage(WelcomeFrame(), gtk.NewLabel("Welcome"))
	notebook.AppendPage(SettingsFrame(), gtk.NewLabel("Settings"))

	window.Add(notebook)
	window.SetSizeRequest(500, 350)
	window.ShowAll()

	gtk.Main()
}
