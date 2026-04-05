package validation

import (
	"fmt"
	"regexp"

	"al.essio.dev/pkg/shellescape"
)

// Simple validation for instance names
var validNamePattern = regexp.MustCompile(`^[a-zA-Z0-9_-]+$`)

type ValidationError error

// EscapeForShell returns the string properly quoted for safe use in shell commands.
// Uses shellescape.Quote to handle all special characters, metacharacters, and control characters.
func EscapeForShell(value string) string {
	return shellescape.Quote(value)
}

func ValidateInstanceName(name string) (string, error) {
	// Validate instance name
	if name == "" {
		return "", ValidationError(fmt.Errorf("name cannot be empty"))
	}
	if !validNamePattern.MatchString(name) {
		return "", ValidationError(fmt.Errorf("name contains invalid characters (only alphanumeric, hyphens, underscores allowed)"))
	}
	if len(name) > 50 {
		return "", ValidationError(fmt.Errorf("name too long (max 50 characters)"))
	}
	return name, nil
}
