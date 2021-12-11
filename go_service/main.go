package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"
)

func main() {

	fmt.Println("==> ALL DONE WITH SETTING UP THE GO ENVIRONMENT.")

	go_location_msg := fmt.Sprintf("==> LOCATION: %s", os.Getenv("GOPATH"))
	fmt.Println(go_location_msg)

	cmd := exec.Command("go", "version")
	stdout, _ := cmd.Output()
	version_msg := fmt.Sprintf("==> VERSION: %s", string(stdout))
	fmt.Println(strings.Trim(version_msg, "\n"))

}
