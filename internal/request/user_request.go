package request

type GetUserByEmailRequest struct {
	Email string `uri:"email" json:"-" binding:"required"`
}
