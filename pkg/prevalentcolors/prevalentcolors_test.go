package prevalentcolors_test

import (
	"fmt"
	"image"
	"pex-prevalent-colors-challenge/pkg/prevalentcolors"
	"testing"

	"github.com/stretchr/testify/suite"
)

//TODO: implement testing

type prevalentColorsSuite struct {
	suite.Suite
	PrevalentColor prevalentcolors.PrevalentColor
	csvCh          chan []string
}

type mockPrevalentColor struct {
	URL    string
	Color1 string
	Color2 string
	Color3 string
	image  image.Image
}

func (m *mockPrevalentColor) FetchImage() (image.Image, error) {
	return image.NewRGBA(image.Rect(0, 0, 10, 10)), nil
}

func (m *mockPrevalentColor) CalculatePrevalentColors(img image.Image) error {
	return nil
}
func (m *mockPrevalentColor) GetUrl() string {
	return m.URL
}
func (m *mockPrevalentColor) GetCalculatedPrevalentColors() (string, string, string) {
	return m.Color1, m.Color2, m.Color3
}
func (m *mockPrevalentColor) SortTopColors(ma map[string]int, s string) {
}

type mockPrevalentColorFetchError struct {
	URL    string
	Color1 string
	Color2 string
	Color3 string
	image  image.Image
}

func (m *mockPrevalentColorFetchError) FetchImage() (image.Image, error) {
	return nil, fmt.Errorf("Error fetching image!")
}

func (m *mockPrevalentColorFetchError) CalculatePrevalentColors(img image.Image) error {
	return nil
}
func (m *mockPrevalentColorFetchError) GetUrl() string {
	return m.URL
}
func (m *mockPrevalentColorFetchError) GetCalculatedPrevalentColors() (string, string, string) {
	return m.Color1, m.Color2, m.Color3
}
func (m *mockPrevalentColorFetchError) SortTopColors(ma map[string]int, s string) {
}

type mockPrevalentColorCalculatingError struct {
	URL    string
	Color1 string
	Color2 string
	Color3 string
	image  image.Image
}

func (m *mockPrevalentColorCalculatingError) FetchImage() (image.Image, error) {
	return image.NewRGBA(image.Rect(0, 0, 10, 10)), nil
}

func (m *mockPrevalentColorCalculatingError) CalculatePrevalentColors(img image.Image) error {
	return fmt.Errorf("Error calculating prevalent colors!")
}
func (m *mockPrevalentColorCalculatingError) GetUrl() string {
	return m.URL
}
func (m *mockPrevalentColorCalculatingError) GetCalculatedPrevalentColors() (string, string, string) {
	return m.Color1, m.Color2, m.Color3
}
func (m *mockPrevalentColorCalculatingError) SortTopColors(ma map[string]int, s string) {
}

func (s *prevalentColorsSuite) SetupTest() {
	s.PrevalentColor = &mockPrevalentColor{
		Color1: "red",
		Color2: "blue",
		Color3: "green",
		URL:    "test.com",
		image:  image.NewRGBA(image.Rect(0, 0, 10, 10)),
	}
	s.csvCh = make(chan []string, 1)
}

func TestAccuratePrevalentColorTestSuite(t *testing.T) {
	suite.Run(t, new(prevalentColorsSuite))
}

func (s *prevalentColorsSuite) Test_ProcessPrevalentColors() {
	s.csvCh = make(chan []string, 1)

	prevalentcolors.ProcessPrevalentColors(s.PrevalentColor, s.csvCh)
	close(s.csvCh)
	for c := range s.csvCh {
		csvLine := c
		s.Equal(csvLine, []string{"test.com", "red", "blue", "green", ""})
		return
	}
}

func (s *prevalentColorsSuite) Test_ProcessPrevalentColors_ErrorFetching() {
	s.csvCh = make(chan []string, 1)
	s.PrevalentColor = &mockPrevalentColorFetchError{
		Color1: "red",
		Color2: "blue",
		Color3: "green",
		URL:    "test.com",
		image:  image.NewRGBA(image.Rect(0, 0, 10, 10)),
	}

	prevalentcolors.ProcessPrevalentColors(s.PrevalentColor, s.csvCh)
	close(s.csvCh)
	for c := range s.csvCh {
		csvLine := c
		s.Equal(csvLine, []string{"test.com", "-", "-", "-", "Error fetching image!"})

		return
	}
}

func (s *prevalentColorsSuite) Test_ProcessPrevalentColors_ErrorCalculating() {
	s.csvCh = make(chan []string, 1)
	s.PrevalentColor = &mockPrevalentColorCalculatingError{
		Color1: "red",
		Color2: "blue",
		Color3: "green",
		URL:    "test.com",
		image:  image.NewRGBA(image.Rect(0, 0, 10, 10)),
	}

	prevalentcolors.ProcessPrevalentColors(s.PrevalentColor, s.csvCh)
	close(s.csvCh)
	for c := range s.csvCh {
		csvLine := c
		s.Equal(csvLine, []string{"test.com", "-", "-", "-", "Error calculating prevalent colors!"})

		return
	}
}
