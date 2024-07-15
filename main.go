package main

import (
	"fmt"
	"barosa.fun/shuferbarosa-ai/microphone"
)

func main() {
	fmt.Println("=================================");
	fmt.Println("          shuferbarosa ai");
	fmt.Println("          testing the scarab");
	fmt.Println("=================================");
	
	microphone.RecordToFile("shuferbarosa.wav", 44100)
}
