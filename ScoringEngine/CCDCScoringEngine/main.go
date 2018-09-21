package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"time"
)

// Check - struct to format checks
type Check struct {
	Title    string `json:"title"`
	Command  string `json:"command"`
	Expected string `json:"expected"`
	Good     bool   `json:"good"`
}

func main() {
	os.Setenv("PATH", "/bin:/usr/bin:/sbin:/usr/local/bin")

	doEvery(1*time.Minute, check)
}

func check(t time.Time) {
	raw, err := ioutil.ReadFile("/etc/gingertechengine/checks.json")
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var c []Check
	json.Unmarshal(raw, &c)

	for i := 0; i < len(c); i++ {
		var args = []string{"/bin/bash", "-c", c[i].Command}
		var output = getCommandOutput("sudo", args)
		if output == c[i].Expected {
			c[i].Good = true
		} else if output != c[i].Expected {
			c[i].Good = false
		}
	}

	currentScore, _ := json.Marshal(c)
	err = ioutil.WriteFile("/opt/site/wwwroot/js/current.json", currentScore, 0644)
	fmt.Printf("%+v", c)
}

func getCommandOutput(command string, args []string) (output string) {
	cmd := exec.Command(command, args...)
	out, err := cmd.CombinedOutput()
	if err != nil {
		log.Fatalf("cmd.Run() failed with %s\n", err)
	}
	sha := string(out)

	return sha
}

// doEvery - Run function f every d length of time
func doEvery(d time.Duration, f func(time.Time)) {
	for x := range time.Tick(d) {
		f(x)
	}
}
