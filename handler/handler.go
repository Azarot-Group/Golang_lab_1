package handler

// git framework
import "github.com/gin-gonic/gin"

type Handler struct {
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	api := router.Group("/api")
	{
		api.GET("/time", h.getTime)
		api.GET("/", h.greet)
	}

	return router
}
