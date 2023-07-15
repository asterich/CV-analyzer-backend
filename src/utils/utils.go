package utils

import (
	"log"

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

	// // Redis configs
	// RedisAddr     string
	// RedisPassword string
	// RedisDB       int

	// OCR service configs
	OCRServerUrl string

	// LLM service configs
	LLMServerUrl         string
	LLMServerUrlPosition string
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
	LoadOCRService(file)
	LoadLLMService(file)
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

func LoadOCRService(file *ini.File) {
	var ocrSection = file.Section("ocr_service")
	OCRServerUrl = ocrSection.Key("OCRServerUrl").String()
	log.Printf("OCRServerAddr: %s\n", OCRServerUrl)
}

func LoadLLMService(file *ini.File) {
	var llmSection = file.Section("llm_service")
	LLMServerUrl = llmSection.Key("LLMServerUrl").String()
	LLMServerUrlPosition = llmSection.Key("LLMServerUrlPosition").String()
	log.Printf("LLMServerAddr: %s\n", LLMServerUrl)
}
