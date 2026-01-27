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
// GetPropertyDetails handles requests for a single property page
func GetPropertyDetails(c *gin.Context) {
    id := c.Param("id")

    // MOCK DATA: Simulating a database response
    // In the future, you will replace this with: property := database.GetProperty(id)
    var property gin.H
    
    if id == "1" {
        property = gin.H{
            "ID": "1",
            "Name": "Luxury Valley Villa",
            "Price": "₹ 4.5 Cr",
            "Location": "Sarjapur Road, Bangalore",
            "LocationURL": "https://www.google.com/maps", 
            "Config": "4 BHK Duplex",
            "Floor": "G + 1 (Top Floor)",
            "Status": "Ready to Move",
            "Possession": "Immediate",
            "Description": "Experience the pinnacle of luxury living in this sprawling 4 BHK duplex. Featuring Italian marble flooring, a private garden, and floor-to-ceiling windows offering panoramic views of the valley. Amenities include a clubhouse, infinity pool, and 24/7 security.",
            "Images": []string{
                "/images/valley-hero.mp4", // Placeholder using your existing video
                "https://images.unsplash.com/photo-1600596542815-6000255ade87?auto=format&fit=crop&w=800&q=80",
                "https://images.unsplash.com/photo-1600607687939-ce8a6c25118c?auto=format&fit=crop&w=800&q=80",
            },
        }
    } else {
        // Fallback data for any other ID so the page always works for testing
        property = gin.H{
            "ID": id,
            "Name": "Modern High-Rise Apartment",
            "Price": "₹ 1.2 Cr",
            "Location": "Whitefield, Bangalore",
            "LocationURL": "https://www.google.com/maps",
            "Config": "3 BHK",
            "Floor": "14th Floor",
            "Status": "Under Construction",
            "Possession": "Dec 2027",
            "Description": "A smart home for the modern professional. Located in the heart of the tech hub, this apartment offers smart-home automation, a rooftop zen garden, and coworking spaces within the building.",
            "Images": []string{
                 "https://images.unsplash.com/photo-1600585154340-be6161a56a0c?auto=format&fit=crop&w=800&q=80",
            },
        }
    }

    c.HTML(http.StatusOK, "property-details.html", gin.H{
        "title": property["Name"],
        "Prop": property,
    })
}