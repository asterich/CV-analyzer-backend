package cv

import (
	"time"

	"gorm.io/gorm"
)

// TODO: Database structures need more consideration.
//       Now just don't use database, use an array instead.

type ContactInfo struct {
	Tel   string `gorm:"type:varchar(20)" json:"tel"`
	Email string `gorm:"type:varchar(50)" json:"email"`
}

type Duration struct {
	Begin time.Time `gorm:"type:time" json:"start_time"`
	End   time.Time `gorm:"type:time" json:"end_time"`
}

type Education struct {
	Duration
	School  string `gorm:"type:varchar(50)" json:"school"`
	Degree  string `gorm:"type:varchar(20)" json:"degree"`
	Major   string `gorm:"type:varchar(20)" json:"major"`
	Lessons string `gorm:"type:varchar(256)" json:"lessons"`
}

type Experience struct {
	Company_or_organization string `gorm:"type:varchar(50)" json:"company_or_organization"`
	Position                string `gorm:"type:varchar(20)" json:"position"`
}

type WorkExperience struct {
	Duration
	Experience
}

type ProjectExperience struct {
	Name string `gorm:"type:varchar(50)" json:"project_name"`
	Desc string `gorm:"type:varchar(512)" json:"project_description"`
}

type Award struct {
	Name  string `gorm:"type:varchar(50)" json:"award_name"`
	Level string `gorm:"type:varchar(20)" json:"level"`
}

type Skill struct {
	Name string `gorm:"type:varchar(32)" json:"skill_name"`
}

type CV struct {
	gorm.Model
	FileUrl               string `gorm:"type:varchar(64)" json:"file_path"`
	Name                  string `gorm:"type:varchar(16)" json:"name"`
	Age                   int    `gorm:"type:int" json:"age"`
	ContactInfo           ContactInfo
	Educations            []Education
	WorkExperiences       []WorkExperience
	SchoolExperiences     []Experience
	InternshipExperiences []Experience
	ProjectExperiences    []ProjectExperience
	Awards                []Award
	Skills                []Skill
	SelfDesc              string `gorm:"type:varchar(1024)" json:"self_description"`
}
