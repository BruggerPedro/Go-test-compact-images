package main

import (
	"image"
	"os"

	"github.com/chai2010/webp"
	"github.com/nfnt/resize"
)

func main() {
	compactWEBPImages()
}

func compactWEBPImages() error {
	inputFile, err := os.Open("input-images/Imagem-webP.webp")
	if err != nil {
		return err
	}
	defer inputFile.Close()

	img, _, err := image.Decode(inputFile)
	if err != nil {
		return err
	}

	if err := resizeImage(400, 300, img, "output-images/output1.webp"); err != nil {
		return err
	}

	if err := resizeImage(300, 200, img, "output-images/output2.webp"); err != nil {
		return err
	}

	return nil
}

func resizeImage(width int, height int, img image.Image, outputFilePath string) error {
	imgResize := resize.Resize(uint(width), uint(height), img, resize.Lanczos3)

	outputFile, err := os.Create(outputFilePath)
	if err != nil {
		return err
	}
	defer outputFile.Close()

	err = webp.Encode(outputFile, imgResize, nil)

	return err
}
