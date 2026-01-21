package routes

import (
	"net/http"
	"valley/internal/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// 1. Load HTML Templates
	r.LoadHTMLGlob("templates/*.html")

	// 2. Serve Static Assets (CSS, JS, Images)
	r.Static("/assets", "./assets")
	r.Static("/images", "./images")

	// 3. PAGE ROUTES

	// Home Page
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.html", gin.H{
			"title": "Valley | Home",
		})
	})

	// About Us Page
	r.GET("/about.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "about.html", gin.H{
			"title": "Valley | About Us",
		})
	})

	// Contact Page
	r.GET("/contact.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "contact.html", gin.H{
			"title": "Valley | Contact Us",
		})
	})

	// Services Page
	r.GET("/services.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "services.html", gin.H{
			"title": "Valley | Our Services",
		})
	})

	// Properties Page
	r.GET("/properties.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "properties.html", gin.H{
			"title": "Valley | Properties",
		})
	})

	// 4. API Routes
	api := r.Group("/api")
	{
		api.GET("/properties", handlers.GetProperties)
		api.GET("/properties/filter/:category", handlers.GetPropertiesByCategory)
		api.GET("/testimonials", handlers.GetTestimonials)
		api.POST("/contact", handlers.SubmitContact)
	}

	return r
}
