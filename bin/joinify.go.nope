///usr/bin/env go run $0 $@ ; exit

package main

import (
	"fmt"
	"os/exec"
	"strings"
)

func main() {
    durations := []string{
    	"fivedays",
    	"oneweek",
    	"threedays",
    }

    sinces := []string{
    	"thelivingroom",
    	"youlaughedatme",
    	"youlookedatme",
    }
	
    for _, duration := range durations {
    	for _, since := range sinces {
    		fmt.Printf("Processing: oneweek, %s, %s\n", duration, since)
    		command := BuildCommand(duration, since)
    		Run(command)
    	}
    }
}

func V(err error) {
	if err != nil {
		panic(err)
	}
}

func BuildCommand(duration string, since string) string {
	output := fmt.Sprintf("itsbeen-%s-%s", duration, since)
	command := fmt.Sprintf(
		"ffmpeg -i assets/itsbeen.wav -i assets/duration-%s.wav -i assets/since-%s.wav -filter_complex [0:0][1:0][2:0]concat=n=3:v=0:a=1[out] -map [out] -y build/wav/%s.wav",
		duration,
		since,
		output)
	
	return command
}

func Run(command string) {
	args := strings.Split(command, " ")
	cmd := exec.Command(args[0], args[1:]...)
	if err := cmd.Run(); err != nil {
		fmt.Printf("command failed: %s\n", command)
		V(err)
	}
}