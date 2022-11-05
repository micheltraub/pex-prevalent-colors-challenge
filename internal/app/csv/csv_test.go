package csv_test

import (
	"os"
	"pex-prevalent-colors-challenge/internal/app/csv"
	"testing"

	"github.com/stretchr/testify/suite"
)

type csvTestSuite struct {
	suite.Suite
	AppendCsv csv.AppendCsv
}

func (s *csvTestSuite) SetupTest() {
	s.AppendCsv = csv.NewAppendCsvImpl("test.csv", []string{"test1", "test2"})
}

func TestCsvTestSuite(t *testing.T) {
	suite.Run(t, &csvTestSuite{})
}

func (s *csvTestSuite) TestCsv_OpenFileToAppend() {
	err := s.AppendCsv.AppendToCsvFile()
	if err != nil {
		s.Fail(err.Error())
	}
	//read the file to check if data was written
	read, _ := os.ReadFile("test.csv")
	str := string(read)
	//check if the correct string was written
	s.Equal("test1", str[0:5])
}
