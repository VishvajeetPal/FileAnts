package utils

import (
	"context"
	"fmt"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/feature/s3/manager"
	"github.com/aws/aws-sdk-go-v2/service/s3"
	"github.com/joho/godotenv"
	"log"
	"os"
	"sync"
)

func AwsFileUpload(path string, fileName string) error {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	f, err := os.OpenFile(path, 0, 0777)
	defer f.Close()
	client := s3.NewFromConfig(cfg)
	uploader := manager.NewUploader(client)
	result, err := uploader.Upload(context.TODO(), &s3.PutObjectInput{
		Bucket: aws.String("healthmate"),
		Key:    aws.String(fileName),
		Body:   f,
	})

	if err != nil {
		fmt.Printf("%v\n", err)
		return err
	}
	fmt.Printf("%v", result.Location)
	return nil
}

func S3FileDownloader(fileNames ...string) error {

	bucket := "healthmate"
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		log.Printf("error: %v", err)
		return err
	}

	var wg sync.WaitGroup

	for _, filename := range fileNames {
		wg.Add(1)
		file := filename
		go func() {
			fmt.Printf("\nDownloding %v", file)
			f := fmt.Sprintf("./recordsTemp/download/%v", file)
			newFile, err := os.Create(f)
			if err != nil {
				log.Println(err)
				return
			}
			defer newFile.Close()

			client := s3.NewFromConfig(cfg)
			downloader := manager.NewDownloader(client)
			numBytes, err := downloader.Download(context.TODO(), newFile, &s3.GetObjectInput{
				Bucket: aws.String(bucket),
				Key:    aws.String(file),
			})
			if err != nil {
				os.Remove(f)
				log.Println(err)
				wg.Done()
				return
			}
			if numBytes == 0 {
				os.Remove(f)
				log.Println(err)
				wg.Done()
				return

			}
			fmt.Printf("\n\nDownloaded file %v", file)
			wg.Done()
		}()
	}
	wg.Wait()

	for _, file := range fileNames {
		res := IsFilePresent("./recordsTemp/download/" + file)
		if !res {
			return fmt.Errorf("%v File was not downloaded check Logs", file)
		}
	}
	return nil

}

func IsFilePresent(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return false
}

func DeleteFiles(path string) error {

	err := os.Remove(path)
	if err != nil {
		return err
	}
	return nil
}
