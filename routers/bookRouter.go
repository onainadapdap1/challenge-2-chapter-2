package routers

import (
	"simple_rest_api_book/controllers"

	"github.com/gin-gonic/gin"
)


func StartServer() *gin.Engine {
	router := gin.Default()

	router.POST("/books", controllers.CreateBook)

	router.PUT("/books/:bookID", controllers.UpdateBook)

	router.GET("/books", controllers.GetAllCar)

	router.GET("/books/:bookID", controllers.GetBookByID)

	return router
}