package controllers

import (
	"net/http"
	"strconv"

	"student-management-api/database"
	"student-management-api/models"
	"github.com/gin-gonic/gin"
)

func CreateStudent(c *gin.Context) {
	var student models.Student

	if err := c.ShouldBindJSON(&student); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error" : err.Error()})
		return
	}

	database.DB.Create(&student)
	c.JSON(http.StatusCreated, gin.H{
		"Message":"Student Created Successfully",
		"data": student,
	})
}

func GetStudent(c *gin.Context) {
	var student []models.Student

	database.DB.Find(&student)
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func GetStudentById(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"message": "Student Not Found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"data": student,
	})
}

func UpdateStudent(c *gin.Context) {
	id := c.Param("id")
	var student models.Student

	if err := database.DB.First(&student, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"Message":"Student Not Found"})
		return
	}

	var input models.Student

	if err := c.ShouldBindJSON(&input); err != nil{
		c.JSON(http.StatusBadRequest, gin.H{"error":err.Error()})
		return
	}

	student.Name = input.Name
	student.Course = input.Course
	student.Email = input.Email
	student.Age = input.Age
	student.Course = input.Course
	student.City = input.City

	database.DB.Save(&student)

	c.JSON(http.StatusOK, gin.H{
		"message":"Student Updated Successfully",
		"data":student,
	})
}

func DeleteStudent(c *gin.Context) {
	id := c.Param("id")
	
	studentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Invalid Student Id",
		})
		return
	}

	var student models.Student

	if err := database.DB.First(&student, studentID).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":"Student Not Found",
		})
		return
	}
	database.DB.Delete(&student)

	c.JSON(http.StatusOK, gin.H{
		"message": "student deleted successfully",
	})
}