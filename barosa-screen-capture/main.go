package main

import (
	"fmt"
	"os"
	"os/exec"
        "strings"
)

func CommandDisplayUsage() {
    fmt.Printf("Usage: %s <window title> <searchMethod> [filename]\n", os.Args[0])
}

func CommandRunLogic() {
    if len(os.Args) < 2 {
        CommandDisplayUsage()
        return
    }

    windowTitle := strings.TrimSpace(os.Args[1])
    if windowTitle == "" {
        CommandDisplayUsage()
        return
    }

    if len(os.Args) < 3 {
        CommandDisplayUsage()
        return
    }

    searchMethod := os.Args[2]  
    windowId, err := WindowGetId(windowTitle, searchMethod)
    if err != nil { 
        fmt.Printf("Error getting window ID for window title '%s': %v", windowTitle, err)
        return
    }

    filename := fmt.Sprintf("%s_screenshot", windowTitle)
    if len(os.Args) >= 4 {
       filename = os.Args[3]  // custom filename provided
    }

    err = WindowScreenshot(windowId, filename)
    if err != nil {
        fmt.Printf("Error screenshotting window '%s' (%s): %v", windowTitle, windowId, err)
        return
    }
    fmt.Printf("%s", filename)
}

func WindowGetId(windowTitle string, searchFlag string) (string, error) {
    cmd := exec.Command("bash", "-c",
        // they do it better, okay?
        fmt.Sprintf("xdotool search --onlyvisible --%s \"%s\" | head -n 1", searchFlag, windowTitle),  

        // head -n 1 in case theres more than one window with the similar string provided
        // we need to graob only the first id .. its not the best solution
    )

    output, err := cmd.Output()
    if err != nil {
        return "", fmt.Errorf("%v (stdout: \"%s\") (ensure window is visible!)", err, string(output))
    }
    return strings.TrimSpace(string(output)), nil
}


func WindowScreenshot(windowId string, filename string) (error) {
    // fuck her up.
    cmd := exec.Command("maim",
        "--window", windowId,
        "--quality", "1",
        "--hidecursor",
        filename)

    output, err := cmd.Output()
    if err != nil {
        os.Remove(filename)
        return fmt.Errorf("%v (stdout: \"%s\")", err, string(output))
    }
    return nil
}

func main() {
    CommandRunLogic()
}
