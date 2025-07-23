package service

import (
	"backend/model"
	"backend/repository"
)

type StudentService struct {
	StudentRepo *repository.StudentRepository
}

func (s *StudentService) CreateStudent(student *model.Student) error {
	return s.StudentRepo.CreateStudent(student)
}

func (s *StudentService) GetStudents() ([]model.Student, error) {
	return s.StudentRepo.GetStudents()
}

func (s *StudentService) GetStudentByID(id uint) (*model.Student, error) {
	return s.StudentRepo.GetStudentByID(id)
}

func (s *StudentService) UpdateStudent(student *model.Student) error {
	return s.StudentRepo.UpdateStudent(student)
}

func (s *StudentService) GetStudentByStudentID(studentID string) (*model.Student, error) {
	return s.StudentRepo.GetStudentByStudentID(studentID)
}

func (s *StudentService) DeleteStudentByStudentID(studentID string) error {
	return s.StudentRepo.DeleteStudentByStudentID(studentID)
}

func (s *StudentService) DeleteStudent(id uint) error {
	return s.StudentRepo.DeleteStudent(id)
}
