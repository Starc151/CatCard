// fyne package -os windows -icon img/icon.png

package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/Starc151/CatCard/pkgs/shell"
)

func main() {
	app := app.New()
	icon, _ := fyne.LoadResourceFromPath("img/icon.png")
	app.SetIcon(icon)

	shell.NewShell(app)
	app.Run()
}