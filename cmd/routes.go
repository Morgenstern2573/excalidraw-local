package main

func (a *Application) SetupRouting() {
	a.Server.GET("/", a.Index)
	a.Server.POST("/new-scene", a.NewScene)
	a.Server.POST("update-scene-data", a.UpdateSceneData)
	//TODO: replace with fig config
	a.Server.Static("/public", "./public")
}
