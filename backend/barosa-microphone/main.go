package main

import (
	"fmt"
	"log"
	"barosa.fun/shuferbarosa-ai/microphone"
)

const (
	AudioRecordingFilePath = "shuferbarosa.wav"
	AudioRecordingSampleRate = 44100
	AudioRecordingMaxSecondsOfSilence = 2
	AudioRecordingSilenceSampleThreshold = 110
)

func main() {

	fmt.Println("=================================");
	fmt.Println("          shuferbarosa ai");
	fmt.Println("          testing the scarab");
	fmt.Println("=================================");
	
	err := microphone.RecordToFile(AudioRecordingFilePath, 
		AudioRecordingSampleRate, 
		AudioRecordingMaxSecondsOfSilence,
		AudioRecordingSilenceSampleThreshold)

	if err != nil {	
		log.Fatalf("Unable to record microphone audio to file :( cuz %v\n", err)	
	}
}
