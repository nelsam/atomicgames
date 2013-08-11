package ui

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

type UiHandler struct {
	baseFileName string
}

func (self UiHandler) ServeHTTP(writer http.ResponseWriter,
	request *http.Request) {
	baseFile, err := os.Open(self.baseFileName)
	if err != nil {
		fmt.Fprint(writer, err)
	} else {
		defer baseFile.Close()

		buffer := make([]byte, 1024)
		for {
			count, err := baseFile.Read(buffer)
			writer.Write(buffer[:count])
			if err == io.EOF {
				break
			}
		}
	}
}

var Handler = UiHandler{"ui/index.html"}
