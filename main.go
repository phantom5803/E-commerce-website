package ecommerceweb


import(
	"github.com/phantom5803/E-commerce-website/controllers"
	"github.com/phantom5803/E-commerce-website/routes"
	"github.com/gin-gonic/gin"
	"github.com/phantom5803/E-commerce-website/database"
	"github.com/phantom5803/E-commerce-website/models"
	"github.com/phantom5803/E-commerce-website/middleware"
	"os"
	"log"
)

func main(){

	port:= os.Getenv("Port")

	if port == " "{
		port = "8080"
	}
	
	router := gin.New()
	router.Use(gin.Logger())
	router.Use(middleware.Authentication())

	router.GET("/addtocart",controllers.AddToCart())
	router.GET("/removeitem",controllers.RemoveItem())
	router.GET("/cartcheckout",controllers.CartCheckout())
	router.GET("/instantorder",controllers.InstantOrder())


	log.Fatal(router.Run(":" + port))



}