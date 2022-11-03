package accurateprevalent_test

import (
	"io"
	"pex-prevalent-colors-challenge/internal/app/accurateprevalent"
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
	s.PrevalentColor = accurateprevalent.NewAccuratePrevalentColor("http://i.imgur.com/FApqk3D.jpg", "-", "-", "-")
}

func TestAccuratePrevalentColorTestSuite(t *testing.T) {
	suite.Run(t, &accuratePrevalentColorTestSuite{})
}

func (s *accuratePrevalentColorTestSuite) Test_FetchImage() {
	img, _ := s.PrevalentColor.FetchImage()
	println("Bounds: ", img.Bounds().Min.X, img.Bounds().Min.Y, img.Bounds().Max.X, img.Bounds().Max.Y)
	s.Equal(img.Bounds().Min.Y, 5)
}
