package main

import (
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"time"
)

const (
	fileSize  = 128 * 1024 // 128KB
	numFiles  = 1000
	chunkSize = 4096 // 4KB
)

func main() {
	if len(os.Args) != 2 {
		fmt.Printf("Usage: %s <temp_folder_path>\n", os.Args[0])
		os.Exit(1)
	}

	tempFolder := os.Args[1]

	startTime := time.Now()

	// create subfolder for files
	timestamp := time.Now().Format("2006-01-02-150405")
	folderName := filepath.Join(tempFolder, fmt.Sprintf("nas-perf-test-%s", timestamp))
	if err := os.Mkdir(folderName, 0755); err != nil {
		fmt.Printf("Error creating folder: %s\n", err)
		os.Exit(1)
	}
	defer os.RemoveAll(folderName)

	// generate files and write them to filesystem
	for i := 0; i < numFiles; i++ {
		filename := fmt.Sprintf("%s/file_%d.txt", folderName, i)
		if err := os.MkdirAll(filepath.Dir(filename), 0755); err != nil {
			fmt.Printf("Error creating folder: %s\n", err)
			os.Exit(1)
		}
		file, err := os.Create(filename)
		if err != nil {
			fmt.Printf("Error creating file: %s\n", err)
			os.Exit(1)
		}

		// write file content in chunks
		for j := 0; j < fileSize/chunkSize; j++ {
			chunk := make([]byte, chunkSize)
			rand.Read(chunk)
			if _, err := file.Write(chunk); err != nil {
				fmt.Printf("Error writing file: %s\n", err)
				os.Exit(1)
			}
		}

		if err := file.Close(); err != nil {
			fmt.Printf("Error closing file: %s\n", err)
			os.Exit(1)
		}

		if (i+1)%100 == 0 {
			progress := float64(i+1) / float64(numFiles) * 100
			fmt.Printf("Wrote %.0f%% of files to filesystem\n", progress)
		}
	}
	writeDuration := time.Since(startTime)

	startTime = time.Now()

	// delete files from filesystem
	if err := os.RemoveAll(folderName); err != nil {
		fmt.Printf("Error deleting folder: %s\n", err)
		os.Exit(1)
	}

	deleteDuration := time.Since(startTime)

	fmt.Printf("Created %d files of size %dKB and wrote them to filesystem in %s\n", numFiles, fileSize/1024, writeDuration)
	fmt.Printf("Deleted %d files and folder from filesystem in %s\n", numFiles, deleteDuration)
}
