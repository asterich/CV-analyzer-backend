package utils

import (
	"log"

	gptapi "github.com/asterich/CV-analyzer-backend/src/utils/gpt-api"
	"gopkg.in/ini.v1"
)

var (
	// Server configs
	AppMode  string
	HttpPort string
	// JwtKey       string
	MaxLoginTime uint

	// Database configs
	Db     string
	DbName string
	DbPath string

	// Upload configs
	UploadPath string

	// GPT configs
	GptApi *gptapi.Api
	Model  gptapi.Model
	// // Redis configs
	// RedisAddr     string
	// RedisPassword string
	// RedisDB       int
)

// degrees
const (
	Bachelor = "本科"
	Master   = "硕士"
	Doctor   = "博士"
)

func init() {

	var file, err = ini.Load("config/config.ini")
	if err != nil {
		log.Fatalln("Failed to load config.ini, err:", err.Error())
	}
	LoadServer(file)
	LoadDb(file)
	LoadUpload(file)
	LoadGPT(file)
}

func LoadGPT(file *ini.File) {
	var gptSection = file.Section("gpt")
	// GPTKey = gptSection.Key("GPTKey").String()
	GptApi = gptapi.NewApi(gptSection.Key("GPTKey").String())
	Model = gptapi.Model(gptSection.Key("Model").String())
}

func LoadServer(file *ini.File) {
	var serverSection = file.Section("server")
	AppMode = serverSection.Key("AppMode").String()
	HttpPort = serverSection.Key("HttpPort").String()
	// JwtKey = serverSection.Key("JwtKey").String()
	MaxLoginTime = serverSection.Key("MaxLoginTime").MustUint()
}

func LoadDb(file *ini.File) {
	var dbSection = file.Section("database")
	Db = dbSection.Key("Db").String()
	DbName = dbSection.Key("DbName").String()
	DbPath = dbSection.Key("DbPath").String()
}

func LoadUpload(file *ini.File) {
	var uploadSection = file.Section("upload")
	UploadPath = uploadSection.Key("UploadPath").String()
	log.Printf("UploadPath: %s\n", UploadPath)
}

func GptGet(role, prompt string) string {
	request := &gptapi.Request{
		Model: Model,
		Messages: []*gptapi.Message{
			{
				Role:    role,
				Content: prompt,
			},
		},
	}

	response, err := GptApi.Chat(request)
	if err != nil {
		panic(err)
	}

	return response.Choices[0].Message.Content
}
