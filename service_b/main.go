package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Hola de servicio B",
		})
	})

	router.GET("/internal/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hola interno de Servicio B",
		})
	})

	router.GET("/openapi.json", func(c *gin.Context) {
		openapi := gin.H{
			"openapi": "3.1.0",
			"info": gin.H{
				"title":   "Servicio B",
				"version": "0.1.0",
			},
			"paths": gin.H{
				"/hello": gin.H{
					"get": gin.H{
						"tags":        []string{"Hello"},
						"summary":     "Hello",
						"operationId": "hello_hello_get",
						"responses": gin.H{
							"200": gin.H{
								"description": "Successful Response",
								"content": gin.H{
									"application/json": gin.H{
										"schema": gin.H{},
									},
								},
							},
						},
						"security": []gin.H{
							{
								"JWT Auth": []string{},
							},
						},
					},
				},
			},
			"components": gin.H{
				"securitySchemes": gin.H{
					"JWT Auth": gin.H{
						"type":        "http",
						"description": "Please use **/auth/login** endpoint to obtain a valid JWT",
						"scheme":      "bearer",
					},
				},
			},
		}
		c.JSON(http.StatusOK, openapi)
	})

	router.Run(":8080")
}
