package handlers

import (
	"net/http"
	"strings"
	"valley/internal/data"
	"valley/internal/models" // <--- Added this import so we can use models.Property

	"github.com/gin-gonic/gin"
)

// GetProperties returns the list of all properties as JSON
func GetProperties(c *gin.Context) {
	c.JSON(http.StatusOK, data.Properties)
}

// GetPropertiesByCategory filters properties by category and returns JSON
func GetPropertiesByCategory(c *gin.Context) {
	category := strings.ToLower(c.Param("category"))

	var filtered []interface{}
	for _, p := range data.Properties {
		// Handle "land" special case or just general matching
		if strings.ToLower(p.Category) == category || (category == "land" && strings.Contains(strings.ToLower(p.Category), "land")) {
			filtered = append(filtered, p)
		} else if category == "all" {
			filtered = append(filtered, p)
		}
	}
	// If the loop finished but filtered is empty (and cat isn't 'all'), return empty list
	if filtered == nil {
		filtered = []interface{}{}
	}
	
	c.JSON(http.StatusOK, filtered)
}

// GetTestimonials returns the list of testimonials
func GetTestimonials(c *gin.Context) {
	c.JSON(http.StatusOK, data.Testimonials)
}

// GetPropertyDetails handles requests for a single property page (HTML)
func GetPropertyDetails(c *gin.Context) {
	id := c.Param("id")
	var selectedProp models.Property
	found := false

	// 1. Search for the property in our data list
	for _, p := range data.Properties {
		if p.ID == id {
			selectedProp = p
			found = true
			break
		}
	}

	// Fallback: If ID not found (or if data is empty), use the first property to prevent crash
	if !found && len(data.Properties) > 0 {
		selectedProp = data.Properties[0]
	}

	// 2. Define the "Extra Details" that aren't in your main list (Description, Config, etc.)
	// We map them here dynamically based on the ID.
	
	details := gin.H{
		"ID":          selectedProp.ID,
		"Name":        selectedProp.Title,
		"Price":       selectedProp.Price,
		"Location":    selectedProp.Location,
		"LocationURL": "https://www.google.com/maps/search/?api=1&query=" + selectedProp.Location, // Dynamic Map Link
		"Category":    selectedProp.Category,
		// Default values for any property
		"Config":      "Standard Configuration",
		"Floor":       "Ground Floor",
		"Status":      "Ready to Move",
		"Possession":  "Immediate",
		"Description": "A premium property opportunity located in " + selectedProp.Location + ", offering excellent returns and strategic value.",
		"Images":      []string{selectedProp.Image},
	}

	// 3. Custom details for the 4 specific Demo IDs
	switch selectedProp.ID {
	case "1": // Tech Park (Corporate)
		details["Config"] = "20,000 sqft Floor Plate"
		details["Floor"] = "Floors 1–15"
		details["Status"] = "Leased (Yield Generative)"
		details["Description"] = "Grade A office space in the heart of Hitech City. LEED Gold certified building with 100% power backup, high-speed elevators, and centralized air conditioning. Ideal for MNC headquarters."
		// Adding a second image for the slideshow
		details["Images"] = []string{selectedProp.Image, "https://images.unsplash.com/photo-1497215728101-856f4ea42174?auto=format&fit=crop&w=800&q=80"}
	
	case "2": // Logistics (Warehousing)
		details["Config"] = "1.5 Lakh sqft Warehouse"
		details["Status"] = "Operational"
		details["Floor"] = "Ground Level (12m Height)"
		details["Description"] = "State-of-the-art warehousing facility with 12m clear height, FM2 flooring, and 10 docking bays. Located on the main highway corridor for easy logistics access."
	
	case "3": // Penthouse (Residential)
		details["Config"] = "5 BHK Sky Villa"
		details["Floor"] = "24th (Top Floor)"
		details["Description"] = "Ultra-luxury penthouse with a private pool and 360-degree views of the Bangalore skyline. Features Italian marble flooring, home automation, and dedicated elevator access."
		details["Images"] = []string{selectedProp.Image, "https://images.unsplash.com/photo-1600607687939-ce8a6c25118c?auto=format&fit=crop&w=800&q=80"}

	case "4": // SEZ Land (Land)
		details["Config"] = "SEZ Notified Land"
		details["Status"] = "Clear Title"
		details["Floor"] = "N/A"
		details["Description"] = "A massive 400-acre contiguous land parcel suitable for setting up a large-scale IT SEZ or manufacturing unit. Fully compliant with government regulations and ready for development."
	}

	c.HTML(http.StatusOK, "property-details.html", gin.H{
		"title": selectedProp.Title,
		"Prop":  details,
	})
}