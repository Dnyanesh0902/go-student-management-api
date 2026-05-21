package dto

type CreateStudentRequest struct {
	Name string `json:"name" binding:"required"`
	Email string `json:"email" binding:"required,email"`
	Age int `json:"age" binding:"required,gte=1"`
	Course string `json:"course" binding:"required"`
	City string `json:"city" binding:"required"`
}