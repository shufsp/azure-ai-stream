package compression

import (
	"fmt"
	"bytes"
	"os"
	"image/jpeg"
	"github.com/gen2brain/avif"
	"golang.org/x/image/bmp"
)

func AvifCompress(filename string, filenameOutput string, encodeOptions avif.Options, jpegQuality int) (string, error) {
	inFile, err := os.Open(filename)
	if err != nil {
	   return "", fmt.Errorf("failed to open BMP file: %v", err)
	}
	defer inFile.Close()

	img, err := bmp.Decode(inFile)
	if err != nil {
	   return "", fmt.Errorf("failed to decode BMP file: %v", err)
	}

	file, err := os.Create(filenameOutput)
	if err != nil {
		return "", err
	}

	err = avif.Encode(file, img, encodeOptions)
	if err != nil {
		return "", fmt.Errorf("Avif encoding failed for %s: %v", filename, err)
	}

	avifFileForReading, err := os.Open(filenameOutput)
	if err != nil {
		return "", fmt.Errorf("failed to open encoded AVIF file: %v", err)
	}
	defer avifFileForReading.Close()

	decodedImg, err := avif.Decode(avifFileForReading)
	if err != nil {
		return "", fmt.Errorf("failed to decode AVIF file: %v", err)
	}

	var buf bytes.Buffer
	err = jpeg.Encode(&buf, decodedImg, &jpeg.Options{
		Quality: jpegQuality,	
	})
	if err != nil {
		return "", fmt.Errorf("JPEG encoding failed: %v", err)
	}

	err = os.WriteFile(filenameOutput, buf.Bytes(), 0644)
	if err != nil {
		return "", fmt.Errorf("failed to write JPEG file: %v", err)
	}

	return filenameOutput, nil
}

func Lanzcos(filename string, filenameOutput string) (string, error) {
	return filenameOutput, nil
}
