package logger

import (
	"fmt"
	"os"

	"github.com/TwiN/go-color"
)

func Detail(a ...any) error {
	a = append([]any{ color.Gray + "DETAIL" + color.Reset + "\n" }, a...)
	_, err := fmt.Print(a...)
	return err
}

func Error(a ...any) error {
	a = append([]any{ color.Red + "ERROR" + color.Reset + " " }, a...)
	_, err := fmt.Print(a...)
	return err
}

func Fatal(a ...any) {
	a = append([]any{ color.Red + "FATAL" + color.Reset + " " }, a...)
	fmt.Print(a...)
	os.Exit(1)
}

func Tip(a ...any) error {
	a = append([]any{ color.Blue + "TIP" + color.Reset + " " }, a...)
	_, err := fmt.Print(a...)
	return err
}

func Detailf(format string, a ...any) error {
	_, err := fmt.Printf(color.Gray + "DETAIL" + color.Reset + "\n" + format, a...)
	return err
}

func Errorf(format string, a ...any) error {
	_, err := fmt.Printf(color.Red + "ERROR" + color.Reset + " " + format, a...)
	return err
}

func Fatalf(format string, a ...any) {
	fmt.Printf(color.Red + "FATAL" + color.Reset + " " + format, a...)
	os.Exit(1)
}

func Tipf(format string, a ...any) error {
	_, err := fmt.Printf(color.Blue + "TIP" + color.Reset + " " + format, a...)
	return err
}