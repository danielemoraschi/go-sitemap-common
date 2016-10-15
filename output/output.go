package output


type OutputInterface interface {
    Write(data []byte) (int, error)
}

type Writer struct {
    OutputInterface
}