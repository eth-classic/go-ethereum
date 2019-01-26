package logger

import (
	"fmt"
	"io"
	"os"

	"github.com/ether-core/go-ethereum/common"
)

func openFile(datadir string, filename string) *os.File {
	path := common.EnsurePathAbsoluteOrRelativeTo(datadir, filename)
	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(fmt.Sprintf("error opening log file '%s': %v", filename, err))
	}
	return file
}

func New(datadir string, logFile string, logLevel int, flags int) LogSystem {
	var writer io.Writer
	if logFile == "" {
		writer = os.Stdout
	} else {
		writer = openFile(datadir, logFile)
	}

	var sys LogSystem
	sys = NewStdLogSystem(writer, flags, LogLevel(logLevel))
	AddLogSystem(sys)

	return sys
}

func BuildNewMLogSystem(datadir string, logFile string, logLevel int, flags int, withTimestamp bool) LogSystem {
	var writer io.Writer
	if logFile == "" {
		writer = os.Stdout
	} else {
		writer = openFile(datadir, logFile)
	}

	var sys LogSystem
	sys = NewMLogSystem(writer, flags, LogLevel(logLevel), withTimestamp)
	AddLogSystem(sys)

	return sys
}

func NewJSONsystem(datadir string, logFile string) LogSystem {
	var writer io.Writer
	if logFile == "-" {
		writer = os.Stdout
	} else {
		writer = openFile(datadir, logFile)
	}

	var sys LogSystem
	sys = NewJsonLogSystem(writer)
	AddLogSystem(sys)

	return sys
}

const (
	reset   = "\x1b[39m"
	green   = "\x1b[32m"
	blue    = "\x1b[36m"
	yellow  = "\x1b[33m"
	red     = "\x1b[31m"
	magenta = "\x1b[35m"
)

func ColorGreen(s string) (coloredString string) {
	return green + s + reset
}
func ColorRed(s string) (coloredString string) {
	return red + s + reset
}
func ColorBlue(s string) (coloredString string) {
	return blue + s + reset
}
func ColorYellow(s string) (coloredString string) {
	return yellow + s + reset
}
func ColorMagenta(s string) (coloredString string) {
	return magenta + s + reset
}
