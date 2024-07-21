package main

import (
	"os"
	"testing"
)

func TestWindowGetId(t *testing.T) {
    bullshitWindowTitles := []string{
        "SOMETITLEEEEEJKKKSSJSKDJSKJSKJSK",
        "ddddddddddddddddddddddddddddddddgnndgnd",
        "*(#UI$HJ#n#*(*(#@(2))!)!___ -__!_)1== == -x092",
        "![[[!}}]}{!][1",
        "@))!)!)@(",
        "11111111111111111111",
        "f ff ff f f f f e e e",
    }

    var id string 
    var err error 

    for i := range len(bullshitWindowTitles) {
        id, err = WindowGetId(bullshitWindowTitles[i], "class")
        if len(id) > 0 {
            t.Errorf("Expected no id returned for getting window id with class '%s'. The id of this window is some how %s?!?!", bullshitWindowTitles[i], id)
        }
        id, err = WindowGetId(bullshitWindowTitles[i], "name")
        if len(id) > 0 {
            t.Errorf("Expected no id returned for getting window id with name '%s'. The id of this window is some how %s?!?!", bullshitWindowTitles[i], id)
        }
    }

    id, err = WindowGetId("a", "name")
    if err != nil && len(id) == 0 {
        t.Errorf("Expected to find the window id with a fuzzy name 'a'? perhaps?")
    }

    id, err = WindowGetId("a", "class")
    if err != nil && len(id) == 0 {
        t.Errorf("Expected to find the window id with a fuzzy class 'a'? perhaps?")
    }

    bullshitClassName := "klsdjflksdjf"
    id, err = WindowGetId("a", bullshitClassName)
    if len(id) > 0 {
        t.Errorf("Expected no id returned for getting window id 'a' with a bullshit class '%s'", bullshitClassName)
    }

    id, err = WindowGetId("", "")
    if len(id) > 0 {
        t.Errorf("Expected no id returned for blank args, but got %s", id)
    }

    id, err = WindowGetId("--", "")
    if len(id) > 0 {
        t.Errorf("Expected no id returned for blank args, but got %s", id)
    }

    id, err = WindowGetId("--rce", "")
    if len(id) > 0 {
        t.Errorf("Expected no id returned when passing arbitrary --flag to shit")
    }

    id, err = WindowGetId("-rce", "")
    if len(id) > 0 {
        t.Errorf("Expected no id returned when passing arbitrary -flag to shit")
    }
}

func TestWindowScreenshot(t *testing.T) {
    bullshitWindowIds := []string{
        "-99918239812",
        "sdjlf",
        "*(#UI$HJ#n",
        "![[[!}}]}{!][1",
        "@))!)!)@(",
        "11111111111111111111",
        "f ff ff f f f f e e e",
    }
    filename := "should_never_exist"
    for i := range len(bullshitWindowIds) {
        err := WindowScreenshot(bullshitWindowIds[i], filename)
        if err == nil {
            t.Errorf("Expected error for screenshotting window ID %s", bullshitWindowIds[i])
        }
        if _, err := os.Stat(filename); os.IsExist(err) {
            t.Errorf("Expected file to get deleted upon error!")
	}
    }

    realWindowId, err := WindowGetId("a", "name")
    if err != nil {
        t.Errorf("Expected at least one window id from fuzzy find on 'a'? maybe not..")
    }

    realFilename := "should_exist"
    err = WindowScreenshot(realWindowId, realFilename)
    defer os.Remove(realFilename)
    if err != nil {
        t.Errorf("Failed to screenshot real window!")
    }
    if _, err := os.Stat(realFilename); os.IsNotExist(err) {
        t.Errorf("Expected real filename to exist on successful screenshot! %v", err)
    }

    err = WindowScreenshot("--rce", realFilename)
    if err == nil {
        t.Errorf("Expected error when passing arbitrary bullshit --flag to window screenshot command.")
    }

    err = WindowScreenshot("-rce", realFilename)
    if err == nil {
        t.Errorf("Expected error when passing arbitrary bullshit -flag to window screenshot command.")
    }

    err = WindowScreenshot("", "")
    if err == nil {
        t.Errorf("Expected error when passing empty args")
    }

    err = WindowScreenshot("--", "")
    if err == nil {
        t.Errorf("Expected error when passing --")
    }
}
