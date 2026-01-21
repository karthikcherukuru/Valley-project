package handlers

import (
	"net/http"
	"strings"
	"valley/internal/data"

	"github.com/gin-gonic/gin"
)

func GetProperties(c *gin.Context) {
	c.JSON(http.StatusOK, data.Properties)
}

func GetPropertiesByCategory(c *gin.Context) {
	category := strings.ToLower(c.Param("category"))

	var filtered []interface{}
	for _, p := range data.Properties {
		if strings.ToLower(p.Category) == category {
			filtered = append(filtered, p)
		}
	}
	c.JSON(http.StatusOK, filtered)
}

func GetTestimonials(c *gin.Context) {
	c.JSON(http.StatusOK, data.Testimonials)
}
