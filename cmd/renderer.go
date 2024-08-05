package main

import "github.com/actanonv/excalidraw-local/ui"

func (a *Application) SetupRenderer() {
	a.Server.Renderer = ui.NewTemplateRenderer()
}
