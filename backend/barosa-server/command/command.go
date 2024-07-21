package command

import (
	"strings"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	BAROSA_SCREENSHOT_BINARY = "barosa-screen-capture"
	BAROSA_AZURE_BINARY = "barosa-azure"
)

func CommandRunBarosaScreenshot(window string, method string) (string, error) {
	if _, err := os.Stat(BAROSA_SCREENSHOT_BINARY); os.IsNotExist(err) {
		log.Fatalf("%s binary is not present!", BAROSA_SCREENSHOT_BINARY)
		return "", err
	}

	cmd := exec.Command(fmt.Sprintf("./%s", BAROSA_SCREENSHOT_BINARY), window, method)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err	
	}

	return strings.TrimSpace(string(output)), nil
}

func CommandRunBarosaAzure(imageFilename string, features string) (string, error) {
	if _, err := os.Stat(BAROSA_AZURE_BINARY); os.IsNotExist(err) {
		log.Fatalf("%s binary is not present!", BAROSA_AZURE_BINARY)
		return "", err
	}

	if _, err := os.Stat(imageFilename); os.IsNotExist(err) {
		log.Fatalf("imageFilename '%s' doesnt exist", imageFilename)
		return "", err
	}

	cmd := exec.Command(fmt.Sprintf("./%s", BAROSA_AZURE_BINARY), imageFilename, features)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	return strings.TrimSpace(string(output)), nil
}
