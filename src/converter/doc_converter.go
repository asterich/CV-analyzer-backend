package converter

import (
	"encoding/json"
	"os"
	"os/exec"
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
			t, err = time.Parse("2006-1-2", yearMonth)
			if err != nil {
				t, err = time.Parse("2006/1/2", yearMonth)
			}
		}
	}
	return t.Year(), int(t.Month())
}

func ExtractDegreeFromEducations(educations []model.Education) string {
	degreeList := []string{"博士", "硕士", "学士", "本科", "专科", "大专", "中专", "高中", "初中", "小学"}
	for _, education := range educations {
		if degree := education.Degree; true {
			for _, degreeWord := range degreeList {
				if strings.Contains(degree, degreeWord) {
					return degreeWord
				}
			}
		}
		if major := education.Major; true {
			for _, degree := range degreeList {
				if strings.Contains(major, degree) {
					return degree
				}
			}
		}
		if school := education.School; true {
			for _, degree := range degreeList {
				if strings.Contains(school, degree) {
					return degree
				}
			}
			zhuankeWords := []string{"职业", "技术", "中专"}
			for _, zhuankeWord := range zhuankeWords {
				if strings.Contains(school, zhuankeWord) {
					return "中专"
				}
			}
			universityWords := []string{"大学", "学院"}
			for _, universityWord := range universityWords {
				if strings.Contains(school, universityWord) {
					return "本科"
				}
			}
			highSchoolWords := []string{"中"}
			for _, highSchoolWord := range highSchoolWords {
				if strings.Contains(school, highSchoolWord) {
					return "高中"
				}
			}
		}
	}

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
	err := exec.Command("python3", "scripts/convert_file.py", "--input_file", path).Run()
	if err != nil {
		return model.CV{}, err
	}

	var imageFiles []string
	// open image_files.json
	imageFilesJsonFile, err := os.Open("tmp/image_files.json")
	if err != nil {
		return model.CV{}, err
	}
	defer imageFilesJsonFile.Close()
	var imageFilesJson []byte
	imageFilesJsonFile.Read(imageFilesJson)
	json.Unmarshal(imageFilesJson, &imageFiles)
	for _, imageFile := range imageFiles {
		// TODO: OCR and SER the image file
		_ = imageFile
	}

	// return model.CV{}, nil
	// TODO: construct test data here
	return model.CV{}, nil
}

func ConvertDocToPositions(path string) ([]model.Position, error) {
	return make([]model.Position, 0), nil
}
