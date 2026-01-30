package data

import "valley/internal/models"

// Updated Property List (4 Real Properties)
var Properties = []models.Property{
    // Property 1: Apartment
    {
        ID:       "1",
        Title:    "Purva Silversky",
        Category: "Residential",
        Price:    "Contact for Price",
        Location: "Hebbagodi, Electronic City",
        Yield:    "Sale",
        Image:    "https://images.unsplash.com/photo-1545324418-cc1a3fa10c00?auto=format&fit=crop&w=800&q=80",
    },
    // Property 2: Apartment
    {
        ID:       "2",
        Title:    "Sobha Whitefield",
        Category: "Residential",
        Price:    "₹39.1 Lakhs Onwards",
        Location: "Whitefield, Bangalore",
        Yield:    "Investment",
        Image:    "https://images.unsplash.com/photo-1486406146926-c627a92ad1ab?auto=format&fit=crop&w=800&q=80",
    },
    // Property 3: Villa (New)
    {
        ID:       "3",
        Title:    "Sobha Lifestyle Legacy",
        Category: "Villa",
        Price:    "₹6.5 Cr Onwards",
        Location: "Devanahalli, North Bangalore",
        Yield:    "Luxury",
        Image:    "https://images.unsplash.com/photo-1613977257363-707ba9348227?auto=format&fit=crop&w=800&q=80", // Luxury Villa Image
    },
    // Property 4: Villa (New)
    {
        ID:       "4",
        Title:    "Chaithanya Oak Ville",
        Category: "Villa",
        Price:    "Contact for Price",
        Location: "Whitefield, Bangalore",
        Yield:    "Luxury",
        Image:    "https://images.unsplash.com/photo-1600596542815-6000255ade87?auto=format&fit=crop&w=800&q=80", // Modern Villa Image
    },
}

var Testimonials = []models.Testimonial{
    {ID: "1", Name: "Vivek Chowhan", Role: "Director", Company: "Treebo Hotels", Content: "The team is thorough in Real Estate and they cover the entire spectrum. Extremely helpful and supportive.", Avatar: "https://randomuser.me/api/portraits/men/32.jpg"},
    {ID: "2", Name: "Arvind S.", Role: "CEO", Company: "Mahindra Happinest", Content: "Went out of their way to deeply understand our requirements and come up with a tailor-made solution.", Avatar: "https://randomuser.me/api/portraits/men/45.jpg"},
    {ID: "3", Name: "Shyamli Mishra", Role: "Manager", Company: "Freight Tiger", Content: "I cannot thank you enough for helping us find our office space in Bangalore. Almost as good as being tailor made.", Avatar: "https://randomuser.me/api/portraits/women/44.jpg"},
}