package simpleLogger

import (
	"fmt"
	"github.com/mattn/go-colorable"
	"os"
	"sync"
	"time"
)

type Logger struct {
	sem sync.Mutex
	DebugMode bool
	LogFile string
	fileMode bool
	fileSem sync.Mutex
}

func New(debug bool, logfile string) (Logger, error){
	fileMode := false
	if logfile != "" {
		file, err := os.OpenFile(logfile, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
		defer file.Close()
		if err != nil {
			return Logger{}, err
		}
		_, err = file.WriteString(time.Now().String() + ":log start!\n")
		if err != nil {
			return Logger{}, err
		}
		fileMode = true
	}
	return Logger {
		sem:       sync.Mutex{},
		DebugMode: debug,
		LogFile:   logfile,
		fileMode: fileMode,
		fileSem:   sync.Mutex{},
	}, nil
}

func (logger *Logger) Success(log string) {
	logger.StdOut(green("[+] "+log +"\n"))
	logger.log2file(log)
}

func (logger *Logger) Info(log string) {
	logger.StdOut(blue("[*] "+log+"\n"))
	logger.log2file(log)
}

func (logger *Logger) Warn(log string) {
	logger.StdOut(yellow("[-] "+log+"\n"))
	logger.log2file(log)
}

func (logger *Logger) Error(log string) {
	logger.StdOut(red("[×] "+log+"\n"))
	logger.log2file(log)
}

func (logger *Logger) Debug(log string) {
	if logger.DebugMode {
		logger.StdOut(Carmine("[~] "+log+"\n"))
	}
	logger.log2file(log)
}

func (logger *Logger) Out(log string) {
	logger.StdOut(green("[>] "+log+"\n"))
	logger.log2file("[+] "+log +"\n")
}

func (logger *Logger)StdOut(s string,a...interface{}) {
	logger.sem.Lock()
	stdout := colorable.NewColorable(os.Stdout)
	fmt.Fprintf(stdout,s,a...)
	logger.sem.Unlock()
}

func (logger *Logger)log2file(log string)() {
	if logger.fileMode {
		// 理论上这里不该有错误
		file, err := os.OpenFile(logger.LogFile, os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			fmt.Println("[×]open log file error.")
			return
		}
		file.WriteString(time.Now().String()+ " " + log + "\n")
	}
}