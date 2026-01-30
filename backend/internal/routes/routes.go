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

	// Properties Page (The list of all properties)
	r.GET("/properties.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "properties.html", gin.H{
			"title": "Valley | Properties",
		})
	})

	// Single Property Details Page
	// This handles links like /property/1, /property/2, etc.
	r.GET("/property/:id", handlers.GetPropertyDetails)

	// Privacy Policy Page
	r.GET("/privacy.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "privacy.html", gin.H{
			"title": "Privacy Policy | Valley",
		})
	})

	// Terms of Service Page
	r.GET("/terms.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "terms.html", gin.H{
			"title": "Terms of Service | Valley",
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

	// --- NEW: 404 Error Handler ---
	// This catches any URL that doesn't match the ones above
	r.NoRoute(func(c *gin.Context) {
		c.HTML(http.StatusNotFound, "404.html", gin.H{
			"title": "Page Not Found | Valley",
		})
	})

	return r
}