package img_module // This.go file will download the image.

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func Getpic(imageUrl, fileName string) error {
	// URL of the image I downloaded.
	//imageUrl := "https://hgtvhome.sndimg.com/content/dam/images/hgtv/fullset/2011/7/18/0/HGTV_Color-Wheel-Full_s4x3.jpg.rend.hgtvcom.1280.1280.suffix/1400967008479.jpeg"

	// Create an HTTP GET request
	response, err := http.Get(imageUrl)
	if err != nil {
		return fmt.Errorf("error making the request: %v", err)
	}
	defer response.Body.Close()

	// Check if the response status code is OK (200)
	if response.StatusCode != http.StatusOK {
		return fmt.Errorf("error: status code %d", response.StatusCode)
	}

	// Create a new file to save the image
	outputFile, err := os.Create("downloaded_image.jpg")
	if err != nil {
		return fmt.Errorf("error creating the file: %v", err)
	}
	defer outputFile.Close()

	// Copy the HTTP response body to the file
	_, err = io.Copy(outputFile, response.Body)
	if err != nil {
		return fmt.Errorf("error saving the image: %v", err)
	}

	fmt.Println("Image downloaded and saved as 'downloaded_image.jpg'")
	return nil
}
