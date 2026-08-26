package techpalace
import (
    "fmt"
    "strings"
    )
// WelcomeMessage returns a welcome message for the customer.
func WelcomeMessage(customer string) string {
    customer = strings.ToUpper(customer)
	return fmt.Sprintf("Welcome to the Tech Palace, %s",customer)
}

// AddBorder adds a border to a welcome message.
func AddBorder(welcomeMsg string, numStarsPerLine int) string {
	border := strings.Repeat("*",numStarsPerLine)
    return fmt.Sprintf("%s\n%s\n%s",border,welcomeMsg,border)
}

// CleanupMessage cleans up an old marketing message.
func CleanupMessage(oldMsg string) string {
	result := strings.ReplaceAll(oldMsg,"*","")
    trim := strings.TrimSpace(result) 
    // res := strings.ReplaceAll(result,"\n","")
    return trim
}
