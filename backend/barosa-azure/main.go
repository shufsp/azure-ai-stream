package main

import (
	"fmt"
	"os"
	"io"
	"strings"
	"github.com/go-resty/resty/v2"
	"github.com/joho/godotenv"
)

type AzureVisionRestAPIConfig struct {
	key		string
	endpoint	string
	region		string
}

type AzureVisionRequestURLConfig struct {
	class		string	
	function	string
	apiVersion	string
	features	string // features values separated by ',' e.g. "denseCaptions,people,captions"
}

func AzureImageAnalyzeFeatures(filename string, visionConfig *AzureVisionRestAPIConfig, urlConfig *AzureVisionRequestURLConfig) (error) {
	imageBinaryData, err := ImageLoadBinaryData(filename)
	if err != nil {
		return fmt.Errorf("%v", err)
	}

	requestUrl := fmt.Sprintf("%s/computervision/%s:%s?api-version=%s&features=%s&language=en&model-version=latest&gender-neutral-caption=false", 
		visionConfig.endpoint,
		urlConfig.class,
		urlConfig.function,
		urlConfig.apiVersion,
		urlConfig.features)

	client := resty.New()
	response, err := client.R().
		SetBody(imageBinaryData). 
		SetHeader("Content-Type", "application/octet-stream"). 
		SetHeader("Ocp-Apim-Subscription-Key", visionConfig.key). 
		Post(requestUrl)
	if err != nil {
		return fmt.Errorf("%v", err)
	}
	
	fmt.Printf("%v", response.String())
	return nil
}

func ImageLoadBinaryData(filename string) ([]byte, error) {
	file, err := os.Open(filename)
	if err != nil {	
		return nil, fmt.Errorf("%v", err)
	}
	defer file.Close()
	return io.ReadAll(file)
}

func CommandDisplayUsage() {
	fmt.Printf("Usage: %s <image filename> <features separated by comma>", os.Args[0])
}

func main() {
	if len(os.Args) < 3 {
		CommandDisplayUsage()
		return
	}

	filename := strings.TrimSpace(os.Args[1])
	features := strings.TrimSpace(os.Args[2])

	err := godotenv.Load()
	if err != nil {
		fmt.Printf("Unable to load .env: %v\n", err)
	}

	azureVisionConfig := AzureVisionRestAPIConfig{
		key:		os.Getenv("VISION_KEY"),	
		endpoint:	os.Getenv("VISION_ENDPOINT"),
		region:		os.Getenv("VISION_REGION"),
	}

	azureVisionUrlConfig := AzureVisionRequestURLConfig{
		class:		"imageanalysis",	
		function:	"analyze",
		apiVersion:	"2023-10-01",
		features:	features,
	}

	err = AzureImageAnalyzeFeatures(filename, &azureVisionConfig, &azureVisionUrlConfig)
	if err != nil {
		fmt.Printf("Failed to analyze image file '%s': %v", filename, err)
	}
}
