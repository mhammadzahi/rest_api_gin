package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	//"rest_api_gin/internal/database/db"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "muhammad"
	password = "k9999"
	dbname   = "testdb"
)

var dbConn *pgx.Conn
var store *db.Store

func main() {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	conn, err := pgx.Connect(context.Background(), connStr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close(context.Background())

	err = conn.Ping(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Successfully connected to the database")

	// Initialize the database store
	store = db.NewStore(conn)

	// Set up the Gin router
	r := gin.Default()

	// Routes
	r.POST("/api/users", createUser)
	r.POST("/api/users/generateotp", generateOTP)
	r.POST("/api/users/verifyotp", verifyOTP)

	// Run the server
	r.Run(":8080")
}

func createUser(c *gin.Context) {
	var user db.User

	// Bind JSON request body to User struct
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Validate input data (you may add more validation as needed)

	// Check if the phone number already exists
	existingUser, err := store.GetUserByPhoneNumber(context.Background(), user.PhoneNumber)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Phone number already exists"})
		return
	}

	// Generate OTP and set expiration time
	otp := generateRandomOTP()
	expirationTime := time.Now().Add(1 * time.Minute)

	// Create the user in the database
	newUser, err := store.CreateUser(context.Background(), user.Name, user.PhoneNumber, otp, expirationTime)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
		return
	}

	c.JSON(http.StatusCreated, newUser)
}

func generateOTP(c *gin.Context) {
	// Implement the logic to generate OTP for a user
}

func verifyOTP(c *gin.Context) {
	// Implement the logic to verify OTP for a user
}

func generateRandomOTP() string {
	// Implement logic to generate a random 4-digit OTP
	return "1234" // Placeholder, replace with actual logic
}
