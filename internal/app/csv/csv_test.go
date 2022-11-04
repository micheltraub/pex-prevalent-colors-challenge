package csv_test

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

// Mock io interface
type WriteCloser struct {
	mock.Mock
}

func (WriteCloser) Close() error {
	return nil
}
func (_m *WriteCloser) Write(p []byte) (int, error) {
	return 0, nil
}

func TestCsv_OpenFileToAppend(t *testing.T) {
	//todo:	csv.OpenFileToAppend("test.csv")
}

func TestCsv_AppendToOsFile(t *testing.T) {
	//todo: csv.AppendToOsFile(&WriteCloser{}, []string{"test"})
}
func TestCsv_AppendToCsvFile(t *testing.T) {
	//todo:	csv.AppendToCsvFile("test.csv",, []string{"test"})
}
