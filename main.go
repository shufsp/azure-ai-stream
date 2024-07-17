package main

import (
	"fmt"
	"log"
	azure "barosa.fun/shuferbarosa-ai/azureshit"
	"barosa.fun/shuferbarosa-ai/microphone"
	"github.com/joho/godotenv"
)

const (
	AudioRecordingFilePath = "shuferbarosa.wav"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	fmt.Println("=================================");
	fmt.Println("          shuferbarosa ai");
	fmt.Println("          testing the scarab");
	fmt.Println("=================================");
	
	err = microphone.RecordToFile(AudioRecordingFilePath, 44100)
	if err != nil {
		log.Fatalf("Unable to record microphone audio to file :( cuz %v\n", err)	
	}
	transcription, err := azure.AzureSpeechToText(AudioRecordingFilePath)
	if err != nil {	
		log.Fatalf("Unable to trancribe what u said to text :( cuz %v\n", err)	
	}
	fmt.Println(transcription)
}
