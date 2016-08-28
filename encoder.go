package main

import (
	"io/ioutil"
	"log"
	"os/exec"
	"path/filepath"
	"strings"
)

func main() {
	baseCommand := "HandBrakeCLI"
	sourceDir := "/mnt/d/MoviesHQ/"
	targetDir := "/mnt/d/Movies/"
	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		log.Fatal(err)
		return
	}

	for _, file := range files {
		fileName := file.Name()
		sourceFile := filepath.Join(sourceDir, fileName)
		targetFile := filepath.Join(targetDir, fileName[0:len(fileName)-len(filepath.Ext(fileName))]+".mp4")
		cmd := exec.Command(baseCommand, "-i", sourceFile, "-t", "1", "-o", targetFile, "-f", "mp4", "-O", "-X", "1920", "--loose-anamorphic",
			"--modulus", "2", "-e", "x264", "-q", "19", "--vfr", "-a", "1", "-E", "av_aac", "-6", "dpl2", "-R", "Auto", "-B", "160", "--audio-fallback",
			"ac3", "--encoder-preset", "slower", "--encoder-tune", "film", "--encoder-level", "4.0", "--encoder-profile", "high", "--verbose", "1")
		println(strings.Join(cmd.Args, " "))

		err = cmd.Run()

		if err != nil {
			log.Fatal(err)
			return
		}
	}
}
