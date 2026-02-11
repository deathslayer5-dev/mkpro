package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	versionNo := "1.0.1"

	GREEN := "\x1b[32m"
	RED := "\x1b[31m"
	NC := "\x1b[0m"

	if len(os.Args) < 2 {
		fmt.Println(RED, "Expected mkpro <name> function", NC)
		os.Exit(-1)
	}
	options := map[string]bool{
		"help":      false,
		"version":   false,
		"java":      false,
		"project":   false,
		"git":       false,
		"git-local": false,
	}
	name := ""
	remote := ""
	next := false
	changed := false
	for _, arg := range os.Args[1:] {
		if next {
			remote = arg
			next = false
			changed = true
			continue
		}
		if arg == "--help" {
			options["help"] = true
			break
		}
		if arg == "--version" {
			options["version"] = true
			break
		}
		if arg == "--git" {
			options["git"] = true
			next = true
			continue
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
	if changed == false && next == true {
		log.Fatal(RED, "Missing remote name", NC)
	}
	if name == "" {
		log.Fatal(RED, "Missing mkpro function", NC)
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
	if strings.Count(name, ".") == 1 {
		cmd := exec.Command("touch", name)
		fmt.Printf("Creating file: %s\n", name)
		_, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully created file: %s%s\n", GREEN, name, NC)
		os.Exit(0)
	}
	cmd := exec.Command("mkdir", name)
	fmt.Printf("Creating directory: %s\n", name)
	_, err := cmd.Output()
	if err != nil {
		log.Fatal(RED, err, NC)
	}
	fmt.Printf("%sSuccessfully created: %s%s\n", GREEN, name, NC)

	if options["project"] {
		dir := fmt.Sprintf("%s%s", name, "/src")
		cmd := exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err := cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%s,Successfully created: %s%s\n", GREEN, dir, NC)

		dir = fmt.Sprintf("%s%s%s", name, "/src", "/resources")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%s,Successfully created: %s%s\n", GREEN, dir, NC)
		dir = fmt.Sprintf("%s%s%s", name, "/src", "/main")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully created: %s%s\n", GREEN, dir, NC)
	}
	if options["java"] {
		dir := fmt.Sprintf("%s%s", name, "/src")
		cmd := exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err := cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully created: %s%s\n", GREEN, dir, NC)

		dir = fmt.Sprintf("%s%s%s", name, "/src", "/main")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully created: %s%s\n", GREEN, dir, NC)

		dir = fmt.Sprintf("%s%s%s", name, "/src", "/resources")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully created: %s%s\n", GREEN, dir, NC)

		dir = fmt.Sprintf("%s%s", name, "/bin")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully created: %s%s\n", GREEN, dir, NC)

		dir = fmt.Sprintf("%s%s", name, "/.project_env")
		cmd = exec.Command("touch", "-p", dir)
		fmt.Printf("Creating file: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully created: %s%s\n", GREEN, dir, NC)

		cmd = exec.Command("echo", "\"alias", "javac='javac -d bin'\"", ">", dir)
		fmt.Printf("Writing to file: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully wrote to : %s%s\n", GREEN, dir, NC)
	}
	if options["git"] {
		cmd := exec.Command("git", "init")
		cmd.Dir = name

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("%sFailed to initialize git: %v\nOutput: %s%s", RED, err, output, NC)
		}
		if len(output) > 0 {
			fmt.Printf("git init output: %s", output)
		} else {
			fmt.Println(GREEN, "✓ Initialized git repository", NC)
		}

		fmt.Printf("Adding remote: %s\n", remote)
		cmd = exec.Command("git", "remote", "add", "origin", remote)
		cmd.Dir = name

		output, err = cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "already exists") {
				fmt.Println(RED, "✓ Remote 'origin' already exists", NC)
			} else {
				log.Fatalf("%sFailed to add remote: %v\nOutput: %s%s", RED, err, output, NC)
			}
		} else {
			if len(output) > 0 {
				fmt.Printf("%sRemote added: %s%s", GREEN, output, NC)
			} else {
				fmt.Println(GREEN, "✓ Added remote 'origin'", NC)
			}
		}
	}

	if options["git-local"] {
		cmd := exec.Command("git", "init")
		cmd.Dir = name

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatal(RED, err, NC)
		}
		fmt.Printf("%sSuccessfully initialized: %s%s\n", GREEN, output, NC)
	}
}
