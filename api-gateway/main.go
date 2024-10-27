package main

import (
	"log"
	"net/http/httputil"
	"net/url"

	"github.com/gin-gonic/gin"
)

// Helper function to create a reverse proxy
func createProxy(target string) gin.HandlerFunc {
	targetURL, err := url.Parse(target)
	if err != nil {
		log.Fatalf("Could not parse target URL: %v", err)
	}
	proxy := httputil.NewSingleHostReverseProxy(targetURL)

	return func(c *gin.Context) {
		proxy.ServeHTTP(c.Writer, c.Request)
	}
}

func main() {
	r := gin.Default()

	// Define service URLs
	userServiceURL := "http://localhost:8081"
	orderServiceURL := "http://localhost:8082"

	// Define routes for user service
	r.POST("/users", createProxy(userServiceURL))
	r.GET("/users", createProxy(userServiceURL))
	r.GET("/users/:id", createProxy(userServiceURL))

	// Define routes for order service
	r.POST("/orders", createProxy(orderServiceURL))
	r.GET("/orders", createProxy(orderServiceURL))
	r.GET("/orders/:id", createProxy(orderServiceURL))

	// Run the API Gateway on port 8080
	if err := r.Run(":8080"); err != nil {
		log.Fatalf("Could not start API Gateway: %v", err)
	}
}
