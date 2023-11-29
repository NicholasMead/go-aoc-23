package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"os"
	"path"
	"strings"
	"time"
)

type options struct {
	year, day int
	session   string
}

func (opts options) getTargetFolder() string {
	return fmt.Sprintf("./y%vd%v", opts.year-2000, opts.day)
}

func main() {
	opts := options{}
	applyEnv(&opts)
	applyFlags(&opts)
	applySystemTime(&opts)

	targetFolder := opts.getTargetFolder()

	copyTemplateFolder(targetFolder)

	if opts.session != "" {
		getInput(opts.day, opts.year, opts.session, targetFolder)
	} else {
		fmt.Println("Unable to retrieve input file (No session key), use flag -s or 'session' key in .env file to include your advent of code session key.")
	}

	replaceTemplateItems(targetFolder)
}

func applySystemTime(opts *options) {
	now := time.Now()

	if now.Month() != 12 && opts.day == 0 || opts.year == 0 {
		log.Fatalln("Its not advent-of-code month. Use arguments -y <year> -d <day> to specify which day you want to create.")
	}

	if opts.year == 0 {
		opts.year = now.Year()
	} else if opts.year < 2015 {
		log.Fatalf("Advent of code started in 2015, year value %v is out of range.\n", opts.year)
	}

	if opts.day == 0 {
		opts.day = now.Day()
	} else if opts.day <= 0 || opts.day > 25 {
		log.Fatalf("Advent of code runs 1st - 25th December, day value %v is out of range.\n", opts.day)
	}
}

func copyTemplateFolder(dest string) {
	os.Mkdir(dest, os.ModePerm)

	if dirents, err := os.ReadDir("./_template"); err != nil {
		panic(err)
	} else {
		for _, direct := range dirents {
			if direct.IsDir() {
				continue
			}

			data, err := os.ReadFile(path.Join("./_template", direct.Name()))
			if err != nil {
				panic(err)
			}

			os.WriteFile(path.Join(dest, direct.Name()), data, os.ModePerm)
		}
	}
}

func getInput(day, year int, session, targetFolder string) {
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

func replaceTemplateItems(folderName string) {
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

func applyFlags(opts *options) {
	f := options{}
	flag.IntVar(&f.day, "d", 0, "Override day")
	flag.IntVar(&f.year, "y", 0, "Override year")
	flag.StringVar(&f.session, "s", "", "Override session key")
	flag.Parse()

	if f.day != 0 {
		opts.day = f.day
	}
	if f.year != 0 {
		opts.year = f.year
	}
	if f.session != "" {
		opts.session = f.session
	}
}

func applyEnv(opts *options) {
	file, err := os.ReadFile("./.env")
	if err != nil {
		if err == os.ErrNotExist {
			return //.env file is optional
		}
		panic(err)
	}

	content := string(file)
	for _, line := range strings.Split(content, "\n") {
		if line == "" {
			continue
		}

		keyValue := strings.Split(line, "=")
		if len(keyValue) == 2 {
			switch keyValue[0] {
			case "session":
				opts.session = keyValue[1]
			default:
				fmt.Println("Unsupported .env key:", keyValue[0])
			}
		} else {
			panic("Invalid env file")
		}
	}
}
