package converter

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/asterich/CV-analyzer-backend/src/model"
)

func convertYearMonthStringToIntPair(yearMonth string) (int, int) {
	if strings.Contains(yearMonth, "至今") {
		return 2023, 4
	}
	if strings.Contains(yearMonth, "暑假") {
		yearMonth = strings.Replace(yearMonth, "暑假", "7月", -1)
	}
	t, err := time.Parse("2006年1月2日", yearMonth)
	if err != nil {
		t, err = time.Parse("2006.1.2", yearMonth)
		if err != nil {
			return 0, 0
		}
	}
	return t.Year(), int(t.Month())
}

func ExtractDegreeFromEducations(educations []model.Education) string {
	// TODO: extract degree from educations
	return ""
}

func ExtractWorkingYearsFromWorkExperiences(workExperiences []model.WorkExperience) int {
	totalWorkingMonths := 0
	for _, workExperience := range workExperiences {
		startYear, startMonth := convertYearMonthStringToIntPair(workExperience.Duration.Begin)
		endYear, endMonth := convertYearMonthStringToIntPair(workExperience.Duration.End)
		if startYear == 0 || endYear == 0 {
			continue
		}
		totalWorkingMonths += (endYear-startYear)*12 + (endMonth - startMonth)
	}
	return (totalWorkingMonths + 12) / 12
}

func ConvertDocToCV(path string) (model.CV, error) {

	// TODO: convert the file to CV

	// return model.CV{}, nil
	// TODO: construct test data here
	return model.CV{
		Filename: filepath.Base(path),
		Name:     "test",
		ContactInfo: model.ContactInfo{
			Email: "fuckers@gmail.com",
			Tel:   "1234567890",
		},
		Age:    22,
		Degree: "本科",
		Educations: []model.Education{
			{
				Duration: model.Duration{
					Begin: "",
					End:   "18",
				},
				School: "HUST",
			},
			{
				Duration: model.Duration{
					Begin: "16",
					End:   "18",
				},
				School: "HUST",
			},
			{
				Duration: model.Duration{
					Begin: "16",
					End:   "18",
				},
				School: "HUST",
			},
		},
		WorkExperiences: []model.WorkExperience{
			{
				Duration: model.Duration{
					Begin: "996",
					End:   "114514",
				},
				Experience: model.Experience{
					CompanyOrOrganization: "Google, LLC.",
					Position:              "CEO",
				},
			},
		},
	}, nil
}

func ConvertDocToPositions(path string) ([]model.Position, error) {
	return make([]model.Position, 0), nil
}
