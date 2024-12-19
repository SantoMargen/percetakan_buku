package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"
)

func generateRandomString(length int) string {
	// Generate a random string of letters (A-Z, a-z)
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	var sb strings.Builder
	for i := 0; i < length; i++ {
		sb.WriteByte(charset[rand.Intn(len(charset))])
	}
	return sb.String()
}

func main() {
	// Open the file where we will save the SQL statements
	file, err := os.Create("dummy_users.sql")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Start the INSERT statements
	file.WriteString("-- Dummy users data\n\n")

	for i := 1; i <= 100000; i++ {
		fullName := fmt.Sprintf("USER %d", i)
		email := fmt.Sprintf("user%d@mail.com", i)
		password := "$2a$10$K0U2/v0ofZce0GUfPaVyB.TfKembIJOj83SCdhAxM00ZlRaaVtjBy" // Example hashed password
		phoneNumber := fmt.Sprintf("123-456-%03d", rand.Intn(999))                 // Random phone number
		dateOfBirth := fmt.Sprintf("1990-01-%02d", rand.Intn(28)+1)                // Random date of birth in January
		profilePicture := "NULL"                                                   // No profile picture
		gender := "M"                                                              // You can alternate between "M" and "F" if needed
		address := fmt.Sprintf("Address %d, Some Street, Some City", i)
		city := "SomeCity"
		country := "SomeCountry"
		role := "USER" // You can randomize roles if needed
		createdBy := 1
		updatedBy := 1
		createdAt := time.Now().Format("2006-01-02 15:04:05")
		updatedAt := createdAt

		// Construct the SQL INSERT statement
		sqlStatement := fmt.Sprintf(
			"INSERT INTO \"public\".\"users\" (\"id\", \"full_name\", \"email\", \"password\", \"phone_number\", \"date_of_birth\", \"profile_picture\", \"gender\", \"address\", \"city\", \"country\", \"role\", \"created_by\", \"updated_by\", \"created_at\", \"updated_at\") VALUES (%d, '%s', '%s', '%s', '%s', '%s', %s, '%s', '%s', '%s', '%s', '%s', '%d', '%d', '%s', '%s');\n",
			i, fullName, email, password, phoneNumber, dateOfBirth, profilePicture, gender, address, city, country, role, createdBy, updatedBy, createdAt, updatedAt)

		// Write the SQL statement to the file
		_, err := file.WriteString(sqlStatement)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}

	fmt.Println("SQL file generated: dummy_users.sql")
}
