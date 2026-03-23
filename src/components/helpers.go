package devops_scripts

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// GetS3Session returns an AWS S3 session
func GetS3Session(awsRegion string) (*session.Session, error) {
	sess, err := session.NewSession(&aws.Config{Region: aws.String(awsRegion)}, nil)
	if err != nil {
		return nil, err
	}
	return sess, nil
}

// GetS3Client returns an AWS S3 client
func GetS3Client(sess *session.Session) (*s3.S3, error) {
	s3Svc := s3.New(sess)
	return s3Svc, nil
}

// UploadFileToS3 uploads a file to S3
func UploadFileToS3(s3Svc *s3.S3, bucket, key, filePath string) error {
	_, err := s3Svc.PutObject(&s3.PutObjectInput{
		Bucket: aws.String(bucket),
		Key:    aws.String(key),
		Body:   os.Stdin,
	})
	return err
}

// GetEnvironmentVariable returns the value of an environment variable
func GetEnvironmentVariable(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatal("Environment variable not set:", name)
	}
	return val
}

// GetIntegerFromEnv returns the integer value of an environment variable
func GetIntegerFromEnv(name string) (int, error) {
	val := GetEnvironmentVariable(name)
	intVal, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		return 0, err
	}
	return int(intVal), nil
}

// GetBoolFromEnv returns the boolean value of an environment variable
func GetBoolFromEnv(name string) (bool, error) {
	val := GetEnvironmentVariable(name)
	boolVal, err := strconv.ParseBool(val)
	if err != nil {
		return false, err
	}
	return boolVal, nil
}

// GetRequiredEnvVar returns the value of an environment variable, logging a fatal error if it's not set
func GetRequiredEnvVar(name string) string {
	val := os.Getenv(name)
	if val == "" {
		log.Fatal("Environment variable not set:", name)
	}
	return val
}

// FormatString returns a formatted string using the provided formatting and arguments
func FormatString(f string, args ...interface{}) string {
	return fmt.Sprintf(f, args...)
}