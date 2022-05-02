package main

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strconv"
	"time"
)

func main() {
	clearTerminal()
	fmt.Printf(" %s[%s*%s] Emily's Club Name Checker%s\n\n", DEFAULT, CYAN, DEFAULT, FLUSH)
	collectTokens()
	collectNames()
	setup()
	checkLoop()
	time.Sleep(3 * time.Second)
	fmt.Printf("\n\n %s[%s*%s] Job Finished.%s\n\n", DEFAULT, CYAN, DEFAULT, FLUSH)
	fmt.Printf(" %s[%s!%s] Press Enter To Exit..%s", DEFAULT, CYAN, DEFAULT, FLUSH)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	os.Exit(0)
}

func clearTerminal() {
	if runtime.GOOS == "windows" {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	} else {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func collectTokens() {
	fmt.Printf("\r %s[%s*%s] Authorizing Check Tokens..%s", DEFAULT, CYAN, DEFAULT, FLUSH)
	checkTokenFile, err := readLines("auth/tokens.txt")
	if err != nil {
		fmt.Printf("\r %s[%s*%s] Failed to read token file.%s\n", DEFAULT, RED, DEFAULT, FLUSH)
		fmt.Printf("\n %s[%s!%s] Press Enter To Exit..%s", DEFAULT, CYAN, DEFAULT, FLUSH)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		os.Exit(0)
	}
	for _, token := range checkTokenFile {
		WaitGroup.Add(1)

		go authorizeTokens(token)
		time.Sleep(8 * time.Millisecond)
	}
	WaitGroup.Wait()
	if len(authorizedTokens) <= 0 {
		fmt.Printf("\r %s[%s*%s] No Tokens Authorized.%s\n", DEFAULT, RED, DEFAULT, FLUSH)
		fmt.Printf("\n %s[%s!%s] Press Enter To Exit..%s", DEFAULT, CYAN, DEFAULT, FLUSH)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		os.Exit(0)
	} else {
		fmt.Printf("\r %s[%s+%s] Checking Tokens Authorized: (%d)%s\n", DEFAULT, CYAN, DEFAULT, len(authorizedTokens), FLUSH)
	}
}

func collectNames() {
	fmt.Printf("\r %s[%s*%s] Collecting Clubs Names...%s", DEFAULT, CYAN, DEFAULT, FLUSH)
	nameFile, err := readLines("lists/club_names.txt")
	if err != nil {
		fmt.Printf("\r %s[%s*%s] Failed to read club names file.%s\n", DEFAULT, RED, DEFAULT, FLUSH)
		fmt.Printf("\n %s[%s!%s] Press Enter To Exit..%s", DEFAULT, CYAN, DEFAULT, FLUSH)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		os.Exit(0)
	}
	clubNames = nameFile
	fmt.Printf("\r %s[%s+%s] Club Names Collected: (%d) %s\n", DEFAULT, CYAN, DEFAULT, len(clubNames), FLUSH)
}

func setup() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("\n %s[%s?%s] Check Delay (in ms): ", DEFAULT, CYAN, DEFAULT)
	if scanner.Scan() {
		line := scanner.Text()
		value, err := strconv.Atoi(line)

		if err != nil {
			fmt.Printf(" %s[%s*%s] Failed to check user input.%s\n", DEFAULT, RED, DEFAULT, FLUSH)
			fmt.Printf("\n %s[%s!%s] Press Enter To Exit..%s", DEFAULT, CYAN, DEFAULT, FLUSH)
			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			os.Exit(0)
		}
		requestDelay = int(value)
		fmt.Println("")
	} else {
		fmt.Printf(" %s[%s*%s] Failed to check user input.%s\n", DEFAULT, RED, DEFAULT, FLUSH)
		fmt.Printf("\n %s[%s!%s] Press Enter To Exit..%s", DEFAULT, CYAN, DEFAULT, FLUSH)
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		os.Exit(0)
	}
}

func checkLoop() {
	for i := range clubNames {
		requests++
		go func() {
			reserve(clubNames[i], authorizedTokens[tokenCounter%len(authorizedTokens)])
		}()
		tokenCounter++
		time.Sleep(time.Duration(requestDelay) * time.Millisecond)
	}
}
