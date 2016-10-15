package output

import "fmt"

type StOutWriter struct {
    *Writer
}

func StOutWriterFactory() OutputInterface {
    return StOutWriter{}
}

func (stOutWriter StOutWriter) Write(data []byte) (int, error) {
    return fmt.Printf(string(data))
}