package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	s "strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

var files = []string{}
var destination = os.Getenv("to")
var processor = runtime.NumCPU()

func worker(id int, jobs <-chan string, results chan<- string) {
	for j := range jobs {
		convert(j)
		results <- j
	}
}

func main() {

	err := filepath.Walk("/convert", visit)
	check(err)

	jobs := make(chan string, len(files))
	results := make(chan string, len(files))

	for w := 1; w <= processor; w++ {
		go worker(w, jobs, results)
	}

	for _, entry := range files {
		jobs <- entry
	}
	close(jobs)

	for a := 1; a <= len(files); a++ {
		<-results
	}
}

func visit(p string, info os.FileInfo, err error) error {

	if err != nil {
		return err
	}

	if !info.IsDir() && !s.HasSuffix(p, destination) {
		files = append(files, p)
	}

	return nil
}

func convert(filename string) {

	println(filename)

	newFilename := substr(filename, 0, s.LastIndex(filename, ".")) + "." + destination

	cmd := exec.Command("ffmpeg", "-n", "-i", filename, newFilename)
	stdout, err := cmd.Output()

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	fmt.Print(string(stdout))
}

func substr(input string, start int, length int) string {
	asRunes := []rune(input)

	if start >= len(asRunes) {
		return ""
	}

	if start+length > len(asRunes) {
		length = len(asRunes) - start
	}

	return string(asRunes[start : start+length])
}
