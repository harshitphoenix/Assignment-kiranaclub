package utils

import (
	"fmt"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io"
	"net/http"
	"os"
	"path/filepath"
)

func GetImage(url string) (*os.File, error) {
	fmt.Println("GetImage")
	_, fileName := filepath.Split(url)
	fmt.Println("fileName", fileName)

	file, err := os.Create(fileName)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return nil, err
	}

	reader, err := os.Open(fileName)
	image, _, err := image.DecodeConfig(reader)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	fmt.Println("image", image)
	return file, nil
}
