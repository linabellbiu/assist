package assist

import (
	"log"
	"os"
)

func WriteToFileString(path, msg string) {
	f, err := os.OpenFile(path, os.O_WRONLY&os.O_CREATE, 0666)
	if err != nil {
		log.Println(err.Error())
	}

	_, err = f.Write([]byte(msg))
	if err != nil {
		log.Println(err.Error())
	}
	f.Close()
}
