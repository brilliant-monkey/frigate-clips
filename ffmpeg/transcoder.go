package ffmpeg

import (
	"fmt"
	"log"
	"os"
	"os/exec"
)

func Transcode(inputPath string, outputPath string, inputArgs []string, outputArgs []string) error {
	log.Println("Building FFMPEG command")
	generalArgs := []string{
		"-hide_banner",
		// "-loglevel", "error",
		"-y",
	}

	inputArgs = append(inputArgs, []string{
		"-i", inputPath,
	}...)

	encodeArgs := []string{
		// "-c:v", "h264_videotoolbox",
	}
	outputArgs = append(outputArgs, outputPath)

	ffmpegArgs := []string{}
	ffmpegArgs = append(ffmpegArgs, generalArgs...)
	ffmpegArgs = append(ffmpegArgs, inputArgs...)
	ffmpegArgs = append(ffmpegArgs, encodeArgs...)
	ffmpegArgs = append(ffmpegArgs, outputArgs...)

	cmd := exec.Command("ffmpeg", ffmpegArgs...)

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	stdin, err := cmd.StdinPipe() // Open stdin pipe
	if err != nil {
		log.Fatalln(err)
		return err
	}

	log.Printf("Starting FFMPEG on url %s.", inputPath)
	err = cmd.Start()
	if err != nil {
		log.Fatalln(err)
		return err
	}

	err = cmd.Wait()
	if err != nil {
		log.Println("Failed to wait for command to finish")
	}

	fmt.Println("Command completed")
	err = stdin.Close()
	if err != nil {
		log.Fatalln("Failed to close pipe")
		return err
	}

	fmt.Println("Conversion complete")

	return nil
}
