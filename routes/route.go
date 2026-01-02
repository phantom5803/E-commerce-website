package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/phantom5803/E-commerce-website/controllers"
)
func UserRoutes(c *gin.Engine) {
	c.POST("/users/signup",controllers.Signup())
	c.POST("/users/login",controllers.Login())
	c.GET("/admin/products",controllers.ProductViewAdmin())
	c.POST("/admin/addproduct",controllers.AddProduct())
	c.GET("/users/search",controllers.SearchProduct())
}

