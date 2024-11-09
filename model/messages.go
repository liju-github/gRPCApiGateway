package model

const (
	RegistrationSuccessful   = "Registration successful"
	LoginSuccessful          = "Login successful"
	EmailVerificationSuccess = "Email verified successfully"
	ProfileRetrieved         = "Profile retrieved successfully"
)

func SuccessResponse(msg string) map[string]string {
	return map[string]string{"message": msg}
}
