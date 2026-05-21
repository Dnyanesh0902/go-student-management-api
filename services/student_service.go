package services

import (
	"student-management-api/models"
	"student-management-api/repositories"
)


func CreateStudent(student *models.Student) error {
	return  repositories.CreateStudent(student)
}

func GetStudents()([]models.Student, error){
	return  repositories.GetStudents()
}

func GetStudentByID(id string)(models.Student, error){
	return repositories.GetStudentByID(id)
}

func UpdateStudent(student *models.Student) error {
	return  repositories.UpdateStudent(student)
}
func DeleteStudent(student *models.Student) error{
	return  repositories.DeleteStudent(student)
}