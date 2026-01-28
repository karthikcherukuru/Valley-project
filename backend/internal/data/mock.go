package data

import "valley/internal/models"

// Only 4 Properties for the Demo
var Properties = []models.Property{
    {ID: "1", Title: "Tech Park One", Category: "Corporate", Price: "₹85/sqft", Location: "Hitech City, Hyderabad", Yield: "Rental", Image: "https://images.unsplash.com/photo-1497366216548-37526070297c?auto=format&fit=crop&w=800&q=80"},
    {ID: "2", Title: "Logistics Hub North", Category: "Warehousing", Price: "₹45 Cr", Location: "Bhiwandi, Mumbai", Yield: "8.5%", Image: "https://images.unsplash.com/photo-1586528116311-ad8dd3c8310d?auto=format&fit=crop&w=800&q=80"},
    {ID: "3", Title: "Skyline Penthouse", Category: "Residential", Price: "₹5.5 Cr", Location: "Whitefield, Bangalore", Yield: "Sale", Image: "https://images.unsplash.com/photo-1512917774080-9991f1c4c750?auto=format&fit=crop&w=800&q=80"},
    {ID: "4", Title: "400 Acre SEZ Land", Category: "Land", Price: "₹150 Cr", Location: "Noida, NCR", Yield: "Sale", Image: "https://images.unsplash.com/photo-1500382017468-9049fed747ef?auto=format&fit=crop&w=800&q=80"},
}

var Testimonials = []models.Testimonial{
    {ID: "1", Name: "Vivek Chowhan", Role: "Director", Company: "Treebo Hotels", Content: "The team is thorough in Real Estate and they cover the entire spectrum. Extremely helpful and supportive.", Avatar: "https://randomuser.me/api/portraits/men/32.jpg"},
    {ID: "2", Name: "Arvind S.", Role: "CEO", Company: "Mahindra Happinest", Content: "Went out of their way to deeply understand our requirements and come up with a tailor-made solution.", Avatar: "https://randomuser.me/api/portraits/men/45.jpg"},
    {ID: "3", Name: "Shyamli Mishra", Role: "Manager", Company: "Freight Tiger", Content: "I cannot thank you enough for helping us find our office space in Bangalore. Almost as good as being tailor made.", Avatar: "https://randomuser.me/api/portraits/women/44.jpg"},
}