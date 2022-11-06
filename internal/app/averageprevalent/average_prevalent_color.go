package averageprevalent

import (
	"image"
	"log"
	"pex-prevalent-colors-challenge/pkg/prevalentcolors"
)

//   - averagePrevalentColor: Get the 3 pixels that appears the most using KMeansClustering
//     Reference: https://en.wikipedia.org/wiki/K-means_clustering
type averagePrevalentColor struct {
	URL       string
	Color1    string
	Color2    string
	Color3    string
	Downscale bool
}

func NewAveragePrevalentColor(url string, color1 string, color2 string, color3 string, downscale bool) prevalentcolors.PrevalentColor {
	return &averagePrevalentColor{
		URL:       url,
		Color1:    color1,
		Color2:    color2,
		Color3:    color3,
		Downscale: downscale,
	}
}

// TODO: Implement and check a better way to avoid repeting code
func (accurate *averagePrevalentColor) FetchImage() (image.Image, error) {
	log.Println("📥 Fetching image from: ", accurate.URL)
	log.Println("⚠️ Not implemented yet")
	return nil, nil
}

// TODO: Use KMeansClustering to calculate prevalent colors
func (accurate *averagePrevalentColor) CalculatePrevalentColors(img image.Image) error {
	log.Println("🧑‍💻 Calculating prevalent color")
	log.Println("⚠️ Not implemented yet")
	return nil
}
func (accurate *averagePrevalentColor) GetUrl() string {
	return accurate.URL
}
func (accurate *averagePrevalentColor) GetCalculatedPrevalentColors() (string, string, string) {
	return accurate.Color1, accurate.Color2, accurate.Color3
}
func (accurate *averagePrevalentColor) SortTopColors(m map[string]int, s string) {
	log.Println("⚠️ Not implemented yet")
}
