package user

type ChangePasswordRequest struct {
	NewPassword string `json:"new_password" binding:"required,min=6"`
}

type UpdateAccountRequest struct {
	PhoneNo string `form:"phone_no"`
	City    string `form:"city"`
	State   string `form:"state"`
	Pincode string `form:"pincode"`
}
