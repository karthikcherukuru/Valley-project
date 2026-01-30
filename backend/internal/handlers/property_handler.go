package handlers

import (
    "net/http"
    "strings"
    "valley/internal/data"
    "valley/internal/models"

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
        pCat := strings.ToLower(p.Category)
        // Smart Filter Logic
        if pCat == category || 
           (category == "land" && strings.Contains(pCat, "land")) ||
           (category == "residential" && (pCat == "residential" || pCat == "villa")) { // Include Villas in Residential
            filtered = append(filtered, p)
        } else if category == "all" {
            filtered = append(filtered, p)
        }
    }
    if filtered == nil {
        filtered = []interface{}{}
    }
    
    c.JSON(http.StatusOK, filtered)
}

// GetTestimonials returns the list of testimonials
func GetTestimonials(c *gin.Context) {
    c.JSON(http.StatusOK, data.Testimonials)
}

// GetPropertyDetails handles requests for a single property page
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

    // Fallback if ID not found
    if !found && len(data.Properties) > 0 {
        selectedProp = data.Properties[0]
    }

    // 2. Define standard details
    details := gin.H{
        "ID":          selectedProp.ID,
        "Name":        selectedProp.Title,
        "Price":       selectedProp.Price,
        "Location":    selectedProp.Location,
        "LocationURL": "http://googleusercontent.com/maps.google.com/?q=" + selectedProp.Location,
        "Category":    selectedProp.Category,
        "Images":      []string{selectedProp.Image},
        // Defaults
        "Config":      "Standard Config",
        "Floor":       "G + Floors",
        "Status":      "Ready to Move",
        "Possession":  "Immediate",
    }

    // 3. Custom Details for All 4 Properties
    switch selectedProp.ID {
    case "1": // Purva Silversky
        details["Config"] = "3 & 4 BHK Premium"
        details["Status"] = "Newly Launched"
        details["Possession"] = "2028"
        details["Floor"] = "G + 20 Floors"
        details["Description"] = `
            <strong>About The Project:</strong> Purva Silversky represents a new benchmark in premium residential living in South East Bangalore. Located in Hebbagodi near Electronic City Phase 2, it features over 70% open green spaces and low-density planning.<br><br>
            <strong>Why Choose This?</strong> Intelligent layouts maximize ventilation and light. Amenities include a grand clubhouse, swimming pool, multiple sports courts, and a party hall.<br><br>
            <strong>Connectivity:</strong> Minutes from Electronic City Phase 2 and Biocon Campus. Excellent access to Hosur Main Road and the upcoming Metro Phase 2 expansion.
        `
        details["Images"] = []string{selectedProp.Image, "https://images.unsplash.com/photo-1574362848149-11496d93a7c7?auto=format&fit=crop&w=800&q=80"}

    case "2": // Sobha Whitefield
        details["Config"] = "2 & 3 BHK Luxury"
        details["Status"] = "Upcoming Project"
        details["Possession"] = "2029"
        details["Floor"] = "Luxury Towers"
        details["Description"] = `
            <strong>Luxury Living in East Bangalore:</strong> Sobha Whitefield apartments redefine premium living with a perfect blend of luxury and connectivity. Designed for homebuyers seeking quality construction and strong appreciation potential.<br><br>
            <strong>Key Highlights:</strong> Close proximity to ITPL, EPIP Zone, and Whitefield Metro Station. Features include a world-class clubhouse, co-working spaces, and landscaped gardens.<br><br>
            <strong>Investment Potential:</strong> High rental demand from IT professionals and strong brand trust of Sobha Limited make this a top-tier investment choice.
        `
        details["Images"] = []string{selectedProp.Image, "https://images.unsplash.com/photo-1502005229766-93976a171e25?auto=format&fit=crop&w=800&q=80"}

    case "3": // Sobha Lifestyle Legacy (Villa)
        details["Config"] = "4 BHK Luxury Villas (6,100+ sq.ft)"
        details["Status"] = "Ready / Legacy"
        details["Possession"] = "Immediate"
        details["Floor"] = "Triplex Villa"
        details["Description"] = `
            <strong>Bengaluru’s Next Luxury Hotspot:</strong> Located on IVC Road, Sobha Lifestyle Legacy is spread across 55 acres of lush greenery. These 4 BHK villas (starting 6,100 sq.ft) are a masterpiece of privacy and precision engineering.<br><br>
            <strong>World-Class Amenities:</strong> Temperature-controlled pool, spa, sauna, and a grand clubhouse. Enveloped by tree-lined avenues and 24/7 advanced security.<br><br>
            <strong>Location Advantage:</strong> Just 15 mins from the International Airport. Surrounded by top schools (Canadian International) and upcoming infrastructure like the ITIR and STRR.
        `
        details["Images"] = []string{
            selectedProp.Image,
            "https://images.unsplash.com/photo-1600596542815-6000255ade87?auto=format&fit=crop&w=800&q=80", // Villa Interior
            "https://images.unsplash.com/photo-1564013799919-ab600027ffc6?auto=format&fit=crop&w=800&q=80", // Garden
        }

    case "4": // Chaithanya Oak Ville (Villa)
        details["Config"] = "Premium Luxury Villas"
        details["Status"] = "Ready to Move"
        details["Possession"] = "Immediate"
        details["Floor"] = "Duplex + Terrace"
        details["Description"] = `
            <strong>Tropical Charm in Whitefield:</strong> Chaithanya Oak Ville offers a harmonious blend of luxury and nature. Known for its double-volume living areas, wooden decks, and central courtyards with water features.<br><br>
            <strong>Resort-Style Living:</strong> Amenities include landscaped gardens, swimming pools, and recreation spaces that offer a serene escape from the city buzz while remaining in the heart of the IT hub.<br><br>
            <strong>Investment Value:</strong> Located in Whitefield's most dynamic neighborhood, ensuring consistent property value appreciation and high demand from expats and top executives.
        `
        details["Images"] = []string{
            selectedProp.Image,
            "https://images.unsplash.com/photo-1600210492486-724fe5c67fb0?auto=format&fit=crop&w=800&q=80", // Modern Deck
            "https://images.unsplash.com/photo-1512917774080-9991f1c4c750?auto=format&fit=crop&w=800&q=80", // Pool
        }
    }

    c.HTML(http.StatusOK, "property-details.html", gin.H{
        "title": selectedProp.Title,
        "Prop":  details,
    })
}