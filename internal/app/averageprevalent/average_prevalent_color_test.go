package averageprevalent_test

import (
	"image"
	"pex-prevalent-colors-challenge/internal/app/averageprevalent"
	"pex-prevalent-colors-challenge/internal/app/models"
	"testing"

	"github.com/stretchr/testify/suite"
)

type averagePrevalentColorTestSuite struct {
	suite.Suite
	URL            string
	PrevalentColor models.PrevalentColor
	mockImage      image.Image
}

func (s *averagePrevalentColorTestSuite) SetupTest() {
	s.PrevalentColor = averageprevalent.NewAveragePrevalentColor("https://www.google.com", "-", "-", "-", false)
	s.mockImage = image.NewRGBA(image.Rect(0, 0, 10, 10))
}

func TestAveragePrevalentColorTestSuite(t *testing.T) {
	suite.Run(t, &averagePrevalentColorTestSuite{})
}

func (s *averagePrevalentColorTestSuite) Test_GetUrl() {
	url := s.PrevalentColor.GetUrl()
	s.Equal("https://www.google.com", url)
}

func (s *averagePrevalentColorTestSuite) Test_GetCalculatedPrevalentColors() {
	c1, c2, c3 := s.PrevalentColor.GetCalculatedPrevalentColors()
	s.Equal("-", c1)
	s.Equal("-", c2)
	s.Equal("-", c3)
}

func (s *averagePrevalentColorTestSuite) Test_SortTopColors() {

	m := make(map[string]int)
	m["-"] = 1
	m["NewColor"] = 2

	s.PrevalentColor.SortTopColors(m, "NewColor")
}

func (s *averagePrevalentColorTestSuite) Test_CalculatePrevalentColors() {
	err := s.PrevalentColor.CalculatePrevalentColors(s.mockImage)
	s.Equal(nil, err)
}
