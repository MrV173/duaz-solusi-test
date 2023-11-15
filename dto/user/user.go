package userdto

type CreateUserRequest struct {
	Email    string `json:"email" form:"email" validate:"required"`
	Password string `json:"password" form:"password" validate:"required"`
	FullName string `json:"fullname" form:"fullname" validate:"required"`
	Gender   string `json:"gender" form:"gender" validate:"required"`
	Phone    string `json:"phone" form:"phone" validate:"required"`
}

type UpdateUserRequest struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
	FullName string `json:"fullname" form:"fullname"`
	Gender   string `json:"gender" form:"gender"`
	Phone    string `json:"phone" form:"phone"`
}