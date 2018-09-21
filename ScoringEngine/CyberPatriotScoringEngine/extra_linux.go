package main

import (
	"strconv"
)

// Checks - Array of all the extra checks
var Checks = []Check{
	{ID: 1, Title: "Extended Checks", Description: "These are checks that are outside what I feel to be the \"Core\" checks.", BashCheck: "echo 1", Expected: "1"},
	{ID: 2, Title: "Thank you", Description: "Thank you for using this practice image! If you like it, you can support this project and others like it on Patrion at\nhttps://www.patreon.com/GingerTechnology", Function: test, Expected: "1"},
}

// ExtraChecks - Iterate through an array of all the checks that fall outside what I
// feel to be outside the "core" values of the CyberPatriot porttion of PackerSystems
func ExtraChecks() {

	var args = []string{}
	var output = ""
	for _, check := range Checks {
		if check.BashCheck != "" {
			args = []string{"bash", "-c", check.BashCheck}
			output = getCommandOutput("sudo", args)
			if output == check.Expected {
				Correct(check.ID, check.Title, check.Description)
			}
		} else {
			// I'm honestly not sure this works
			output = Call(check.Function)
			var compareTo = check.Expected
			if output == compareTo {
				Correct(check.ID, check.Title, check.Description)
			}
		}
	}
}

// Correct - If the check gives the correct result, append it to post
func Correct(id int, title string, description string) {
	AppendStringToFile("/etc/gingertechengine/post", title+" ("+strconv.Itoa(id)+"/"+strconv.Itoa(len(Checks))+")")
	AppendStringToFile("/etc/gingertechengine/post", "  - "+description)
	AppendStringToFile("/etc/gingertechengine/post", "")
}

// Call - I think it returns the output of the function passed to it?
func Call(f func() string) (output string) {
	var funcOut = f()
	return funcOut
}

// Here go the functions that an extra check may need.
// Make sure they return a string (return 1 or 0 for boolean results).

func test() (result string) {
	return "1"
}
