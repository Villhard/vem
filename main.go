package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	ss := flag.String("ss", "", "Start time (HH:MM:SS), optional")
	to := flag.String("to", "", "End time (HH:MM:SS), optional")
	inputFile := flag.String("i", "", "Input file")
	flag.Parse()

	if *inputFile == "" {
		fmt.Println("Error: input file is required")
		os.Exit(1)
	}

	if !strings.HasSuffix(strings.ToLower(*inputFile), ".mov") {
		fmt.Println("Error: input file must have .mov extension")
		os.Exit(1)
	}

	inputDir := filepath.Dir(*inputFile)
	inputBase := strings.TrimSuffix(filepath.Base(*inputFile), filepath.Ext(*inputFile))
	outputFile := filepath.Join(inputDir, inputBase+".m4a")

	args := []string{
		"-hide_banner",
		"-y",
	}

	if *ss != "" {
		args = append(args, "-ss", *ss)
	}
	if *to != "" {
		args = append(args, "-to", *to)
	}

	args = append(args, "-i", *inputFile)
	args = append(args,
		"-vn",
		"-ar", "16000",
		"-b:a", "64k",
		"-c:a", "aac",
		outputFile,
	)

	cmd := exec.Command("ffmpeg", args...)

	fmt.Println("Converting...")
	if err := cmd.Run(); err != nil {
		fmt.Printf("ffmpeg error: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Done! Output: %s\n", outputFile)
}
