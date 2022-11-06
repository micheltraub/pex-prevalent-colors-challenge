package averageprevalent

import (
	"fmt"
	"image"
	"log"
	"math/rand"
	"pex-prevalent-colors-challenge/pkg/prevalentcolors"
	"time"

	"github.com/muesli/clusters"
	"github.com/muesli/kmeans"
	"github.com/muesli/kmeans/plotter"
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
	rand.Seed(time.Now().UnixNano())
	var d clusters.Observations
	for x := 0; x < 1024; x++ {
		d = append(d, clusters.Coordinates{
			rand.Float64(),
			rand.Float64(),
		})
	}

	km, _ := kmeans.NewWithOptions(0.01, plotter.SimplePlotter{})
	clusters, _ := km.Partition(d, 3)

	for _, c := range clusters {
		fmt.Printf("Centered at x: %.2f y: %.2f\n", c.Center[0], c.Center[1])
		fmt.Printf("Matching data points: %+v\n\n", c.Observations)
	}

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
