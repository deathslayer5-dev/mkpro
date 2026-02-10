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
		log.Fatal("Missing remote name")
	}
	if name == "" {
		log.Fatal("Missing mkpro function")
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

		dir = fmt.Sprintf("%s%s%s", name, "src", "/resources")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)
		dir = fmt.Sprintf("%s%s%s", name, "src", "/main")
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

		dir = fmt.Sprintf("%s%s%s", name, "/src", "/main")
		cmd = exec.Command("mkdir", "-p", dir)
		fmt.Printf("Creating directory: %s\n", dir)
		_, err = cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully created: %s\n", dir)

		dir = fmt.Sprintf("%s%s%s", name, "/src", "/resources")
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
	if options["git"] {
		cmd := exec.Command("git", "init")
		cmd.Dir = name

		output, err := cmd.CombinedOutput()
		if err != nil {
			log.Fatalf("Failed to initialize git: %v\nOutput: %s", err, output)
		}
		if len(output) > 0 {
			fmt.Printf("git init output: %s", output)
		} else {
			fmt.Println("✓ Initialized git repository")
		}

		fmt.Printf("Adding remote: %s\n", remote)
		cmd = exec.Command("git", "remote", "add", "origin", remote)
		cmd.Dir = name

		output, err = cmd.CombinedOutput()
		if err != nil {
			if strings.Contains(string(output), "already exists") {
				fmt.Println("✓ Remote 'origin' already exists")
				fmt.Println("Replacing 'origin'...")
			} else {
				log.Fatalf("Failed to add remote: %v\nOutput: %s", err, output)
			}
		} else {
			if len(output) > 0 {
				fmt.Printf("Remote added: %s", output)
			} else {
				fmt.Println("✓ Added remote 'origin'")
			}
		}
	}

	if options["git-local"] {
		cmd := exec.Command("git", "init")
		cmd.Dir = name

		output, err := cmd.Output()
		if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("Successfully initialized: %s\n", output)
	}
}
