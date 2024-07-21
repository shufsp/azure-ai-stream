package microphone

import (
	"fmt"
        "os"
        "github.com/gordonklaus/portaudio"
	"github.com/youpy/go-wav"
)


func RecordToFile(file_path string, sampleRate uint32, silenceMaxDurationSeconds uint32, sampleSilenceThreshold int16) error {
        //
        // Setup recording
        //
        err := portaudio.Initialize()
        if err != nil {
                return fmt.Errorf("Couldn't init portaudio: %v\n", err)
        }
        defer portaudio.Terminate()

        in := make([]int16, 1024)
        stream, err := portaudio.OpenDefaultStream(1, 0, float64(sampleRate), len(in), in)
        if err != nil {
                return fmt.Errorf("error opening stream: %w", err)
        }
        defer stream.Close()


        //
        // Start recording
        //
        var samples []int16
        err = stream.Start()
        if err != nil {
                return fmt.Errorf("error starting stream: %w", err)
        }
        fmt.Print("Recording... ")

        stop := make(chan struct{})
        done := make(chan struct{})

        var sampleCounter uint32 = 0
        var silenceCounter uint32 = 0
        var secondsLeftUntilStop uint32 = silenceMaxDurationSeconds
        isTalkingStarted := false

        // its janky. but it works. therefore i do not care
        go func() {
        defer close(done)
        for {
                select {
                case <-stop:
                        return // Stop recording when stop signal is received
                default:
                        stream.Read()
                        
                        for _, sample := range in {
                                sampleCounter++

                                if sample < 0 {
                                        sample = -sample
                                }

                                if isTalkingStarted && sample < sampleSilenceThreshold {
                                        // silence occurring
                                        silenceCounter++
                                } else {
                                        // squawkin
                                        silenceCounter = 0
                                        isTalkingStarted = true
                                }

                                if silenceCounter % sampleRate == 0 {
                                        secondsLeftUntilStop = (((silenceMaxDurationSeconds * sampleRate) - silenceCounter) / sampleRate) 
                                }
                                

                                if sampleCounter % sampleRate == 0 {
                                        fmt.Printf("Stopping recording in %ds ...\n", secondsLeftUntilStop)
                                }

                                if silenceCounter > silenceMaxDurationSeconds * sampleRate {
                                        // silent for too long
                                        close(stop) 
                                        return
                                }
                        }
                        
                        samples = append(samples, in...)
                }
        }
        }()

        <-done
        defer stream.Stop()
        fmt.Println("\nRecording stopped")


        //
        // Save recording
        //
        file, err := os.Create(file_path)
        if err != nil {
                return fmt.Errorf("error creating WAV file: %w", err)
        }
        defer file.Close()
	numSamples := uint32(len(samples))
	w := wav.NewWriter(file, numSamples, 1, sampleRate, 16)

        // convert int16 buffer to bytes
        data := make([]byte, len(samples)*2) // 2 bytes per sample
        for i, sample := range samples {
                data[i*2] = byte(sample)
                data[i*2+1] = byte(sample >> 8) // Little-endian byte order
        }
        w.Write(data)

        return nil
}
