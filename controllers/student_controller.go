package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"student-management-api/dto"
	"student-management-api/models"
	"student-management-api/services"
	"student-management-api/utils"
)

func CreateStudent(c *gin.Context) {
	var input dto.CreateStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.GetValidationErrors(err)})
		return
	}
	student := models.Student{
		Name:   input.Name,
		Email:  input.Email,
		Age:    input.Age,
		Course: input.Course,
		City:   input.City,
	}
	if err := services.CreateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to create student"})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"Message": "Student Created Successfully",
		"data":    student,
	})
}

func GetStudents(c *gin.Context) {
	students, err := services.GetStudents()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to fetch students"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": students,
	})
}

func GetStudentByID(c *gin.Context) {
	id := c.Param("id")
	student, err := services.GetStudentByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Student Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	student, err := services.GetStudentByID(id)

	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message": "Student Not Found"})
		return
	}

	var input dto.CreateStudentRequest

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": utils.GetValidationErrors(err)})
		return
	}

	student.Name = input.Name
	student.Email = input.Email
	student.Age = input.Age
	student.Course = input.Course
	student.City = input.City

	if err := services.UpdateStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to update student"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "Student Updated Successfully",
		"data":    student,
	})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")

	_, err := strconv.Atoi(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Invalid Student Id",
		})
		return
	}
	student, err := services.GetStudentByID(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "Student Not Found",
		})
		return
	}
	if err := services.DeleteStudent(&student); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to delete Student"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted successfully",
	})
}
