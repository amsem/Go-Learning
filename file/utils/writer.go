package utils

import "os"


func FileWriter(filename string, c string) error {
	return os.WriteFile(filename,[]byte(c), 0666)
}
