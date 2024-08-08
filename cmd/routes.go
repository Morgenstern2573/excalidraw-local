package main

func (a *Application) SetupRouting() {
	a.Server.GET("/", a.Index)
	a.Server.POST("/new-drawing", a.NewDrawing)
	a.Server.POST("update-drawing-data", a.UpdateDrawingData)
	a.Server.POST("/new-collection", a.NewCollection)
	a.Server.GET("/drawing-list", a.DrawingList)
	a.Server.DELETE("/", a.DeleteDrawing)
	//TODO: replace with fig config
	a.Server.Static("/public", "./public")
}
