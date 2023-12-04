package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path"
	"strings"
)

func main() {
	opts := options{}
	applyEnv(&opts)
	applyFlags(&opts)
	applySystemTime(&opts)

	targetFolder := opts.getTargetFolder()

	if _, err := os.Stat(targetFolder); !os.IsNotExist(err) {
		fmt.Println("Folder already exists, delete folder to re-generate.")
	} else {
		createTargetFolderFromTemplate(targetFolder)
		replaceTemplatePlaceholders(targetFolder)
	}

	if opts.session != "" {
		downloadDailyInput(opts.day, opts.year, opts.session, targetFolder)
	} else {
		fmt.Println("Unable to retrieve input file (No session key), use flag -s or 'session' key in .env file to include your advent of code session key.")
	}
}

func createTargetFolderFromTemplate(dest string) {
	os.Mkdir(dest, os.ModePerm)

	if dirents, err := os.ReadDir("./cmd/create/template"); err != nil {
		panic(err)
	} else {
		for _, direct := range dirents {
			if direct.IsDir() {
				continue
			}

			data, err := os.ReadFile(path.Join("./cmd/create/template", direct.Name()))
			if err != nil {
				panic(err)
			}

			os.WriteFile(path.Join(dest, direct.Name()), data, os.ModePerm)
		}
	}
}

func downloadDailyInput(day, year int, session, targetFolder string) {
	var (
		sessionCookie = http.Cookie{Name: "session", Value: session}
		jar, _        = cookiejar.New(nil)
		client        = http.Client{
			Jar:       jar,
			Transport: &http.Transport{},
		}
	)

	req, err := http.NewRequest("GET", fmt.Sprintf("https://adventofcode.com/%v/day/%v/input", year, day), nil)
	if err != nil {
		panic(err)
	}
	client.Jar.SetCookies(req.URL, []*http.Cookie{&sessionCookie})

	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	dest := path.Join(targetFolder, "input.txt")
	file, err := os.OpenFile(dest, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(resp.Body)
	writer := bufio.NewWriter(file)
	defer writer.Flush()
	for {
		line, err := reader.ReadBytes('\n')
		writer.Write(line)

		if err != nil {
			if err == io.EOF {
				return
			}
			panic(err)
		}
	}
}

func replaceTemplatePlaceholders(folderName string) {
	fileName := path.Join(folderName, "main.go")

	lines := func() (lines []string) {
		file, err := os.Open(fileName)
		if err != nil {
			panic(err)
		}
		defer file.Close()

		reader := bufio.NewScanner(file)
		reader.Split(bufio.ScanLines)
		for reader.Scan() {
			lines = append(lines, reader.Text())
		}
		return
	}()

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	defer writer.Flush()

	for _, line := range lines {
		line = strings.ReplaceAll(line, "./_template", folderName) + "\n"
		writer.WriteString(line)
	}
}
