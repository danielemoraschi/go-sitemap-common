package output

import (
    "io/ioutil"
)


type FileWriter struct {
    *Writer
    fileName string
}

func FileWriterFactory(fileName string) OutputInterface {
    return FileWriter{fileName: fileName}
}

func (fileWriter FileWriter) Write(data []byte) (int, error) {
    return len(data), ioutil.WriteFile(fileWriter.fileName, data, 0644)
}