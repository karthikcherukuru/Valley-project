package handlers

import (
	"fmt"
	"net/http"
	"net/smtp"

	"github.com/gin-gonic/gin"
)

// --- CONFIGURATION (FILL THESE IN LOCALLY) ---
const (
	SenderEmail    = "cherukuru.karthi@gmail.com" // <--- Update this locally
	SenderPassword = "xokh drpr qstv owdy"        // <--- Update this locally
	SMTPHost       = "smtp.gmail.com"
	SMTPPort       = "587"
)

// Define the data we expect from the frontend
type ContactRequest struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required"`
	Phone    string `json:"phone"` // <--- Correct
	Interest string `json:"interest"`
	Message  string `json:"message" binding:"required"`
}

func SubmitContact(c *gin.Context) {
	var req ContactRequest

	// 1. Bind the JSON data
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	// 2. Format the Email Body
	subject := "Subject: New Inquiry from Valley Website\n"
	mime := "MIME-version: 1.0;\nContent-Type: text/html; charset=\"UTF-8\";\n\n"

	// FIX IS HERE: Added the Phone line to the HTML string below
	body := fmt.Sprintf(`
		<html>
		<body>
			<h2>New Contact Inquiry</h2>
			<p><strong>Name:</strong> %s</p>
			<p><strong>Email:</strong> %s</p>
			<p><strong>Phone:</strong> %s</p>
			<p><strong>Interest:</strong> %s</p>
			<p><strong>Message:</strong><br/>%s</p>
		</body>
		</html>
	`, req.Name, req.Email, req.Phone, req.Interest, req.Message)

	msg := []byte(subject + mime + body)

	// 3. Send the Email using Gmail SMTP
	auth := smtp.PlainAuth("", SenderEmail, SenderPassword, SMTPHost)
	addr := SMTPHost + ":" + SMTPPort

	// We send the email TO ourselves (SenderEmail)
	err := smtp.SendMail(addr, auth, SenderEmail, []string{SenderEmail}, msg)

	if err != nil {
		fmt.Println("Email Error:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to send email"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Email sent successfully!"})
}
