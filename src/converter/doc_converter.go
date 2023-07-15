package converter

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strings"
	"time"

	"github.com/ledongthuc/pdf"

	"github.com/asterich/CV-analyzer-backend/src/model"
	"github.com/asterich/CV-analyzer-backend/src/utils"
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

func postRequest(url, body, contentType string) ([]byte, error) {
	resp, err := http.Post(url, contentType, strings.NewReader(body))
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	defer resp.Body.Close()
	log.Println(resp.Status)
	respBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
		return []byte{}, err
	}
	return respBody, nil
}

func removeKeysFromMap(m *map[string]interface{}, keys []string) {
	for _, key := range keys {
		delete(*m, key)
	}
}

func removePointsAndPredIDFromMap(m *map[string]interface{}) {
	removeKeysFromMap(m, []string{"points", "pred_id"})
}

func joinCVs(cvs []model.CV) model.CV {
	joinedCV := model.CV{}
	joinedCV.Educations = []model.Education{}
	joinedCV.WorkExperiences = []model.WorkExperience{}
	joinedCV.ProjectExperiences = []model.ProjectExperience{}
	joinedCV.InternshipExperiences = []model.InternshipExperience{}
	joinedCV.SchoolExperiences = []model.SchoolExperience{}
	joinedCV.Awards = []model.Award{}
	joinedCV.Skills = []model.Skill{}

	for _, cv := range cvs {
		if cv.Name != "" {
			joinedCV.Name = cv.Name
		}
		if cv.Age != 0 {
			joinedCV.Age = cv.Age
		}
		if cv.Birthday != "" {
			joinedCV.Birthday = cv.Birthday
		}
		if cv.Email != "" {
			joinedCV.Email = cv.Email
		}
		if cv.Tel != "" {
			joinedCV.Tel = cv.Tel
		}
		if cv.SelfDesc != "" {
			joinedCV.SelfDesc = cv.SelfDesc
		}
		if cv.WorkingYears != 0 {
			joinedCV.WorkingYears = cv.WorkingYears
		}
		if cv.Degree != "" {
			joinedCV.Degree = cv.Degree
		}

		joinedCV.Educations = append(joinedCV.Educations, cv.Educations...)
		joinedCV.WorkExperiences = append(joinedCV.WorkExperiences, cv.WorkExperiences...)
		joinedCV.ProjectExperiences = append(joinedCV.ProjectExperiences, cv.ProjectExperiences...)
		joinedCV.InternshipExperiences = append(joinedCV.InternshipExperiences, cv.InternshipExperiences...)
		joinedCV.SchoolExperiences = append(joinedCV.SchoolExperiences, cv.SchoolExperiences...)
		joinedCV.Awards = append(joinedCV.Awards, cv.Awards...)
		joinedCV.Skills = append(joinedCV.Skills, cv.Skills...)
	}
	return joinedCV
}

func ConvertDocToCV(path string) (model.CV, error) {
	resultCV := model.CV{}

	// TODO: convert the file to CV
	log.Printf("Converting file %s", path)
	var pythonCommandStr string
	if runtime.GOOS == "windows" {
		pythonCommandStr = "python"
	} else {
		pythonCommandStr = "python3"
	}
	absScriptPath, _ := filepath.Abs("scripts/convert_file.py")
	cmd := exec.Command(pythonCommandStr, absScriptPath, "--input_file", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		log.Printf("Failed to convert file %s, error = %s", path, err.Error()+": "+string(output))
		return model.CV{}, err
	}

	var imageFiles []string
	// open image_files.json
	absImageFilesJsonPath, _ := filepath.Abs("tmp/image_files.json")
	imageFilesJsonFile, err := os.Open(absImageFilesJsonPath)
	if err != nil {
		return model.CV{}, err
	}
	defer imageFilesJsonFile.Close()
	imageFilesJson := make([]byte, 1024)
	nbytes, err := imageFilesJsonFile.Read(imageFilesJson)
	if err != nil || nbytes == 0 {
		return model.CV{}, err
	}
	log.Printf("image_files.json = %s", string(imageFilesJson[:nbytes]))
	err = json.Unmarshal(imageFilesJson[:nbytes], &imageFiles)
	if err != nil {
		return model.CV{}, err
	}
	var resultCVs []model.CV
	for _, imageFile := range imageFiles {
		log.Printf("Processing image file %s", imageFile)
		imageData, err := ioutil.ReadFile(imageFile)
		if err != nil {
			return model.CV{}, err
		}
		encodedString := base64.StdEncoding.Strict().EncodeToString(imageData)
		// log.Printf("Encoded string = %s", encodedString)
		data := map[string]interface{}{
			"images": []string{encodedString},
		}
		dataJson, _ := json.Marshal(data)

		// post to ser server
		respJson, err := postRequest(utils.OCRServerUrl, string(dataJson), "application/json")
		if err != nil {
			return model.CV{}, err
		}
		// log.Println(respJson)
		var resp map[string][][][]map[string]interface{}
		json.Unmarshal(respJson, &resp)
		ocrResult := resp["results"]
		for _, ocrResultBox := range ocrResult[0][0] {
			removePointsAndPredIDFromMap(&ocrResultBox)
		}
		// log.Println(ocrResult[0][0])

		// post to llm server
		ocrResultJson, _ := json.Marshal(ocrResult[0][0])
		dataJson, _ = json.Marshal(map[string]string{
			"content": string(ocrResultJson),
		})
		var llmResp map[string]string
		var cv model.CV
		llmResultJson, err := postRequest(utils.LLMServerUrl, string(dataJson), "application/json")
		if err != nil {
			log.Println(err)
			return model.CV{}, err
		}
		// log.Println(string(llmResultJson))
		err = json.Unmarshal(llmResultJson, &llmResp)
		if err != nil {
			log.Println(err)
			return model.CV{}, err
		}
		// log.Println("llmResp content: " + llmResp["content"])
		err = json.Unmarshal([]byte(llmResp["content"]), &cv)
		if err != nil {
			log.Println(err)
			return model.CV{}, err
		}
		// fmt.Println(cv)
		resultCVs = append(resultCVs, cv)

	}

	resultCV = joinCVs(resultCVs)

	// resultCVJson, _ := json.Marshal(resultCV)
	// log.Println(string(resultCVJson))

	resultCV.WorkingYears = ExtractWorkingYearsFromWorkExperiences(resultCV.WorkExperiences)
	resultCV.Degree = ExtractDegreeFromEducations(resultCV.Educations)

	return resultCV, nil
}

func readPdf(path string) (string, error) {
	f, r, err := pdf.Open(path)
	if err != nil {
		return "", err
	}
	defer f.Close()

	var buf bytes.Buffer
	b, err := r.GetPlainText()
	if err != nil {
		return "", err
	}
	buf.ReadFrom(b)
	return buf.String(), nil
}

func readTxt(path string) (string, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil
}

func ConvertDocToPositions(path string) ([]model.Position, error) {
	var positions []model.Position
	var positionStr string
	var err error

	// first, examine the file type
	ext := filepath.Ext(path)
	switch ext {
	case ".pdf":
		positionStr, err = readPdf(path)
		if err != nil {
			return make([]model.Position, 0), err
		}
	case ".txt":
		positionStr, err = readTxt(path)
		if err != nil {
			return make([]model.Position, 0), err
		}
	default:
		log.Println("Unsupported file type:", path)
		return make([]model.Position, 0), fmt.Errorf("unsupported file type: %s", path)
	}

	// second, extract positions from the file
	dataJson, err := json.Marshal(map[string]interface{}{
		"content": positionStr,
	})
	if err != nil {
		log.Println(err)
		return make([]model.Position, 0), err
	}
	llmResultJson, err := postRequest(utils.LLMServerUrlPosition, string(dataJson), "application/json")
	if err != nil {
		log.Println(err)
		return make([]model.Position, 0), err
	}
	var llmResp map[string]string
	err = json.Unmarshal(llmResultJson, &llmResp)
	if err != nil {
		log.Println(err)
		return make([]model.Position, 0), err
	}
	err = json.Unmarshal([]byte(llmResp["content"]), &positions)
	if err != nil {
		log.Println(err)
		return make([]model.Position, 0), err
	}

	return positions, nil
}
