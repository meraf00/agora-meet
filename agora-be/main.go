package main

import (
	"context"
	"fmt"

	"github.com/gin-gonic/gin"

	"github.com/meraf00/agora-meet/agora-be/internal/users/app"
	"github.com/meraf00/agora-meet/agora-be/internal/users/ports"
	"github.com/meraf00/agora-meet/agora-be/shared/config"
)

func main() {
	config.LoadConfig()

	router := gin.Default()

	ctx := context.Background()

	application := app.NewApplication(ctx)

	ports.NewHttpServer(application, router)

	router.Run(fmt.Sprintf(":%d", config.AppConfig.Port))
}
