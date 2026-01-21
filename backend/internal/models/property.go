package models

type Property struct {
    ID       string `json:"id"`
    Title    string `json:"title"`
    Category string `json:"category"`
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