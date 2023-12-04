package main

import (
	"flag"
	"fmt"
	"log"
	"os"
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

func applyFlags(opts *options) {
	f := options{}
	flag.IntVar(&f.day, "d", -1, "Override day")
	flag.IntVar(&f.year, "y", -1, "Override year")
	flag.StringVar(&f.session, "s", "", "Override session key")
	flag.Parse()

	if f.day != -1 {
		opts.day = f.day
	}
	if f.year != -1 {
		opts.year = f.year

		// Handle short year formats (Y2K YOLO)
		if year < 50 {
			opts.year += 2000
		} else if year < 99 {
			opts.year += 1900
		}
	}
	if f.session != "" {
		opts.session = f.session
	}
}

func applyEnv(opts *options) {
	file, err := os.ReadFile("./.env")
	if err != nil {
		if err == os.IsNotExist(err) {
			return //.env file is optional
		} else {
			panic(err)
		}
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

func applySystemTime(opts *options) {
	now := time.Now()
	fmt.Println(now)

	if now.Month() != 12 && (opts.day == 0 || opts.year == 0) {
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
