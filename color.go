package simpleLogger

import "fmt"

func red(format string) string {
	return fmt.Sprintf("\x1b[31m"+format+"\x1b[0m")
}

func green(format string) string {
	return fmt.Sprintf("\x1b[32m"+format+"\x1b[0m")
}

func yellow(format string) string {
	return fmt.Sprintf("\x1b[33m"+format+"\x1b[0m")
}

func blue(format string) string {
	return fmt.Sprintf("\x1b[34m"+format+"\x1b[0m")
}

func Carmine(format string) string {
	return fmt.Sprintf("\x1b[35m"+format+"\x1b[0m")
}