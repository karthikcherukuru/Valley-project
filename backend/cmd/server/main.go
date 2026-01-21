package main

import (
    "os"  // <--- This was missing!
    "valley/internal/routes"
)

func main() {
    // 1. Get the PORT from Render
    port := os.Getenv("PORT")
    if port == "" {
        port = "5000" // Default for local testing
    }

    // 2. Setup Routes
    r := routes.SetupRouter()

    // 3. Start the server
    r.Run(":" + port)
}