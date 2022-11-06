package accurateprevalent_test

import (
	"image"
	"image/color"
	"pex-prevalent-colors-challenge/internal/app/accurateprevalent"
	"pex-prevalent-colors-challenge/pkg/prevalentcolors"
	"testing"

	"github.com/stretchr/testify/suite"
)

type accuratePrevalentColorTestSuite struct {
	suite.Suite
	URL            string
	PrevalentColor prevalentcolors.PrevalentColor
	mockImage      image.Image
}

type mockImage struct {
	img image.Image
}

func (m *mockImage) ColorModel() { return }
func (m *mockImage) Bounds() image.Rectangle {
	return image.Rect(0, 0, 1, 1)
}
func (m *mockImage) At(x, y int) color.Color {
	return color.Black
}

func (s *accuratePrevalentColorTestSuite) SetupTest() {
	s.PrevalentColor = accurateprevalent.NewAccuratePrevalentColor("https://www.google.com", "-", "-", "-", false)
	s.mockImage = image.NewRGBA(image.Rect(0, 0, 10, 10))
}

func TestAccuratePrevalentColorTestSuite(t *testing.T) {
	suite.Run(t, &accuratePrevalentColorTestSuite{})
}

func (s *accuratePrevalentColorTestSuite) Test_GetUrl() {
	url := s.PrevalentColor.GetUrl()
	s.Equal("https://www.google.com", url)
}

func (s *accuratePrevalentColorTestSuite) Test_GetCalculatedPrevalentColors() {
	c1, c2, c3 := s.PrevalentColor.GetCalculatedPrevalentColors()
	s.Equal("-", c1)
	s.Equal("-", c2)
	s.Equal("-", c3)
}

func (s *accuratePrevalentColorTestSuite) Test_SortTopColors2Color() {

	m := make(map[string]int)
	m["-"] = 1
	m["NewColor"] = 2

	s.PrevalentColor.SortTopColors(m, "NewColor")
	c1, c2, c3 := s.PrevalentColor.GetCalculatedPrevalentColors()
	s.Equal("NewColor", c1)
	s.Equal("-", c2)
	s.Equal("-", c3)
}

func (s *accuratePrevalentColorTestSuite) Test_SortTopColors3Color() {
	s.PrevalentColor = accurateprevalent.NewAccuratePrevalentColor("https://www.google.com", "Color1", "-", "-", false)
	m := make(map[string]int)
	m["-"] = 1
	m["Color1"] = 3
	m["NewColor"] = 2

	s.PrevalentColor.SortTopColors(m, "NewColor")
	c1, c2, c3 := s.PrevalentColor.GetCalculatedPrevalentColors()
	s.Equal("Color1", c1)
	s.Equal("NewColor", c2)
	s.Equal("-", c3)
}

func (s *accuratePrevalentColorTestSuite) Test_CalculatePrevalentColors() {
	err := s.PrevalentColor.CalculatePrevalentColors(s.mockImage)
	if err != nil {
		s.Error(err)
	}
	c1, c2, c3 := s.PrevalentColor.GetCalculatedPrevalentColors()
	s.Equal("#000000", c1)
	s.Equal("-", c2)
	s.Equal("-", c3)
}
