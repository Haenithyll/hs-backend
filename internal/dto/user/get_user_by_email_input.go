package dto

type GetUserByEmailInput struct {
	Email string `form:"email" binding:"required,email"`
}
