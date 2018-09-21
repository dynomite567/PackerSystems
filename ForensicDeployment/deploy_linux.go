package main

import (
	"fmt"
	"os"
)

// Deploy - Function to handle any needed deployment of a forensic question
func Deploy(toDeploy Question) {
	if toDeploy.Deployment == "none" {
		var args = []string{"bash", "-c", "rm -rf /etc/gingertechengine/files"}
		var clean = getCommandOutput("sudo", args)
		fmt.Println(clean, "Done")
	} else if toDeploy.Deployment != "none" {
		var args = []string{"bash", "-c", toDeploy.Deployment}
		var deploy = getCommandOutput("sudo", args)
		args = []string{"bash", "-c", "rm -rf /etc/gingertechengine/files"}
		var clean = getCommandOutput("sudo", args)
		fmt.Println(clean, deploy)
	}

	if _, err := os.Stat("/home/administrator/Desktop/Forensic One.txt"); err == nil {
		createFile("/home/administrator/Desktop/Forensic Two.txt")
		AppendStringToFile("/home/administrator/Desktop/Forensic Two.txt", toDeploy.Question)
	} else {
		createFile("/home/administrator/Desktop/Forensic One.txt")
		AppendStringToFile("/home/administrator/Desktop/Forensic One.txt", toDeploy.Question)
	}

	key := []byte("WjNJKFcSZejKNzPP")
	answer, _ := encrypt(key, toDeploy.Answer)

	if _, err := os.Stat("/home/administrator/Desktop/Forensic Two.txt"); err == nil {
		AppendKeyToFile("/etc/gingertechengine/key1", "\n"+answer)
	} else {
		AppendKeyToFile("/etc/gingertechengine/key", "\n"+answer)
	}
}
