package main

import (
	"fmt"
	"os"
	"testing"

	"barosa.fun/azure-ai-stream-backend/command"
	"barosa.fun/azure-ai-stream-backend/environment"
)

func UtilCheckEnv(envKey string, minLength int, method func() error, checkValid bool) error {
	os.Setenv(envKey, "")
	err := method() 
	if err == nil {
		return fmt.Errorf("Expected fail when env has value of \"\"")
	}


	var key string = ""
	for range minLength - 1 {
		key = key + "a"
	}
	os.Setenv(envKey, key)
	err = method() 
	if err == nil {
		return fmt.Errorf("Expected fail when env has value of \"%s\" which is 1 less than the min length required of %d", key, environment.AUTH_ENV_KEY_MIN_LENGTH)
	}


	if !checkValid {
		return nil
	}

	var validAuthToken string = ""
	for range minLength + 1 {
		validAuthToken = validAuthToken + "a"
	}
	os.Setenv(envKey, validAuthToken)
	err = method() 
	if err != nil {
		return fmt.Errorf("Expected \"%s\" to be valid and not throw error (error was \"%v\")", validAuthToken, err)
	}
	return nil
}

func TestCheckAuthToken(t *testing.T) {
	err := UtilCheckEnv(environment.AUTH_ENV_KEY, environment.AUTH_ENV_KEY_MIN_LENGTH, environment.CheckAuthToken, true)
	if err != nil {
		t.Fatalf("%v", err)
	}
}


func TestCheckAuthSecret(t *testing.T) {
	err := UtilCheckEnv(environment.AUTH_SECRET_KEY, environment.AUTH_SECRET_KEY_MIN_LENGTH, environment.CheckAuthSecret, true)
	if err != nil {
		t.Fatalf("%v", err)
	}
}

func TestCheckClientPort(t *testing.T) {
	err := UtilCheckEnv(environment.BAROSA_CLIENT_PORT_KEY, 4, environment.CheckClientPort, false)
	if err != nil {
		t.Fatalf("%v", err)
	}

	os.Setenv(environment.BAROSA_CLIENT_PORT_KEY, "4444")
	err = environment.CheckClientPort()
	if err != nil {
		t.Fatalf("Expected successs with env value 4444, but got error: %v", err)
	}

	os.Setenv(environment.BAROSA_CLIENT_PORT_KEY, "44a44")
	err = environment.CheckClientPort()
	if err == nil {
		t.Fatalf("Expected failure with invalid port 44a44")
	}
}


func TestCommandRunBarosaAzure(t *testing.T) {
	actualBinary := command.BAROSA_AZURE_BINARY
	command.BAROSA_AZURE_BINARY = "slkdfjlksdfjlskdjflskdjflskdfjsoosdfosodooosodof"
	_, err := command.CommandRunBarosaAzure("somefile", "captions,people")
	if err == nil {
		t.Fatalf("Expected fail when setting azure binary to %s", command.BAROSA_AZURE_BINARY)
	}
	command.BAROSA_AZURE_BINARY = actualBinary

	
	_, err = command.CommandRunBarosaAzure("lsdjliofslidfjsdlifhsdlifhsdf", "captions,people")
	if err == nil {
		t.Fatalf("Expected fail when giving bullshit filename to send to azure")
	}
}


func TestCommandRunBarosaScreenshot(t *testing.T) {
	actualBinary := command.BAROSA_SCREENSHOT_BINARY
	command.BAROSA_SCREENSHOT_BINARY = "skdfjhskldjfhosdfhosduifhosdnhrfjskldnf"
	_, err := command.CommandRunBarosaScreenshot("some_window_doesnt_matter", "class")
	if err == nil {
		t.Fatalf("Expected fail when setting screenshot to %s", command.BAROSA_SCREENSHOT_BINARY)
	}
	command.BAROSA_SCREENSHOT_BINARY = actualBinary

	_, err = command.CommandRunBarosaScreenshot("this_window_probably_doesnt_exist", "class")

	if err == nil {
		t.Fatalf("Expected error when giving bullshit window class")
	}

	_, err = command.CommandRunBarosaScreenshot("plasma", "class")
	defer os.Remove("plasma_screenshot")
	if err != nil {
		t.Fatalf("Expected success screenshotting plasma window class.. dont tell me.. are u on windows? err.. no.. surely ur just a gnome user.. (%v)", err)
	}

	_, err = command.CommandRunBarosaScreenshot("plasma", "METHOD4")
	if err == nil {
		t.Fatalf("Expected error when giving bullshit method")
	}
}
