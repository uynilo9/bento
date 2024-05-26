package main

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/alexflint/go-arg"
	"github.com/joho/godotenv"

	"github.com/uynilo9/bento/pkg/logger"
)

var cmd = strings.Join(os.Args, " ")

type run struct {
	File string `arg:"positional" help:"the target file supposed to be run" placeholder:"<file>"`
}

type args struct {
	Run *run `arg:"subcommand:run" help:"run file by setting the argument <file>"`
	Version bool `arg:"--version,-v" help:"display the version and exit"`
	License bool `arg:"--license,-l" help:"display the license and exit"`
}

func (args) Description() string {
	return "🍱 Welcome to Bento " + os.Getenv("VERSION") + "\n"
}

func (args) Epilogue() string {
	return "✨ Visit " + os.Getenv("WEBSITE") + " to get more infomation about Bento"
}

func main() {
	if err := godotenv.Load(); err != nil {
		logger.Fatalf("Failed to find or load the file `.env` while running the command `%s`\n", cmd)
	}
	var args args
	goarg, err := arg.NewParser(arg.Config{Program: os.Args[0], Exit: os.Exit}, &args)
	if err != nil {
		logger.Fatal("Failed to create the argument parser\n")
	}
	goarg.Parse(os.Args[1:])
	switch {
		case args.Version:
			fmt.Println("🍱 Bento " + os.Getenv("VERSION"))
			os.Exit(0)
		case args.License:
			fmt.Println("📜 Apache License 2.0 Copyright " + os.Getenv("YEAR") + " @uynilo9")
			os.Exit(0)
		case args.Run != nil:
			file := args.Run.File
			if file != "" {
				if path, err := filepath.Abs(file); err != nil {
					logger.Fatalf("Failed to resolve the input file path `%s`\n", file)
				} else {
					if bytes, err := os.ReadFile(path); err != nil {
						if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
							logger.Fatalf("The input file `%s` doesn't exist\n", path)
						} else {
							logger.Fatalf("Failed to read the input file `%s`\n", path)
						}
					} else if filepath.Ext(file) != ".bento" {
						logger.Fatalf("`%s` isn't a legal Bento source file\n", path)
					} else {
						source := string(bytes)
						fmt.Println(source)
						os.Exit(0)
					}
				}
			} else {
				logger.Error("The argument <file> was missing while running the `run` subcommand\n")
				goarg.WriteHelp(os.Stdout)
			}
		default:
			goarg.WriteHelp(os.Stdout)
	}
}