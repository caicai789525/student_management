package repository

import (
	"backend/model"
	"gorm.io/gorm"
)

type StudentRepository struct {
	DB *gorm.DB
}

func (r *StudentRepository) CreateStudent(student *model.Student) error {
	return r.DB.Create(student).Error
}

func (r *StudentRepository) GetStudents() ([]model.Student, error) {
	var students []model.Student
	err := r.DB.Find(&students).Error
	return students, err
}

func (r *StudentRepository) GetStudentByID(id uint) (*model.Student, error) {
	var student model.Student
	err := r.DB.First(&student, id).Error
	return &student, err
}

func (r *StudentRepository) UpdateStudent(student *model.Student) error {
	return r.DB.Save(student).Error
}

func (r *StudentRepository) GetStudentByStudentID(studentID string) (*model.Student, error) {
	var student model.Student
	err := r.DB.Where("student_id = ?", studentID).First(&student).Error
	return &student, err
}

func (r *StudentRepository) DeleteStudentByStudentID(studentID string) error {
	return r.DB.Where("student_id = ?", studentID).Delete(&model.Student{}).Error
}

func (r *StudentRepository) DeleteStudent(id uint) error {
	return r.DB.Delete(&model.Student{}, id).Error
}
