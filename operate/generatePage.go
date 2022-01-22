package operate

import (
	"fmt"
	"os"
)

func WriteToFile(msg string, fileNames ...string) {
	for _, fileName := range fileNames {
		writeFile(msg, fileName)
	}
	return
}

func writeFile(msg string, fileName string) {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		fmt.Println("file create failed. err: " + err.Error())
	} else {
		_, err := f.Write([]byte(msg))
		if err != nil {
			return
		}
		defer func(f *os.File) {
			err := f.Close()
			if err != nil {
				fmt.Println("close", fileName, "fail")
			}
		}(f)
	}
	return
}
