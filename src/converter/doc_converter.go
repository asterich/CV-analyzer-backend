package converter

import (
	"path/filepath"

	"github.com/asterich/CV-analyzer-backend/src/model"
)

func ExtractDegreeFromEducations(educations []model.Education) string {
	// TODO: extract degree from educations
	return ""
}

func ExtractWorkingYearsFromWorkExperiences(workExperiences []model.WorkExperience) int {
	// TODO: extract working years from work experiences
	return 0
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
					Begin: 16,
					End:   18,
				},
				School: "HUST",
			},
			{
				Duration: model.Duration{
					Begin: 16,
					End:   18,
				},
				School: "HUST",
			},
			{
				Duration: model.Duration{
					Begin: 16,
					End:   18,
				},
				School: "HUST",
			},
		},
		WorkExperiences: []model.WorkExperience{
			{
				Duration: model.Duration{
					Begin: 996,
					End:   114514,
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
