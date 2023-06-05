package model

import "gorm.io/gorm"

type Requirement struct {
	Degree       string `gorm:"type:varchar(20)" json:"degree"`
	Age          int    `gorm:"type:int" json:"age"` // TODO: the "age" field should be a range
	WorkingYears int    `gorm:"type:int" json:"working_years"`
	Major        string `gorm:"type:varchar(20)" json:"major"`
}

type Salary struct {
	MinSalary int `gorm:"type:int" json:"min_salary"`
	MaxSalary int `gorm:"type:int" json:"max_salary"`
}

type Position struct {
	gorm.Model
	Filename     string `gorm:"type:varchar(64)" json:"file_path"`
	PositionName string `gorm:"type:varchar(20)" json:"position_name"`
	Company      string `gorm:"type:varchar(20)" json:"company"`
	Department   string `gorm:"type:varchar(20)" json:"department"`
	Number       int    `gorm:"type:int" json:"number"`
	Desc         string `gorm:"type:varchar(512)" json:"desc"`
	Salary       `gorm:"embedded"`
	Requirement  `gorm:"embedded"`
}

// TODO: Maybe many bugs, further consideration needed

func GetAllPositions(limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Limit(limit).Offset(offset).Find(&positions)
	return positions, result.Error
}

// NOTE: This function is not used
func GetPositionById(id uint, limit int, offset int) (Position, error) {
	var position Position
	result := Db.Model(&Position{}).Limit(limit).Offset(offset).First(&position, id)
	return position, result.Error
}

func GetPositionsByName(positionName string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("position_name LIKE ?", positionName).
		Limit(limit).Offset(offset).Find(&positions)
	return positions, result.Error
}

func GetPositionsByCompany(company string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("company LIKE ?", company).
		Limit(limit).Offset(offset).Find(&positions)
	return positions, result.Error
}

func GetPositionsByDepartment(department string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("department LIKE ?", department).
		Limit(limit).Offset(offset).Find(&positions)
	return positions, result.Error
}

func GetPositionsByMajor(major string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("major LIKE ?", major).
		Limit(limit).Offset(offset).Find(&positions)
	return positions, result.Error
}

func GetPositionsByDegree(degree string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("degree LIKE ?", degree).
		Limit(limit).Offset(offset).Find(&positions)
	return positions, result.Error
}

func GetPositionsGreaterThanWorkingYears(workingYears int, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("working_years >= ?", workingYears).
		Limit(limit).Offset(offset).Find(&positions)
	return positions, result.Error
}

func CreatePosition(position *Position) error {
	result := Db.Model(&Position{}).Create(position)
	return result.Error
}

func DeletePositionByFilename(filename string) error {
	result := Db.Model(&Position{}).Delete(&Position{Filename: filename})
	return result.Error
}
