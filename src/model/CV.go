package model

import (
	"log"

	"gorm.io/gorm"
)

type ContactInfo struct {
	Tel   string `gorm:"type:varchar(20)" json:"tel"`
	Email string `gorm:"type:varchar(50)" json:"email"`
}

type Duration struct {
	Begin string `gorm:"type:varchar(20)" json:"start_time"`
	End   string `gorm:"type:varchar(20)" json:"end_time"`
}

type Education struct {
	ID       int `gorm:"type:int;primaryKey;autoIncrement"`
	CVId     int `gorm:"type:int" json:"cv_id"`
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
	ID         int `gorm:"type:int;primaryKey" json:"id"`
	CVId       int `gorm:"type:int" json:"cv_id"`
	Duration   `gorm:"embedded"`
	Experience `gorm:"embedded"`
}

type SchoolExperience struct {
	ID         int `gorm:"type:int;primaryKey" json:"id"`
	CVId       int `gorm:"type:int" json:"cv_id"`
	Experience `gorm:"embedded"`
}

type InternshipExperience struct {
	ID         int `gorm:"type:int;primaryKey" json:"id"`
	CVId       int `gorm:"type:int" json:"cv_id"`
	Experience `gorm:"embedded"`
}

type ProjectExperience struct {
	ID   int    `gorm:"type:int;primaryKey"`
	CVId int    `gorm:"type:int" json:"cv_id"`
	Name string `gorm:"type:varchar(50)" json:"project_name"`
	Desc string `gorm:"type:varchar(512)" json:"project_description"`
}

type Award struct {
	ID    int    `gorm:"type:int;primaryKey"`
	CVId  int    `gorm:"type:int" json:"cv_id"`
	Name  string `gorm:"type:varchar(50)" json:"award_name"`
	Level string `gorm:"type:varchar(20)" json:"level"`
}

type Skill struct {
	ID   int    `gorm:"type:int;primaryKey"`
	CVId int    `gorm:"type:int" json:"cv_id"`
	Name string `gorm:"type:varchar(32)" json:"skill_name"`
}

type CV struct {
	ID                    int    `gorm:"type:int;primaryKey"`
	Filename              string `gorm:"type:varchar(64)" json:"filename"`
	Name                  string `gorm:"type:varchar(16)" json:"name"`
	Age                   uint64 `gorm:"type:uint" json:"age"`
	Birthday              string `gorm:"type:varchar(20)" json:"birthday"`
	ContactInfo           `gorm:"embedded"`
	Degree                string                 `gorm:"type:varchar(20)" json:"degree"`
	WorkingYears          uint64                 `gorm:"type:uint" json:"working_years"` // WorkingYears is extracted from work experience
	Educations            []Education            `gorm:"-" json:"educations"`
	WorkExperiences       []WorkExperience       `gorm:"-" json:"work_experiences"`
	SchoolExperiences     []SchoolExperience     `gorm:"-" json:"school_experiences"`
	InternshipExperiences []InternshipExperience `gorm:"-" json:"internship_experiences"`
	ProjectExperiences    []ProjectExperience    `gorm:"-" json:"project_experiences"`
	Awards                []Award                `gorm:"-" json:"awards"`
	Skills                []Skill                `gorm:"-" json:"skills"`
	SelfDesc              string                 `gorm:"type:varchar(1024)" json:"self_description"`
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

func GetCVById(id int) (CV, error) {
	var cv CV
	err := Db.Model(&CV{}).Where("id = ?", id).First(&cv).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by id, err:", err.Error())
		return CV{}, err
	}

	err = constructCVArrayFields(&cv)
	if err != nil {
		log.Println("Failed to construct CV array fields, err:", err.Error())
		return CV{}, err
	}

	return cv, nil
}

func GetCVByFilename(path string, limit int, offset int) (CV, error) {
	var cv CV
	err := Db.Model(&CV{}).Where("filename = ?", path).
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

func GetCVLesserThanAge(age int, limit int, offset int) ([]CV, error) {
	var cvs []CV
	err := Db.Model(&CV{}).Where("age <= ?", age).
		Offset((offset - 1) * limit).Limit(limit).Find(&cvs).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return []CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by age, err:", err.Error())
		return nil, err
	}

	cvs_for_return := []CV{}
	for _, cv := range cvs {
		err = constructCVArrayFields(&cv)
		if err != nil {
			log.Println("Failed to construct CV array fields, err:", err.Error())
			return nil, err
		}
		cvs_for_return = append(cvs_for_return, cv)
	}

	return cvs_for_return, nil
}

func GetCVsByName(name string, limit int, offset int) ([]CV, error) {
	var cvs []CV
	err := Db.Model(&CV{}).Where("name = ?", name).
		Offset((offset - 1) * limit).Limit(limit).Find(&cvs).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return []CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by name, err:", err.Error())
		return nil, err
	}

	cvs_for_return := []CV{}
	for _, cv := range cvs {
		err = constructCVArrayFields(&cv)
		if err != nil {
			log.Println("Failed to construct CV array fields, err:", err.Error())
			return nil, err
		}
		cvs_for_return = append(cvs_for_return, cv)
	}

	return cvs_for_return, nil
}

func GetCVsByDegree(degree string, limit int, offset int) ([]CV, error) {
	var cvs []CV
	err := Db.Model(&CV{}).Where("degree = ?", degree).
		Offset((offset - 1) * limit).Limit(limit).Find(&cvs).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return []CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by degree, err:", err.Error())
		return nil, err
	}

	cvs_for_return := []CV{}
	for _, cv := range cvs {
		err = constructCVArrayFields(&cv)
		if err != nil {
			log.Println("Failed to construct CV array fields, err:", err.Error())
			return nil, err
		}
		cvs_for_return = append(cvs_for_return, cv)
	}

	return cvs_for_return, nil
}

func GetCVsGreaterThanWorkingYears(workingYears int, limit int, offset int) ([]CV, error) {
	var cvs []CV
	err := Db.Model(&CV{}).Where("working_years >= ?", workingYears).
		Offset((offset - 1) * limit).Limit(limit).Find(&cvs).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return []CV{}, err
	}
	if err != nil {
		log.Println("Failed to get CV by working years, err:", err.Error())
		return nil, err
	}

	cvs_for_return := []CV{}
	for _, cv := range cvs {
		err = constructCVArrayFields(&cv)
		if err != nil {
			log.Println("Failed to construct CV array fields, err:", err.Error())
			return nil, err
		}
		cvs_for_return = append(cvs_for_return, cv)
	}

	return cvs_for_return, nil
}

func GetAllCVs(limit int, offset int) ([]CV, error) {
	var CVs []CV
	result := Db.Model(&CV{}).Limit(limit).Offset((offset - 1) * limit).Find(&CVs)
	log.Println("CVs:", CVs)
	log.Println("error:", result.Error)
	return CVs, result.Error
}

func SetCV(cv *CV) error {
	err := Db.Model(&CV{}).Save(cv).Error
	if err != nil {
		log.Println("Failed to save CV, err:", err.Error())
		return err
	}

	// err = saveCVArrayFields(cv)
	// if err != nil {
	// 	log.Println("Failed to save CV array fields, err:", err.Error())
	// 	return err
	// }

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

func DeleteCVByID(id int) error {
	var cv CV
	err := Db.Model(&CV{}).Where("id = ?", id).First(&cv).Error
	if err == gorm.ErrRecordNotFound {
		log.Println("CV not found, err:", err.Error())
		return err
	}
	if err != nil {
		log.Println("Failed to get CV by id, err:", err.Error())
		return err
	}

	err = deleteCVArrayFields(&cv)
	if err != nil {
		log.Println("Failed to delete CV array fields, err:", err.Error())
		return err
	}

	err = Db.Delete(&cv, "id = ?", id).Error
	if err != nil {
		log.Println("Failed to delete CV, err:", err.Error())
		return err
	}

	return nil
}

func GetCountDegree(limit int) (map[string]int, error) {
	var degrees []string
	result := Db.Model(&Education{}).Limit(limit).Pluck("degree", &degrees)
	log.Println("degrees:", degrees)
	if result.Error != nil {
		log.Println("error:", result.Error)
	}
	degreeMap := make(map[string]int)
	for _, v := range degrees {
		if _, ok := degreeMap[v]; !ok {
			degreeMap[v] = 1
		} else {
			degreeMap[v] += 1
		}
	}
	return degreeMap, result.Error
}

func GetCountWorkingyears(limit int) (map[int]int, error) {
	var workingYears []int
	result := Db.Model(&CV{}).Limit(limit).Pluck("working_years", &workingYears)
	log.Println("working_years:", workingYears)
	if result.Error != nil {
		log.Println("error:", result.Error)
	}
	workingYearsMap := make(map[int]int)
	for _, v := range workingYears {
		if _, ok := workingYearsMap[v]; !ok {
			workingYearsMap[v] = 1
		} else {
			workingYearsMap[v] += 1
		}
	}
	return workingYearsMap, result.Error
}
