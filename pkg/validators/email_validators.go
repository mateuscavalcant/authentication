package validators

import (
	CON "authentication/pkg/database"

	"github.com/badoux/checkmail"
)


// ValidateFormatEmail checks the format of the provided email address.
func ValidateFormatEmail(email string) error {
	// Use the checkmail library to validate the email format.
	err := checkmail.ValidateFormat(email)
	// Return any errors encountered during validation.
	return err
}

// ExistEmail checks if the provided email already exists in the database.
func ExistEmail(email string) (bool, error) {
	// Get the database connection.
	db := CON.DB()

	var emailCount int

	// Query the database to count the number of occurrences of the provided email.
	err := db.QueryRow("SELECT COUNT(id) AS emailCount FROM user1 WHERE email=?", email).Scan(&emailCount)
	if err != nil {
		// Return false and any errors encountered during the query.
		return false, err
	}

	// Return true if the email count is greater than 0 (email exists), otherwise return false.
	return emailCount > 0, nil
}