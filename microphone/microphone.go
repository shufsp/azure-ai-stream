package microphone

import (
	"fmt"
        "time"
        "os"
        "github.com/gordonklaus/portaudio"
	"github.com/youpy/go-wav"
)


func RecordToFile(file_path string, sampleRate uint32) error {
        //
        // Setup recording
        //
        err := portaudio.Initialize()
        if err != nil {
                return fmt.Errorf("Couldn't init portaudio: %v\n", err)
        }
        defer portaudio.Terminate()

        in := make([]int16, 1024)
        stream, err := portaudio.OpenDefaultStream(1, 0, 44100, len(in), in)
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

        // TODO record into samples buffer 
        stop := make(chan struct{})
        go func() {
                for {
                        select {
                        case <-stop:
                                return // Stop recording when stop signal is received
                        default:
                                stream.Read()
                                samples = append(samples, in...)
                        }
                }
        }()

        // TODO mechanism to stop recording after a certain duration of silence??
        <-time.After(3 * time.Second)
        close(stop)
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
