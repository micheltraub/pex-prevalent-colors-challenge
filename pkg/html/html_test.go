package html_test

import (
	"errors"
	"os"
	"pex-prevalent-colors-challenge/pkg/html"
	"testing"

	"github.com/stretchr/testify/suite"
)

type htmlTestSuite struct {
	suite.Suite
}

func (s *htmlTestSuite) SetupTest() {

}

func TestHtmlTestSuite(t *testing.T) {
	suite.Run(t, &htmlTestSuite{})
}

func (s *htmlTestSuite) TestCsv_OpenCsvOnBrowser() {
	err := html.CreateHtmlFromCsv("test/test.csv", "test/result.tmpl", "test/index.html")

	if err != nil {
		s.Fail(err.Error())
	}

	//read the file to check if data was written
	if _, err := os.Stat("test/index.html"); errors.Is(err, os.ErrNotExist) {
		s.Fail(err.Error())
	}
}
