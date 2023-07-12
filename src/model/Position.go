package model

import "log"

type IntDuration struct {
	Begin int `gorm:"type:int" json:"begin"`
	End   int `gorm:"type:int" json:"end"`
}

type Requirement struct {
	Degree       string      `gorm:"type:varchar(20)" json:"degree"`
	Age          IntDuration `gorm:"embedded;embeddedPrefix:age_" json:"age"`
	WorkingYears IntDuration `gorm:"embedded;embeddedPrefix:working_years_" json:"working_years"`
	Major        string      `gorm:"type:varchar(20)" json:"major"`
	Others       string      `gorm:"type:varchar(512)" json:"others"`
}

// type Salary struct {
// 	MinSalary int `gorm:"type:int" json:"min_salary"`
// 	MaxSalary int `gorm:"type:int" json:"max_salary"`
// }

type Position struct {
	ID int `gorm:"type:int;primaryKey" json:"id"`
	// Filename     string `gorm:"type:varchar(64)" json:"file_path"`
	PositionName string `gorm:"type:varchar(20)" json:"position_name"`
	// Company      string `gorm:"type:varchar(20)" json:"company"`
	// Department string `gorm:"type:varchar(20)" json:"department"`
	// Number       int    `gorm:"type:int" json:"number"`
	Desc string `gorm:"type:varchar(512)" json:"desc"`
	// Salary       `gorm:"embedded"`
	Requirement `gorm:"embedded"`
}

// TODO: Maybe many bugs, further consideration needed

func GetAllPositions(limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Limit(limit).Offset((offset - 1) * limit).Find(&positions)
	log.Println("positions:", positions)
	log.Println("error:", result.Error)
	return positions, result.Error
}

// NOTE: This function is not used
func GetPositionById(id uint, limit int, offset int) (Position, error) {
	var position Position
	result := Db.Model(&Position{}).Limit(limit).Offset((offset-1)*limit).First(&position, id)
	return position, result.Error
}

func GetPositionsByName(positionName string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("position_name LIKE ?", positionName).
		Limit(limit).Offset((offset - 1) * limit).Find(&positions)
	return positions, result.Error
}

// func GetPositionsByCompany(company string, limit int, offset int) ([]Position, error) {
// 	var positions []Position
// 	result := Db.Model(&Position{}).Where("company LIKE ?", company).
// 		Limit(limit).Offset((offset - 1) * limit).Find(&positions)
// 	return positions, result.Error
// }

// func GetPositionsByDepartment(department string, limit int, offset int) ([]Position, error) {
// 	var positions []Position
// 	result := Db.Model(&Position{}).Where("department LIKE ?", department).
// 		Limit(limit).Offset((offset - 1) * limit).Find(&positions)
// 	return positions, result.Error
// }

func GetPositionsByMajor(major string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("major LIKE ?", major).
		Limit(limit).Offset((offset - 1) * limit).Find(&positions)
	return positions, result.Error
}

func GetPositionsByDegree(degree string, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("degree LIKE ?", degree).
		Limit(limit).Offset((offset - 1) * limit).Find(&positions)
	return positions, result.Error
}

func GetPositionsInRangeOfWorkingYears(workingYears Duration, limit int, offset int) ([]Position, error) {
	var positions []Position
	result := Db.Model(&Position{}).Where("working_years between ? and ?", workingYears.Begin, workingYears.End).
		Limit(limit).Offset((offset - 1) * limit).Find(&positions)
	return positions, result.Error
}

func CreatePosition(position *Position) error {
	result := Db.Model(&Position{}).Create(position)
	return result.Error
}

func DeletePositionByID(id int) error {
	result := Db.Model(&Position{}).Delete(&Position{ID: id})
	return result.Error
}
