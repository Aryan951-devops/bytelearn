// Package user handles user profile data and adjustments.
package user

// ChangePasswordRequest holds data to change a password.
type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

// UpdateAccountRequest holds data to update account settings.
type UpdateAccountRequest struct {
	PhoneNo string `form:"phone_no"`
	City    string `form:"city"`
	State   string `form:"state"`
	Pincode string `form:"pincode"`
}
