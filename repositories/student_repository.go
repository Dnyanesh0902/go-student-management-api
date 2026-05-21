package repositories

import (
	"student-management-api/models"
	"student-management-api/database"
)

func CreateStudent(student *models.Student) error {
	return  database.DB.Create(student).Error
}

func GetStudents() ([]models.Student, error){
	var students []models.Student
	err := database.DB.Find(&students).Error
	return students, err
}

func GetStudentByID(id string) (models.Student, error){
	var student models.Student
	err := database.DB.First(&student, id).Error
	return  student,err
}

func UpdateStudent(student *models.Student) error{
	return  database.DB.Save(student).Error
}

func DeleteStudent(student *models.Student) error {
	return  database.DB.Delete(student).Error
}