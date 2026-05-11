package main

import (
	"embed"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"sync"
)

//go:embed embedded/*
var toolsFS embed.FS

const banner = `
   ╔══════════════════════════════════════════╗
   ║  403-Killchain – Unified 403 Bypass    ║
   ║  by Archsec-Emman (@Archsec-Emman)     ║
   ║  Arc-Cyber-Arsenal Edition             ║
   ║  https://github.com/Arc-Cyber-Arsenal/403-Killchain  ║
   ╚══════════════════════════════════════════╝
`

func main() {
	urlPtr := flag.String("u", "", "Target URL (e.g., https://example.com/admin)")
	listTools := flag.Bool("l", false, "List embedded tools and exit")
	onlyPtr := flag.String("only", "", "Run only the specified tool (nomore403,byp4xx,dontgo403,bypass403)")
	skipPtr := flag.String("skip", "", "Skip a tool (comma-separated)")

	flag.Usage = func() {
		fmt.Print(banner)
		fmt.Println("\nUsage:  403-killchain -u <URL>  [--only <tool>]  [--skip <tool>]")
		fmt.Println("Example: 403-killchain -u https://example.com/admin")
		fmt.Println("         Runs all four tools and prints a combined report.")
		fmt.Println("\nFlags:")
		flag.PrintDefaults()
	}

	flag.Parse()
	if *listTools {
		fmt.Println("Embedded tools:")
		fmt.Println("  - nomore403")
		fmt.Println("  - byp4xx")
		fmt.Println("  - dontgo403")
		fmt.Println("  - bypass-403.sh (Bash)")
		return
	}
	if *urlPtr == "" {
		flag.Usage()
		os.Exit(1)
	}

	tmpDir, err := ioutil.TempDir("", "403killchain")
	if err != nil {
		log.Fatalf("Cannot create temp dir: %v", err)
	}
	defer os.RemoveAll(tmpDir)

	tools := []string{"nomore403", "byp4xx", "dontgo403", "bypass-403.sh"}
	toolPaths := make(map[string]string)
	for _, name := range tools {
		data, err := toolsFS.ReadFile("embedded/" + name)
		if err != nil {
			log.Fatalf("Failed to read embedded tool %s: %v", name, err)
		}
		path := filepath.Join(tmpDir, name)
		err = ioutil.WriteFile(path, data, 0755)
		if err != nil {
			log.Fatalf("Failed to write %s: %v", name, err)
		}
		toolPaths[name] = path
	}

	runSet := make(map[string]bool)
	if *onlyPtr != "" {
		allowed := strings.Split(*onlyPtr, ",")
		for _, t := range allowed {
			t = strings.TrimSpace(t)
			if _, ok := toolPaths[t]; ok {
				runSet[t] = true
			}
		}
	} else {
		for _, t := range tools {
			runSet[t] = true
		}
	}
	if *skipPtr != "" {
		skipped := strings.Split(*skipPtr, ",")
		for _, t := range skipped {
			t = strings.TrimSpace(t)
			delete(runSet, t)
		}
	}

	fmt.Println(banner)
	fmt.Printf("[*] Target: %s\n\n", *urlPtr)

	var wg sync.WaitGroup
	output := make(map[string]string)
	var mu sync.Mutex

	for tool := range runSet {
		wg.Add(1)
		go func(tool string) {
			defer wg.Done()
			cmdPath := toolPaths[tool]
			var cmd *exec.Cmd
			if tool == "bypass-403.sh" {
				cmd = exec.Command("bash", cmdPath, *urlPtr)
			} else {
				cmd = exec.Command(cmdPath, *urlPtr)
			}
			cmd.Dir = tmpDir
			stdout, err := cmd.Output()
			if err != nil {
				mu.Lock()
				output[tool] = fmt.Sprintf("[!] %s error: %v", tool, err)
				mu.Unlock()
				return
			}
			mu.Lock()
			output[tool] = string(stdout)
			mu.Unlock()
		}(tool)
	}
	wg.Wait()

	fmt.Println("=== COMBINED REPORT ===")
	for _, tool := range tools {
		out, ok := output[tool]
		if !ok {
			continue
		}
		fmt.Printf("\n--- %s ---\n", tool)
		fmt.Print(out)
	}
	fmt.Println("\n=== DONE ===")
}
