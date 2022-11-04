package accurateprevalent_test

import (
	"io"
	"pex-prevalent-colors-challenge/internal/app/models"
	"testing"

	"github.com/stretchr/testify/suite"
)

type accuratePrevalentColorTestSuite struct {
	suite.Suite
	URL            string
	PrevalentColor models.PrevalentColor
	FileReader     io.ReadCloser
	Color1         string
	Color2         string
	Color3         string
}

func (s *accuratePrevalentColorTestSuite) SetupTest() {
	s.URL = "https://www.google.com"
}

func TestAccuratePrevalentColorTestSuite(t *testing.T) {
	suite.Run(t, &accuratePrevalentColorTestSuite{})
}

func (s *accuratePrevalentColorTestSuite) Test_FetchImage() {
	s.Equal(s.URL, "https://www.google.com")
}
