package main

import (
	"valley/internal/routes"
)
func main() {
    // 1. Get the PORT from Render (they use "PORT", not "port")
    port := os.Getenv("PORT")
    if port == "" {
        port = "5000" // Default for local testing
    }

    // 2. Setup Routes
    r := routes.SetupRouter()

    // 3. Start the server
    // IMPORTANT: We use the 'port' variable here
    r.Run(":" + port)
}