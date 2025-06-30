package main

import (
	"flag"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"

	"gopkg.in/yaml.v3"
)

type AudioConfig struct {
	Bitrate    string `yaml:"bitrate"`
	SampleRate int    `yaml:"sample_rate"`
}

type Config struct {
	Audio AudioConfig `yaml:"audio"`
}

func loadConfig() (*Config, error) {
	config := &Config{
		Audio: AudioConfig{
			Bitrate:    "64k",
			SampleRate: 16000,
		},
	}

	configPaths := []string{
		"config.yaml",
		"./config.yaml",
		fmt.Sprintf("%s/.config/vem/config.yaml", os.Getenv("HOME")),
	}

	for _, path := range configPaths {
		if data, err := os.ReadFile(path); err == nil {
			if err := yaml.Unmarshal(data, config); err != nil {
				return nil, fmt.Errorf("error parsing config %s: %v", path, err)
			}
			break
		}
	}

	return config, nil
}

func main() {
	ss := flag.String("ss", "", "Start time (HH:MM:SS), optional")
	to := flag.String("to", "", "End time (HH:MM:SS), optional")
	inputFile := flag.String("i", "", "Input file")
	flag.Parse()

	config, err := loadConfig()
	if err != nil {
		fmt.Printf("Config error: %v\n", err)
		os.Exit(1)
	}

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
		"-ar", strconv.Itoa(config.Audio.SampleRate),
		"-b:a", config.Audio.Bitrate,
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
