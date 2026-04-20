package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
)

//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	app := NewApp()

	trayMenu := menu.NewMenu()
	trayMenu.AddText("Quit", nil, func(_ *menu.CallbackData) {
		app.Quit()
	})

	err := wails.Run(&options.App{
			Title:  "translate-64",
			Width:  700,
			Height: 775,
			AssetServer: &assetserver.Options{
				Assets: assets,
			},
			OnStartup: app.startup,
			Bind: []interface{}{
				app,
			},
			// This handles the icon in the window title bar and taskbar/panel
			Linux: &linux.Options{
				Icon: icon,
				WindowIsTranslucent: false,
			},
		})

	if err != nil {
		println("Error:", err.Error())
	}
}