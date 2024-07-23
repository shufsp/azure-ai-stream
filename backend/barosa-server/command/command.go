package command

import (
	"strings"
	"fmt"
	"os"
	"os/exec"
)

// not const because unit test :) and idgaf about interface dp bullshhit. leave me alone
var BAROSA_SCREENSHOT_BINARY = "barosa-screen-capture"
var BAROSA_AZURE_BINARY = "barosa-azure"

func CommandRunBarosaScreenshot(window string, method string) (string, error) {
	if _, err := os.Stat(BAROSA_SCREENSHOT_BINARY); os.IsNotExist(err) {
		return "", fmt.Errorf("%s binary is not present!", BAROSA_SCREENSHOT_BINARY)

	}

	cmd := exec.Command(fmt.Sprintf("./%s", BAROSA_SCREENSHOT_BINARY), window, method)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err	
	}

	outputString := strings.TrimSpace(string(output))
	if output != nil && strings.Contains(outputString, "Error") {
		return "", fmt.Errorf("%s", outputString) 
	}

	return outputString, nil
}

func CommandRunBarosaAzure(imageFilename string, features string) (string, error) {
	if _, err := os.Stat(BAROSA_AZURE_BINARY); os.IsNotExist(err) {
		return "", fmt.Errorf("%s binary is not present!", BAROSA_AZURE_BINARY)

	}

	if _, err := os.Stat(imageFilename); os.IsNotExist(err) {	
		return "", fmt.Errorf("imageFilename '%s' doesnt exist", imageFilename)
	}

	cmd := exec.Command(fmt.Sprintf("./%s", BAROSA_AZURE_BINARY), imageFilename, features)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return "", err
	}

	outputString := strings.TrimSpace(string(output))
	if output != nil && strings.Contains(outputString, "Failed") {
		return "", fmt.Errorf("%s", outputString) 
	}

	return outputString, nil
}
