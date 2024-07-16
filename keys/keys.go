package keys

import (
	"fmt"
        evdev "github.com/gvalkov/golang-evdev"
)

func WaitForKey(key uint16) error {
        // TODO need to be in sudo to read the global key but 
        // the microphone will stop working in sudo !
        // what the hell do we do?!

        // UPDATE ^^ you know what we do? we DONT! 
        keyboard, err := evdev.Open("/dev/input/event11") // Adjust event device as per your system
        if err != nil {
                return fmt.Errorf("error opening input device: %v", err)
        }

        for {
                events, err := keyboard.Read()
                if err != nil {
                        return fmt.Errorf("error reading input device: %v", err)
                }
                for _, ev := range events {
                        if ev.Type == evdev.EV_KEY && ev.Code == key && ev.Value == 1 {
                                return nil
                        }
                }
        }
}
