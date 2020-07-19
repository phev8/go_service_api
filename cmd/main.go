package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	gc "github.com/phev8/go_service_api/pkg/grpc/clients"
	v1 "github.com/phev8/go_service_api/pkg/http/v1"
	"github.com/phev8/go_service_api/pkg/types"
)

// Conf holds all static configuration information

var grpcClients *types.APIClients

func init() {
	grpcClients = &types.APIClients{}
}

func healthCheckHandle(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"status": "ok"})
}

func main() {
	clientA, closeA := gc.ConnectToAService("localhost:4301")
	defer closeA()
	clientB, closeB := gc.ConnectToBService("localhost:4302")
	defer closeB()

	grpcClients.ServiceA = clientA
	grpcClients.ServiceB = clientB

	// Start webserver
	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowAllOrigins:  true,
		AllowMethods:     []string{"POST", "GET", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type", "Content-Length"},
		ExposeHeaders:    []string{"Authorization", "Content-Type", "Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/", healthCheckHandle)
	v1Root := router.Group("/v1")

	v1APIHandlers := v1.NewHTTPHandler(grpcClients)
	v1APIHandlers.AddAPI(v1Root)

	port := "4300"
	log.Printf("gateway listening on port %s", port)
	log.Fatal(router.Run(":" + port))
}
