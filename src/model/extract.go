package model

import "log"

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
