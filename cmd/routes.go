package main

func (a *Application) SetupRouting() {
	a.Server.GET("/:sceneID", a.Index)

	//TODO: replace with fig config
	a.Server.Static("/public", "./public")
}
