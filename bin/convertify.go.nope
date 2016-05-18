///usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strings"
)

func main() {
    wavs := ReadWavs()

    for _, wav := range wavs {
    	fmt.Printf("Processing: %s\n", wav)

	    command := BuildCommand(wav)
	    Run(command)
    }
}

func ReadWavs() []string {
    files, _ := ioutil.ReadDir("./build/wav/")
    wavs := []string{}
    for _, f := range files {
        if IsWav(f.Name()) {
        	wavs = append(wavs, f.Name())
        }
    }

    return wavs
}

func BuildCommand(wav string) string {
	return fmt.Sprintf(
		"ffmpeg -y -i build/wav/%s -c:a libmp3lame -q:a 2 build/dist/audio/%s -c:a libvorbis -q:a 4 build/dist/audio/%s",
		wav,
		strings.Replace(wav, ".wav", ".mp3", -1),
		strings.Replace(wav, ".wav", ".ogg", -1))
}

func IsWav(name string) bool {
	return strings.HasSuffix(name, ".wav")
}

func V(err error) {
	if err != nil {
		panic(err)
	}
}

func Run(command string) {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		V(err)
	}
}