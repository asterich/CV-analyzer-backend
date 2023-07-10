package model

import "log"

func GetAllDegree(limit int, offset int) ([]string, error) {
	var degrees []string
	result := Db.Model(&Education{}).Limit(limit).Pluck("degree", &degrees).Offset((offset - 1) * limit)
	log.Println("degrees:", degrees)
	if result.Error != nil {
		log.Println("error:", result.Error)
	}
	return degrees, result.Error
}

func GetAllWorkingyears(limit int, offset int) ([]int, error) {
	var workingYears []int
	result := Db.Model(&CV{}).Limit(limit).Pluck("working_years", &workingYears).Offset((offset - 1) * limit)
	log.Println("working_years:", workingYears)
	if result.Error != nil {
		log.Println("error:", result.Error)
	}
	return workingYears, result.Error
}
