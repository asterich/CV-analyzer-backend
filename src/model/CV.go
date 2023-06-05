package model

import (
	"log"
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
	gorm.Model
	CVId     uint `gorm:"type:int" json:"cv_id"`
	Duration `gorm:"embedded"`
	School   string `gorm:"type:varchar(50)" json:"school"`
	Degree   string `gorm:"type:varchar(20)" json:"degree"`
	Major    string `gorm:"type:varchar(20)" json:"major"`
	Lessons  string `gorm:"type:varchar(256)" json:"lessons"`
}

type Experience struct {
	CompanyOrOrganization string `gorm:"type:varchar(50)" json:"company_or_organization"`
	Position              string `gorm:"type:varchar(20)" json:"position"`
}

type WorkExperience struct {
	gorm.Model
	CVId       uint `gorm:"type:int" json:"cv_id"`
	Duration   `gorm:"embedded"`
	Experience `gorm:"embedded"`
}

type SchoolExperience struct {
	gorm.Model
	CVId       uint `gorm:"type:int" json:"cv_id"`
	Experience `gorm:"embedded"`
}

type InternshipExperience struct {
	gorm.Model
	CVId       uint `gorm:"type:int" json:"cv_id"`
	Experience `gorm:"embedded"`
}

type ProjectExperience struct {
	gorm.Model
	CVId uint   `gorm:"type:int" json:"cv_id"`
	Name string `gorm:"type:varchar(50)" json:"project_name"`
	Desc string `gorm:"type:varchar(512)" json:"project_description"`
}

type Award struct {
	gorm.Model
	CVId  uint   `gorm:"type:int" json:"cv_id"`
	Name  string `gorm:"type:varchar(50)" json:"award_name"`
	Level string `gorm:"type:varchar(20)" json:"level"`
}

type Skill struct {
	gorm.Model
	CVId uint   `gorm:"type:int" json:"cv_id"`
	Name string `gorm:"type:varchar(32)" json:"skill_name"`
}

type CV struct {
	gorm.Model
	Filename              string      `gorm:"type:varchar(64)" json:"filename"`
	Name                  string      `gorm:"type:varchar(16)" json:"name"`
	Age                   int         `gorm:"type:int" json:"age"`
	ContactInfo           ContactInfo `gorm:"embedded"`
	Degree                string      `gorm:"type:varchar(20)" json:"degree"` // Degree is extracted from education
	WorkingYears          int         `gorm:"type:int" json:"working_years"`  // WorkingYears is extracted from work experience
	Educations            []Education
	WorkExperiences       []WorkExperience
	SchoolExperiences     []SchoolExperience
	InternshipExperiences []InternshipExperience
	ProjectExperiences    []ProjectExperience
	Awards                []Award
	Skills                []Skill
	SelfDesc              string `gorm:"type:varchar(1024)" json:"self_description"`
}

func constructCVArrayFields(cv *CV) error {
	cv.Educations = []Education{}
	err := Db.Model(&Education{}).Where("cv_id = ?", cv.ID).Find(&cv.Educations).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get educations, err:", err.Error())
		return err
	}

	cv.WorkExperiences = []WorkExperience{}
	err = Db.Model(&WorkExperience{}).Where("cv_id = ?", cv.ID).Find(&cv.WorkExperiences).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get work experiences, err:", err.Error())
		return err
	}

	cv.SchoolExperiences = []SchoolExperience{}
	err = Db.Model(&SchoolExperience{}).Where("cv_id = ?", cv.ID).Find(&cv.SchoolExperiences).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get school experiences, err:", err.Error())
		return err
	}

	cv.InternshipExperiences = []InternshipExperience{}
	err = Db.Model(&InternshipExperience{}).Where("cv_id = ?", cv.ID).Find(&cv.InternshipExperiences).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get internship experiences, err:", err.Error())
		return err
	}

	cv.ProjectExperiences = []ProjectExperience{}
	err = Db.Model(&ProjectExperience{}).Where("cv_id = ?", cv.ID).Find(&cv.ProjectExperiences).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get project experiences, err:", err.Error())
		return err
	}

	cv.Awards = []Award{}
	err = Db.Model(&Award{}).Where("cv_id = ?", cv.ID).Find(&cv.Awards).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get awards, err:", err.Error())
		return err
	}

	cv.Skills = []Skill{}
	err = Db.Model(&Skill{}).Where("cv_id = ?", cv.ID).Find(&cv.Skills).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Println("Failed to get skills, err:", err.Error())
		return err
	}
	return nil
}

func saveCVArrayFields(cv *CV) error {
	for i := range cv.Educations {
		cv.Educations[i].CVId = cv.ID
		err := Db.Save(&cv.Educations[i]).Error
		if err != nil {
			log.Println("Failed to save education, err:", err.Error())
			return err
		}
	}

	for i := range cv.WorkExperiences {
		cv.WorkExperiences[i].CVId = cv.ID
		err := Db.Save(&cv.WorkExperiences[i]).Error
		if err != nil {
			log.Println("Failed to save work experience, err:", err.Error())
			return err
		}
	}

	for i := range cv.SchoolExperiences {
		cv.SchoolExperiences[i].CVId = cv.ID
		err := Db.Save(&cv.SchoolExperiences[i]).Error
		if err != nil {
			log.Println("Failed to save school experience, err:", err.Error())
			return err
		}
	}

	for i := range cv.InternshipExperiences {
		cv.InternshipExperiences[i].CVId = cv.ID
		err := Db.Save(&cv.InternshipExperiences[i]).Error
		if err != nil {
			log.Println("Failed to save internship experience, err:", err.Error())
			return err
		}
	}

	for i := range cv.ProjectExperiences {
		cv.ProjectExperiences[i].CVId = cv.ID
		err := Db.Save(&cv.ProjectExperiences[i]).Error
		if err != nil {
			log.Println("Failed to save project experience, err:", err.Error())
			return err
		}
	}

	for i := range cv.Awards {
		cv.Awards[i].CVId = cv.ID
		err := Db.Save(&cv.Awards[i]).Error
		if err != nil {
			log.Println("Failed to save award, err:", err.Error())
			return err
		}
	}

	for i := range cv.Skills {
		cv.Skills[i].CVId = cv.ID
		err := Db.Save(&cv.Skills[i]).Error
		if err != nil {
			log.Println("Failed to save skill, err:", err.Error())
			return err
		}
	}

	return nil
}

func deleteCVArrayFields(cv *CV) error {
	// TODO: check whether it's right
	err := Db.Where("cv_id = ?", cv.ID).Delete(&Education{}).Error
	if err != nil {
		log.Println("Failed to delete education, err:", err.Error())
		return err
	}

	err = Db.Where("cv_id = ?", cv.ID).Delete(&WorkExperience{}).Error
	if err != nil {
		log.Println("Failed to delete work experience, err:", err.Error())
		return err
	}

	err = Db.Where("cv_id = ?", cv.ID).Delete(&SchoolExperience{}).Error
	if err != nil {
		log.Println("Failed to delete school experience, err:", err.Error())
		return err
	}

	err = Db.Where("cv_id = ?", cv.ID).Delete(&InternshipExperience{}).Error
	if err != nil {
		log.Println("Failed to delete internship experience, err:", err.Error())
		return err
	}

	err = Db.Where("cv_id = ?", cv.ID).Delete(&ProjectExperience{}).Error
	if err != nil {
		log.Println("Failed to delete project experience, err:", err.Error())
		return err
	}

	err = Db.Where("cv_id = ?", cv.ID).Delete(&Award{}).Error
	if err != nil {
		log.Println("Failed to delete award, err:", err.Error())
		return err
	}

	err = Db.Where("cv_id = ?", cv.ID).Delete(&Skill{}).Error
	if err != nil {
		log.Println("Failed to delete skill, err:", err.Error())
		return err
	}

	return nil
}

func GetCVByFilename(path string, limit int, offset int) (CV, error) {
	var cv CV
	err := Db.Model(&CV{}).Where("Filename = ?", path).
		Offset((offset - 1) * limit).Limit(limit).First(&cv).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by path, err:", err.Error())
		return CV{}, err
	}

	err = constructCVArrayFields(&cv)
	if err != nil {
		log.Println("Failed to construct CV array fields, err:", err.Error())
		return CV{}, err
	}

	return cv, nil
}

func GetCVsByName(name string, limit int, offset int) ([]CV, error) {
	var cvs []CV
	err := Db.Model(&CV{}).Where("Name = ?", name).
		Offset((offset - 1) * limit).Limit(limit).Find(&cvs).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return []CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by name, err:", err.Error())
		return nil, err
	}

	for _, cv := range cvs {
		err = constructCVArrayFields(&cv)
		if err != nil {
			log.Println("Failed to construct CV array fields, err:", err.Error())
			return nil, err
		}
	}

	return cvs, nil
}

func GetCVsByDegree(degree string, limit int, offset int) ([]CV, error) {
	var cvs []CV
	err := Db.Model(&CV{}).Where("Degree = ?", degree).
		Offset((offset - 1) * limit).Limit(limit).Find(&cvs).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return []CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by degree, err:", err.Error())
		return nil, err
	}

	for _, cv := range cvs {
		err = constructCVArrayFields(&cv)
		if err != nil {
			log.Println("Failed to construct CV array fields, err:", err.Error())
			return nil, err
		}
	}

	return cvs, nil
}

func GetCVsGreaterThanWorkingYears(workingYears int, limit int, offset int) ([]CV, error) {
	var cvs []CV
	err := Db.Model(&CV{}).Where("WorkingYears >= ?", workingYears).
		Offset((offset - 1) * limit).Limit(limit).Find(&cvs).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return []CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by working years, err:", err.Error())
		return nil, err
	}

	for _, cv := range cvs {
		err = constructCVArrayFields(&cv)
		if err != nil {
			log.Println("Failed to construct CV array fields, err:", err.Error())
			return nil, err
		}
	}

	return cvs, nil
}

func SetCV(cv *CV) error {
	err := Db.Model(&CV{}).Save(cv).Error
	if err != nil {
		log.Println("Failed to save CV, err:", err.Error())
		return err
	}

	err = saveCVArrayFields(cv)
	if err != nil {
		log.Println("Failed to save CV array fields, err:", err.Error())
		return err
	}

	return nil
}

func DeleteCVByFilename(filename string) error {
	var cv CV
	err := Db.Model(&CV{}).Where("filename = ?", filename).First(&cv).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return err
	}
	if err != nil {
		log.Println("Failed to get CV by filename, err:", err.Error())
		return err
	}

	err = deleteCVArrayFields(&cv)
	if err != nil {
		log.Println("Failed to delete CV array fields, err:", err.Error())
		return err
	}

	err = Db.Delete(&cv, "Filename = ?", filename).Error
	if err != nil {
		log.Println("Failed to delete CV, err:", err.Error())
		return err
	}

	return nil
}
