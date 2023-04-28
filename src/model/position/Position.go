package position

import "gorm.io/gorm"

type Requirement struct {
	Degree       string `gorm:"type:varchar(20)" json:"degree"`
	Age          int    `gorm:"type:int" json:"age"`
	WorkingYears int    `gorm:"type:int" json:"working_years"`
	Major        string `gorm:"type:varchar(20)" json:"major"`
}

type Position struct {
	gorm.Model
	FileUrl      string `gorm:"type:varchar(64)" json:"file_path"`
	PositionName string `gorm:"type:varchar(20)" json:"position_name"`
	Department   string `gorm:"type:varchar(20)" json:"department"`
	Number       int    `gorm:"type:int" json:"number"`
	Desc         string `gorm:"type:varchar(512)" json:"desc"`
	Requirement
}
