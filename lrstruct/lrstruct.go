// Package lrstruct contains response structs used by SDK tests and examples.
package lrstruct

// AuthSecurityQuestion is the response shape returned by authentication
// security-question lookup APIs.
type AuthSecurityQuestion []SecurityQuestion

type SecurityQuestion struct {
	QuestionID string `json:"QuestionId"`
}
