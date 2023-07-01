package converter

import "github.com/asterich/CV-analyzer-backend/src/model"

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

	return model.CV{}, nil
}

func ConvertDocToPositions(path string) ([]model.Position, error) {
	return make([]model.Position, 0), nil
}
