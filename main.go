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

type flags struct {
	year, day int
	session   string
}

func main() {
	f := flags{}
	flag.IntVar(&f.day, "d", -1, "Override day")
	flag.IntVar(&f.year, "y", -1, "Override year")
	flag.StringVar(&f.session, "s", "", "Override session key")
	flag.Parse()

	day, year := parseDayYearFlags(f)

	folderName := fmt.Sprintf("./y%vd%v", year-2000, day)

	copyTemplateFolder(folderName)
	getInput(day, year, f.session, path.Join(folderName, "input.txt"))
	replaceTemplateItems(folderName)
}

func parseDayYearFlags(f flags) (day, year int) {
	now := time.Now()
	day, year = f.day, f.year

	if now.Month() != 12 && day == -1 || year == -1 {
		log.Fatalln("Its not advent-of-code month. Use arguments -y <year> -d <day> to specify which day you want to create.")
	}

	if year == -1 {
		year = now.Year()
	} else if year < 2015 {
		log.Fatalf("Advent of code started in 2015, year value %v is out of range.\n", year)
	}

	if day == -1 {
		day = now.Day()
	} else if day <= 0 || day > 25 {
		log.Fatalf("Advent of code runs 1st - 25th December, day value %v is out of range.\n", day)
	}

	return
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

func getInput(day, year int, session string, dest string) {
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
