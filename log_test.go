package simpleLogger

import (
	"fmt"
	"testing"
)

func TestColor(t *testing.T) {
	fmt.Println(green("[+] green test"))
	fmt.Println(red("[×] red test"))
	fmt.Println(blue("[×] blue test"))
	fmt.Println(yellow("[×] yellow test"))
	fmt.Printf("%s,%s,%s,%s\n", green("test"), red("test"), blue("test"), yellow("test"))
}

func TestLogger1(t *testing.T) {
	var logger = Logger{}
	logger, err := New(true, "G:/xxx/")
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Debug("Debug test.")
	logger.Warn("Warn test.")
	logger.Error("Error test.")
	logger.Info("Info test.")
	logger.Success("Success test.")
}

func TestLogger2(t *testing.T) {
	var logger = Logger{}
	logger, err := New(true, "log.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	logger.Debug("Debug test.")
	logger.Warn("Warn test.")
	logger.Error("Error test.")
	logger.Info("Info test.")
	logger.Success("Success test.")
}