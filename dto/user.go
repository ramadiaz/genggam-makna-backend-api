package dto

type User struct {
	ID        string `json:"id"`
	Email     string `json:"email" form:"email" binding:"required"`
	Password  string `json:"password,omitempty" form:"password"`
	FirstName string `json:"first_name" form:"first_name" binding:"required"`
	LastName  string `json:"last_name" form:"last_name" binding:"required"`
	GoogleUID string `json:"google_uid,omitempty" form:"google_uid"`
}
