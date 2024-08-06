package main

func (a *Application) SetupRouting() {
	a.Server.GET("/", a.Index)
	a.Server.POST("/new-scene", a.NewScene)

	//TODO: replace with fig config
	a.Server.Static("/public", "./public")
}
