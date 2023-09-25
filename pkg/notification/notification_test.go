package notification

import "testing"

func TestIsEmailValid_WithValidEmails(t *testing.T) {
	validEmails := []string{
		"john.doe@example.com",
		"jane_doe1234@example.co.uk",
		"info@sub.domain.com",
	}

	for _, email := range validEmails {
		isValid := isValidEmail(email)
		if !isValid {
			t.Errorf("Expected %s to be a valid email, but it was not", email)
		}
	}
}

func TestIsEmailValid_WithInvalidEmails(t *testing.T) {
	invalidEmails := []string{
		"notanemail",
		"john.doe",
		"example.com",
		"john.doe@.com",
		"john.doe@example",
	}

	for _, email := range invalidEmails {
		isValid := isValidEmail(email)
		if isValid {
			t.Errorf("Expected %s to be an invalid email, but it was not", email)
		}
	}
}

func TestIsPhoneNumberValid_ValidNumbers(t *testing.T) {
	validPhoneNumbers := []string{
		"+359888987654",
		"+359899555123",
	}

	for _, phoneNumber := range validPhoneNumbers {
		isValid := isValidBulgarianPhoneNumber(phoneNumber)
		if !isValid {
			t.Errorf("Expected %s to be a valid Bulgarian phone number, but it was not", phoneNumber)
		}
	}
}

func TestIsPhoneNumberValid_InvalidNumbers(t *testing.T) {
	invalidPhoneNumbers := []string{
		"+35982345",
		"+3598234567890",
		"+359600123456",
		"1234567890",
		"notaphonenumber",
	}

	for _, phoneNumber := range invalidPhoneNumbers {
		isValid := isValidBulgarianPhoneNumber(phoneNumber)
		if isValid {
			t.Errorf("Expected %s to be an invalid Bulgarian phone number, but it was not", phoneNumber)
		}
	}
}

func TestIsValidSlackUsername_ValidUsernames(t *testing.T) {
	validUsernames := []string{
		"john_doe",
		"jane.doe",
		"alice123",
		"user.name123",
	}

	for _, username := range validUsernames {
		isValid := isValidSlackUsername(username)
		if !isValid {
			t.Errorf("Expected %s to be a valid Slack username, but it was not", username)
		}
	}
}

func TestIsValidSlackUsername_InvalidUsernames(t *testing.T) {
	invalidUsernames := []string{
		"user name",
		"user@name",
		"too_long_username12345678901234567890",
	}

	for _, username := range invalidUsernames {
		isValid := isValidSlackUsername(username)
		if isValid {
			t.Errorf("Expected %s to be an invalid Slack username, but it was not", username)
		}
	}
}

func TestIsMessageNotEmpty_ValidMessages(t *testing.T) {
	validMessages := []string{
		"Hello, World!",
		"   This is a message with leading and trailing spaces   ",
		"Non-empty message",
		"1",
	}

	for _, message := range validMessages {
		isValid := isValidMessage(message)
		if !isValid {
			t.Errorf("Expected message '%s' to be not empty, but it was considered empty", message)
		}
	}
}

func TestIsMessageNotEmpty_EmptyMessages(t *testing.T) {
	emptyMessages := []string{
		"",
		"  ",
		"\t",
		"\n",
	}

	for _, message := range emptyMessages {
		isValid := isValidMessage(message)
		if isValid {
			t.Errorf("Expected message '%s' to be empty, but it was considered not empty", message)
		}
	}
}
