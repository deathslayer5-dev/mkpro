package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	versionNo := "1.0.0"

	if len(os.Args) < 2 {
		fmt.Println("Expected mkpro <name> function")
		os.Exit(1)
	}
	options := map[string]bool{
		"help":    false,
		"version": false,
		"java":    false,
		"project": false,
	}
	name := ""

	for _, arg := range os.Args[1:] {
		if arg == "--help" {
			options["help"] = true
			break
		}
		if arg == "--version" {
			options["version"] = true
			break
		}
		if !strings.HasPrefix(arg, "--") {
			name = arg
			break
		}
		if strings.HasPrefix(arg, "--") {
			str, _ := strings.CutPrefix(arg, "--")
			options[str] = true
		}
	}
	if name == "" {
		fmt.Println("Missing mkpro function")
	}
	if options["help"] {
		fmt.Println("Usage: mkpro \"--options\" <name>")
		for n := range options {
			fmt.Printf("\t%s\n", n)
		}
		os.Exit(0)
	}
	if options["version"] {
		fmt.Printf("Version: v%s\n", versionNo)
		os.Exit(0)
	}
	cmd := exec.Command("mkdir", name)
	fmt.Printf("Creating directory: %s\n", name)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Successfully created: %s\n", name)

	if options["project"] {
		dir := fmt.Sprintf("%s%s", name, "/src")
		cmd := exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)

		dir = fmt.Sprintf("%s%s", name, "/resources")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)
	}
	if options["java"] {
		dir := fmt.Sprintf("%s%s", name, "/src")
		cmd := exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)

		dir = fmt.Sprintf("%s%s", name, "/resources")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)

		dir = fmt.Sprintf("%s%s", name, "/bin")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)

		dir = fmt.Sprintf("%s%s", name, "/.project_env")
		cmd = exec.Command("touch", "-p", dir)
		fmt.Printf("Creating file: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)

		cmd = exec.Command("echo", "\"alias", "javac='javac -d bin'\"", ">", dir)
		fmt.Printf("Writing to file: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully wrote to : %s\n", dir)
	}
}
