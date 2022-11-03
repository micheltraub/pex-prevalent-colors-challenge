package helpers_test

import (
	"pex-prevalent-colors-challenge/internal/app/helpers"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type rgbaSuite struct {
	suite.Suite
	ExpectedHexString string
}

func (s *rgbaSuite) SetupTest() {
}

func (s *rgbaSuite) TestRgbaToHex_ReturnsCorrectWhiteString() {
	s.ExpectedHexString = "#ffffff"
	assert.Equal(s.T(), s.ExpectedHexString, helpers.RgbaToHex(65535, 65535, 65535, 65535))
}

func (s *rgbaSuite) TestRgbaToHex_ReturnsCorrectBlackString() {
	s.ExpectedHexString = "#000000"
	assert.Equal(s.T(), s.ExpectedHexString, helpers.RgbaToHex(0, 0, 0, 0))
}

func TestRgbaTestSuite(t *testing.T) {
	suite.Run(t, new(rgbaSuite))
}
