package controller

import (
	"backend/model"
	"backend/service"
	"github.com/gin-gonic/gin"
	"net/http"
	"regexp"
	"strconv"
	"time"
)

type StudentController struct {
	StudentService *service.StudentService
}

// 统一错误处理函数
func handleError(ctx *gin.Context, statusCode int, errMsg string) {
	ctx.JSON(statusCode, ErrorResponse{Error: errMsg})
	ctx.JSON(statusCode, gin.H{"error": errMsg})
}

// 解析 ID 函数
func parseID(ctx *gin.Context) (uint, error) {
	idStr := ctx.Param("id")
	id, err := strconv.ParseUint(idStr, 10, 64)
	return uint(id), err
}

// CreateStudent 创建学生信息
// @Summary 创建学生信息
// @Description 创建新的大学学生信息
// @Tags 学生管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param student body model.Student true "学生信息"
// @Success 201 {object} SuccessResponse "创建成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /students [post]
func (c *StudentController) CreateStudent(ctx *gin.Context) {
	var student model.Student
	if err := ctx.ShouldBindJSON(&student); err != nil {
		handleError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	// 校验学号是否为 11 位数字，且入学年份合理
	matched, err := regexp.MatchString(`^(20[0-9]{2})(\d{4})(\d{3})$`, student.StudentID)
	if err != nil || !matched {
		handleError(ctx, http.StatusBadRequest, "学号必须由 4 位入学年份、4 位专业序号和 3 位学生编号组成的 11 位数字")
		return
	}

	// 检查入学年份是否合理
	year, _ := strconv.Atoi(student.StudentID[:4])
	currentYear := time.Now().Year()
	if year > currentYear || year < currentYear-10 {
		handleError(ctx, http.StatusBadRequest, "入学年份不合理")
		return
	}

	if err := c.StudentService.CreateStudent(&student); err != nil {
		handleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusCreated, SuccessResponse{Message: "Student created successfully"})
}

// UpdateStudent 更新学生信息
// @Summary 更新学生信息
// @Description 根据学生 ID 更新学生的信息
// @Tags 学生管理
// @Accept json
// @Produce json
// @Security BearerAuth
// @Param id path uint true "学生 ID"
// @Param student body model.Student true "更新后的学生信息"
// @Success 200 {object} SuccessResponse "更新成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 404 {object} ErrorResponse "学生信息不存在"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /students/{studentID} [put]
func (c *StudentController) UpdateStudent(ctx *gin.Context) {
	// 获取学号
	studentID := ctx.Param("studentID")

	var updateData model.Student
	if err := ctx.ShouldBindJSON(&updateData); err != nil {
		handleError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	student, err := c.StudentService.GetStudentByStudentID(studentID)
	if err != nil {
		handleError(ctx, http.StatusNotFound, "Student not found")
		return
	}

	student.Name = updateData.Name
	student.Gender = updateData.Gender
	student.Age = updateData.Age
	student.Class = updateData.Class
	student.Major = updateData.Major
	student.Grade = updateData.Grade

	// 简单校验性别，仅允许男、女
	if student.Gender != "男" && student.Gender != "女" {
		handleError(ctx, http.StatusBadRequest, "性别只能为男或女")
		return
	}

	if err := c.StudentService.UpdateStudent(student); err != nil {
		handleError(ctx, http.StatusInternalServerError, "Failed to update student")
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Message: "Student updated successfully"})
}

// GetStudents 获取学生信息列表
// @Summary 获取学生信息列表
// @Description 获取所有学生的信息列表
// @Tags 学生管理
// @Produce json
// @Security BearerAuth
// @Success 200 {array} model.Student "获取成功"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /students [get]
func (c *StudentController) GetStudents(ctx *gin.Context) {
	students, err := c.StudentService.GetStudents()
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, err.Error())
		return
	}
	ctx.JSON(http.StatusOK, students)
}

// GetStudent 获取单个学生信息
// @Summary 获取单个学生信息
// @Description 根据学生学号获取单个学生的信息
// @Tags 学生管理
// @Produce json
// @Security BearerAuth
// @Param studentID path string true "学生学号"
// @Success 200 {object} model.Student "获取成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 404 {object} ErrorResponse "学生信息不存在"
// @Router /students/{studentID} [get]
func (c *StudentController) GetStudent(ctx *gin.Context) {
	studentID := ctx.Param("studentID")
	// 假设学号必须是 11 位数字
	if len(studentID) != 11 {
		handleError(ctx, http.StatusBadRequest, "学号必须为 11 位数字")
		return
	}
	// 使用正则表达式验证是否全为数字
	match, err := regexp.MatchString("^[0-9]+$", studentID)
	if err != nil {
		handleError(ctx, http.StatusInternalServerError, "学号格式验证出错")
		return
	}
	if !match {
		handleError(ctx, http.StatusBadRequest, "学号必须全为数字")
		return
	}

	student, err := c.StudentService.GetStudentByStudentID(studentID)
	if err != nil {
		handleError(ctx, http.StatusNotFound, "未找到该学生信息")
		return
	}

	ctx.JSON(http.StatusOK, student)
}

// DeleteStudent 删除学生信息
// @Summary 删除学生信息
// @Description 根据学生 ID 删除学生的信息
// @Tags 学生管理
// @Produce json
// @Security BearerAuth
// @Param id path uint true "学生 ID"
// @Success 200 {object} SuccessResponse "删除成功"
// @Failure 400 {object} ErrorResponse "请求参数错误"
// @Failure 500 {object} ErrorResponse "服务器内部错误"
// @Router /students/{studentID} [delete]
func (c *StudentController) DeleteStudent(ctx *gin.Context) {
	// 获取学号
	studentID := ctx.Param("studentID")

	if err := c.StudentService.DeleteStudentByStudentID(studentID); err != nil {
		handleError(ctx, http.StatusInternalServerError, "Failed to delete student")
		return
	}

	ctx.JSON(http.StatusOK, SuccessResponse{Message: "Student deleted successfully"})
}
