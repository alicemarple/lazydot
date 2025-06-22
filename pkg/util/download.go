package util

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// download using the newURL

func DownloadFromURL(url, filename string) error {
	resp, err := http.Get(url)
	if err != nil {
		return fmt.Errorf("error during GET request: %w", err)
	}
	defer resp.Body.Close()

	// 2️⃣ Check HTTP status
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("bad status: %s", resp.Status)
	}

	// 3️⃣ Create the destination file
	out, err := os.Create(filename)
	if err != nil {
		return fmt.Errorf("failed to create file: %w", err)
	}
	defer out.Close()

	// 4️⃣ Copy response body to the file
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return fmt.Errorf("error while downloading: %w", err)
	}

	fmt.Printf("File saved to %s\n", filename)
	return nil
}
