package ports

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/meraf00/agora-meet/agora-be/internal/users/app"
	"github.com/meraf00/agora-meet/agora-be/internal/users/app/query"
)

type HttpServer struct {
	app app.Application
}

func NewHttpServer(app app.Application, router *gin.Engine) HttpServer {
	server := HttpServer{
		app: app,
	}

	userRoutes := router.Group("/users")
	userRoutes.GET("/:id", server.GetProfileDetail)

	return server
}

func (h HttpServer) GetProfileDetail(ctx *gin.Context) {
	result, err := h.app.Queries.User.Handle(ctx.Request.Context(), query.User{
		Id: ctx.Param("id"),
	})

	if err != nil {
		fmt.Println(err)
		ctx.AbortWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.JSON(http.StatusOK, gin.H{"data": result})
}
