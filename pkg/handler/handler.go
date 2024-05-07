package handler

import (
	"github.com/ctrlVadim/go-chat/pkg/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	services *service.Service
}

func NewHandler(services *service.Service) *Handler {
	return &Handler{services: services}
}

func (h *Handler) InitRoutes() *gin.Engine {
	router := gin.New()

	auth := router.Group("/auth")
	{
		auth.POST("/sign-up", h.signUp)
		auth.POST("/sign-in", h.signIn)
	}

	// api := router.Group("/api", h.userIdentity)
	// {
	// 	chats := api.Group("/chats")
	// 	{
	// 		api.GET("/", h.getChats)
	// 		api.POST("/:id", h.storeChat)
	// 		api.PUT("/:id", h.updateChat)
	// 		api.DELETE("/:id", h.deleteChat)
	// 	}

	// }

	return router
}
