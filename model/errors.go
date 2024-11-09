package model

const (
	InvalidRequest          = "Invalid request data"
	InternalServerError     = "Internal server error"
	UserNotFound            = "User not found"
	Unauthorized            = "Unauthorized access"
	TokenMissing            = "Token is missing"
	TokenInvalid            = "Invalid or expired token"
)

func ErrorResponse(errMsg string) map[string]string {
	return map[string]string{"error": errMsg}
}
