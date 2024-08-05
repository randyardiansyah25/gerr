package gerr

import (
	"errors"
	"fmt"
	"runtime"
)

func New(msg string) error {
	file, line := getCallerInfo()
	erMsg := fmt.Sprintf("%s\n\t%s:%d", msg, file, line)
	return errors.New(erMsg)
}

func AddNew(er error, msg string) error {
	file, line := getCallerInfo()
	erMsg := fmt.Sprintf("%s\n\t%s:%d", msg, file, line)
	return errors.Join(er, errors.New(erMsg))
}

func getCallerInfo() (string, int) {
	// `2` Call stack level (1 is GetCallerInfo function, 2 for caller function)
	_, file, line, ok := runtime.Caller(2)
	if !ok {
		return "Unknown", 0
	}

	// Mengambil nama file saja dari path lengkap
	//file := filepath.Base(file)

	return file, line
}
