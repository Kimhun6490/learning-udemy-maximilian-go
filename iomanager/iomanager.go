package iomanager

type IOManager interface {
	ReadLinesFromFile() ([]string, error)
	WriteLinesToFile(data interface{}) error
}
