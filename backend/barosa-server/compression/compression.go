package compression

import (
	"fmt"
	"github.com/disintegration/imaging"
	"github.com/gen2brain/avif"
	"golang.org/x/image/bmp"
	"image"
	"image/color"
	"os"
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

	return filenameOutput, nil
}

func Lanzcos(filename string, width int, filenameOutput string) (string, error) {
	src, err := imaging.Open(filename)
	if err != nil {
		return "", fmt.Errorf("failed to open image: %v", err)
	}

	// preserves aspect ratio
	method := imaging.Resize(src, width, 0, imaging.Lanczos)
	dst := imaging.New(width, method.Rect.Dy(), color.NRGBA{0, 0, 0, 0})
	dst = imaging.Paste(dst, method, image.Pt(0, 0))

	err = imaging.Save(dst, filenameOutput)
	if err != nil {
		return "", fmt.Errorf("failed to save image: %v", err)
	}
	return filenameOutput, nil
}
