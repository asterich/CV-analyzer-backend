package model_test

import (
	"log"
	"os"
	"path/filepath"
	"testing"

	"github.com/asterich/CV-analyzer-backend/src/model"
)

func TestGetCountDegree(t *testing.T) {
	input := 50

	projectRootDir, _ := filepath.Abs("../")
	os.Chdir(projectRootDir)
	model.InitDb()

	mapp, err := model.GetCountDegree(input)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(mapp)
}

func TestGetCountWorkingyears(t *testing.T) {
	input := 50

	projectRootDir, _ := filepath.Abs("../")
	os.Chdir(projectRootDir)
	model.InitDb()

	mapp, err := model.GetCountWorkingyears(input)

	if err != nil {
		log.Fatalln(err)
	}
	log.Println(mapp)
}
