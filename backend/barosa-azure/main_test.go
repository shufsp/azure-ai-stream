package main

import (
	"os"
	"testing"
)


func TestImageLoadBinaryData(t *testing.T) {
	bullshitFiles := []string{
		"_))!@)(#*@()#2--.png",
		"lkdj.jpg.bwebm",
		"od99mdm,s. ..ds",
		"--rce",
		"d(( --)0123.p1np",
		"~",
		"",
		"; echo uhhhhhhhhhh",
	}

	for i := range len(bullshitFiles) {	
		buffer, err := ImageLoadBinaryData(bullshitFiles[i])
		if buffer != nil || err == nil {
			t.Errorf("Loading binary data from path %s should have failed. Instead, got %s", bullshitFiles[i], string(buffer))
		}
	}

	filename := "load_me_up"
	random_data := []byte{ '0', '6', '9', '2', '5' }
	err := os.WriteFile(filename, random_data, 0777)
	defer os.Remove(filename)

	if err != nil {
		t.Errorf("Failed to create test file %s: %v", filename, err)
	}

	buffer, err := ImageLoadBinaryData(filename)
	if err != nil {
		t.Errorf("Expected write to %s not to fail, but got error: %v", filename, err)
	}
	if len(buffer) != len(random_data) {
		t.Errorf("Expected random data buffer and retrieved file buffer to have same len")
	}
	for i := range(len(buffer)) {
		byte_a := buffer[i]
		byte_b := random_data[i]
		if byte_a != byte_b {
			t.Errorf("Expected buffer to be '%v', but was '%v'", string(random_data), string(buffer))
		}
	}
}
