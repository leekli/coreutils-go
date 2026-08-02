package utils

import (
	"os"
	"syscall"
)

func CheckError(err error) {
	if err != nil {
		panic(err)
	}
}

func CheckIsStdIn() bool {
	// Checks if there is anything in standard input (stdin), or if there is an error reading stdin
    var stat syscall.Stat_t

    if err := syscall.Fstat(int(os.Stdin.Fd()), &stat); err != nil {
		return false
    }

    if (stat.Mode & syscall.S_IFIFO) != 0 || (stat.Mode & syscall.S_IFREG) != 0 {
        if stat.Size != 0 {
            return true
        }
    }

	return false
}

func HasArgs(args []string) bool {
	return len(args) > 0
}

func CheckFileExists (filePath string) bool {
	if _, err := os.Stat(filePath); err == nil {
        return true
    } else if os.IsNotExist(err) {
        return false
    } else {
        return false
    }
}

func ReadFile(filePath string) []byte {
	data, err := os.ReadFile(filePath)

	CheckError(err)

	return data
}