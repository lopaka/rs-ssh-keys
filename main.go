package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/user"
	"regexp"
	"strings"
)

// Search policy file for username
func searchFile(file string, username string) string {

	contents, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	regex := fmt.Sprintf("^%s:|^[0-9a-z_\\-]*:%s:", username, username)
	re := regexp.MustCompile(regex)
	lines := strings.Split(string(contents), "\n")

	for _, line := range lines {
		if re.FindString(line) != "" {
			return (line)
		}
	}
	return ("")
}

func main() {

	keys := ""

	// Read in username argument
	args := os.Args
	if len(args) != 2 {
		fmt.Println("single username argument required")
		os.Exit(1)
	}
	username := args[1]

	// Read in policy file
	policyFile := "/var/lib/rightlink/login_policy"

	// found can be blank
	foundEntry := searchFile(policyFile, username)

	if foundEntry != "" {
		// Determine if the entry found has matching system UID.
		// If there is another user from another NSS plugin, this is not our user, so return no keys.
		policyFileUsername := strings.Split(foundEntry, ":")

		systemUsername, err := user.Lookup(username)
		// This should return at least our user. If nothing or error is returned, exit with no keys.
		if err != nil {
			os.Exit(1)
		}
		if policyFileUsername[3] == systemUsername.Uid {
			// User is from policyFile so get and set keys
			// Currently, keys are the in the 6th and on location in the array (starting with 0)
			for i := 6; i < len(policyFileUsername); i++ {
				keys = keys + policyFileUsername[i] + "\n"
			}
		}
	}

	fmt.Printf("%s", keys)
}
