package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

// --- STRUCTS ---

type Property struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Category string `json:"category"` // Corporate, Warehousing, Residential, Land
	Price    string `json:"price"`
	Location string `json:"location"`
	Yield    string `json:"yield"`
	Image    string `json:"image"`
}

type Testimonial struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
	Role    string `json:"role"`
	Company string `json:"company"`
	Content string `json:"content"`
	Avatar  string `json:"avatar"`
}

// --- MOCK DATA ---

var properties = []Property{
	// --- CORPORATE ---
	{ID: "1", Title: "Tech Park One", Category: "Corporate", Price: "₹85/sqft", Location: "Hitech City, Hyderabad", Yield: "Rental", Image: "https://images.unsplash.com/photo-1497366216548-37526070297c?auto=format&fit=crop&w=800&q=80"},
	{ID: "2", Title: "Prestige Trade Tower", Category: "Corporate", Price: "₹120/sqft", Location: "CBD, Bangalore", Yield: "Rental", Image: "https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=800&q=80"},
	{ID: "3", Title: "DLF Cyber City Hub", Category: "Corporate", Price: "₹140/sqft", Location: "Gurgaon, NCR", Yield: "Rental", Image: "https://images.unsplash.com/photo-1554469384-e58fac16e23a?auto=format&fit=crop&w=800&q=80"},
	{ID: "4", Title: "Financial District Summit", Category: "Corporate", Price: "₹90/sqft", Location: "Gachibowli, Hyderabad", Yield: "Rental", Image: "https://images.unsplash.com/photo-1497215728101-856f4ea42174?auto=format&fit=crop&w=800&q=80"},
	
	// --- WAREHOUSING ---
	{ID: "5", Title: "Logistics Hub North", Category: "Warehousing", Price: "₹45 Cr", Location: "Bhiwandi, Mumbai", Yield: "8.5%", Image: "https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?auto=format&fit=crop&w=800&q=80"},
	{ID: "6", Title: "Chennai Industrial Park", Category: "Warehousing", Price: "₹32 Cr", Location: "Oragadam, Chennai", Yield: "9.0%", Image: "https://images.unsplash.com/photo-1553413077-190dd305871c?auto=format&fit=crop&w=800&q=80"},
	{ID: "8", Title: "Kolkata Cold Chain", Category: "Warehousing", Price: "₹18 Cr", Location: "Dankuni, Kolkata", Yield: "9.5%", Image: "https://images.unsplash.com/photo-1587293852726-70cdb56c2866?auto=format&fit=crop&w=800&q=80"},

	// --- RESIDENTIAL ---
	{ID: "9", Title: "Skyline Penthouse", Category: "Residential", Price: "₹5.5 Cr", Location: "Whitefield, Bangalore", Yield: "Sale", Image: "https://images.unsplash.com/photo-1512917774080-9991f1c4c750?auto=format&fit=crop&w=800&q=80"},
	{ID: "10", Title: "Eon Waterfront", Category: "Residential", Price: "₹3.8 Cr", Location: "Kharadi, Pune", Yield: "Sale", Image: "https://images.unsplash.com/photo-1493809842364-78817add7ffb?auto=format&fit=crop&w=800&q=80"},
	{ID: "11", Title: "Jubilee Hills Villa", Category: "Residential", Price: "₹12 Cr", Location: "Jubilee Hills, Hyderabad", Yield: "Sale", Image: "https://images.unsplash.com/photo-1600585154340-be6161a56a0c?auto=format&fit=crop&w=800&q=80"},
	{ID: "12", Title: "Worli Sea Face Apt", Category: "Residential", Price: "₹15 Cr", Location: "Worli, Mumbai", Yield: "Sale", Image: "https://images.unsplash.com/photo-1600607687939-ce8a6c25118c?auto=format&fit=crop&w=800&q=80"},

	// --- LAND ---
	{ID: "13", Title: "400 Acre SEZ Land", Category: "Land", Price: "₹150 Cr", Location: "Noida, NCR", Yield: "Sale", Image: "https://images.unsplash.com/photo-1500382017468-9049fed747ef?auto=format&fit=crop&w=800&q=80"},
	{ID: "14", Title: "Coastal Resort Plot", Category: "Land", Price: "₹12 Cr", Location: "Goa", Yield: "Sale", Image: "https://images.unsplash.com/photo-1613977257363-707ba9348227?auto=format&fit=crop&w=800&q=80"},
	{ID: "15", Title: "Organic Farmland", Category: "Land", Price: "₹2.5 Cr", Location: "Coorg, Karnataka", Yield: "Sale", Image: "https://images.unsplash.com/photo-1470723710355-95304d8aece4?auto=format&fit=crop&w=800&q=80"},
	{ID: "16", Title: "Industrial Plot Phase 1", Category: "Land", Price: "₹8 Cr", Location: "Sri City, Andhra Pradesh", Yield: "Sale", Image: "https://images.unsplash.com/photo-1590247813693-5541d1c609fd?auto=format&fit=crop&w=800&q=80"},
}

var testimonials = []Testimonial{
	{ID: "1", Name: "Vivek Chowhan", Role: "Director", Company: "Treebo Hotels", Content: "The team is thorough in Real Estate and they cover the entire spectrum. Extremely helpful and supportive.", Avatar: "https://randomuser.me/api/portraits/men/32.jpg"},
	{ID: "2", Name: "Arvind S.", Role: "CEO", Company: "Mahindra Happinest", Content: "Went out of their way to deeply understand our requirements and come up with a tailor-made solution.", Avatar: "https://randomuser.me/api/portraits/men/45.jpg"},
	{ID: "3", Name: "Shyamli Mishra", Role: "Manager", Company: "Freight Tiger", Content: "I cannot thank you enough for helping us find our office space in Bangalore. Almost as good as being tailor made.", Avatar: "https://randomuser.me/api/portraits/women/44.jpg"},
}

func main() {
	r := gin.Default()

	// 1. CORS Middleware
	r.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS, POST")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	// 2. Serve Static Assets (Images)
	r.Static("/images", "./images")

	// 3. Serve HTML Pages
	// This allows localhost:5000/ to load index.html
	r.GET("/", func(c *gin.Context) {
		c.File("./index.html")
	})

	// This allows direct access to the HTML files
	r.StaticFile("/index.html", "./index.html")
	r.StaticFile("/about.html", "./about.html")
	r.StaticFile("/properties.html", "./properties.html") // Coming next
	r.StaticFile("/services.html", "./services.html")     // Coming soon
	r.StaticFile("/contact.html", "./contact.html")       // Coming soon

	// 4. API Routes
	r.GET("/api/properties", func(c *gin.Context) {
		c.JSON(http.StatusOK, properties)
	})

	r.GET("/api/properties/filter/:category", func(c *gin.Context) {
		category := c.Param("category")
		var filtered []Property
		for _, p := range properties {
			if strings.EqualFold(p.Category, category) {
				filtered = append(filtered, p)
			}
		}
		c.JSON(http.StatusOK, filtered)
	})

	r.GET("/api/testimonials", func(c *gin.Context) {
		c.JSON(http.StatusOK, testimonials)
	})

	// Start Server
	r.Run(":5000")
}