package models

import "path/filepath"

type DatabaseConfig struct {
	Local DatabaseInstanceConfig `yaml:"local"`
	Live  DatabaseInstanceConfig `yaml:"live"`
	Test  DatabaseInstanceConfig `yaml:"test"`
}
type DatabaseInstanceConfig struct {
	POSTGRES_USER     string `yaml:"POSTGRES_USER"`
	POSTGRES_DB       string `yaml:"POSTGRES_DB"`
	POSTGRES_PASSWORD string `yaml:"POSTGRES_PASSWORD"`
	POSTGRES_HOST     string `yaml:"POSTGRES_HOST"`
	POSTGRES_PORT     int    `yaml:"POSTGRES_PORT"`
	CertificateFolder string `yaml:"CERTIFICATE_FOLDER"`
}
type ServerConfig struct {
	//SERVER_PORT int    `yaml:"SERVER_PORT"`
	JWT_SECRET string `yaml:"JWT_SECRET"`
}
type GoogleStorageBucketConfig struct {
	Prod string `yaml:"live"`
	Stag string `yaml:"stag"`
	Dev  string `yaml:"dev"`
}
type GoogleConfig struct {
	CLOUD_PROJECT_ID              string                    `yaml:"CLOUD_PROJECT_ID"`
	CLOUD_SERVICEACCOUNT_JSONPATH string                    `yaml:"CLOUD_SERVICEACCOUNT_JSONPATH"`
	CLOUD_STORAGE_BUCKET_UPLOADS  GoogleStorageBucketConfig `yaml:"storage_buckets"`
}
type FirebaseConfig struct {
	SERVICEACCOUNT_JSONPATH string `yaml:"SERVICEACCOUNT_JSONPATH"`
}
type AppConfig struct {
	Server                       ServerConfig    `yaml:"server"`
	Autopilot                    AutopilotConfig `yaml:"autopilot"`
	Twilio                       TwilioConfig    `yaml:"twilio"`
	Google                       GoogleConfig    `yaml:"google"`
	Firebase                     FirebaseConfig  `yaml:"firebase"`
	Db                           DatabaseConfig  `yaml:"db"`
	BackendId                    string
	ConfigFolder                 string
	LiveHost                     string `yaml:"LIVE_HOST"`
	AddressForRedictedTestEmails string `yaml:"addressForRedictedTestEmails"`
}

func (c *AppConfig) GetPath(file string) string {
	return filepath.Join(c.ConfigFolder, file)
}

type AutopilotConfig struct {
	AUTOPILOTHQ_APIKEY string `yaml:"AUTOPILOTHQ_APIKEY"`
}
type TwilioConfig struct {
	API_KEY     string `yaml:"API_KEY"`
	Sid         string `yaml:"sid"`
	Token       string `yaml:"token"`
	From        string `yaml:"from"`
	Emailapikey string `yaml:"email_apikey"`
}
