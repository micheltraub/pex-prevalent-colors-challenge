package accurateprevalent

import (
	"fmt"
	"image"
	"log"
	"net/http"
	"pex-prevalent-colors-challenge/internal/app/helpers"
	"pex-prevalent-colors-challenge/pkg/prevalentcolors"
	"sort"

	"github.com/nfnt/resize"
)

// - accuratePrevalentColor: Get the exact 3 pixels that appears the most
type AccuratePrevalentColor struct {
	URL       string
	Color1    string
	Color2    string
	Color3    string
	Downscale bool
}

func NewAccuratePrevalentColor(url string, color1 string, color2 string, color3 string, downscale bool) prevalentcolors.PrevalentColor {
	return &AccuratePrevalentColor{
		URL:       url,
		Color1:    color1,
		Color2:    color2,
		Color3:    color3,
		Downscale: downscale,
	}
}

func (accurate *AccuratePrevalentColor) FetchImage() (image.Image, error) {
	if len(accurate.URL) <= 0 {
		return nil, fmt.Errorf("URL is empty")
	}
	log.Println("📥 Fetching image from: ", accurate.URL)
	//Get the response bytes from the url
	response, err := http.Get(accurate.URL)
	if err != nil {
		return nil, err
	}
	img, filename, err := image.Decode(response.Body)
	if err != nil {
		log.Println("🚧🚨 Error decoding image: " + err.Error())
		return nil, err
	}
	log.Println("🖼️ Image type: " + filename)
	defer response.Body.Close()
	return img, nil
}

func (accurate *AccuratePrevalentColor) CalculatePrevalentColors(img image.Image) error {
	log.Println("🧑‍💻 Calculating prevalent color")
	//Reducing the size of the image with interpolation makes the pixels counting faster,
	//consuming less memory but less accurate. Diving by 5 the width andk eeping the ratio
	if accurate.Downscale {
		img = resize.Resize(uint(img.Bounds().Max.X)/5, 0, img, resize.Lanczos3)
	}
	//count the pixels by hex code in map
	m := make(map[string]int)
	m["-"] = -1
	for x := 0; x < img.Bounds().Max.X; x++ {
		for y := 0; y < img.Bounds().Max.Y; y++ {
			hexPixel := helpers.RgbaToHex(img.At(x, y).RGBA())
			m[hexPixel] = m[hexPixel] + 1
			accurate.SortTopColors(m, hexPixel)
		}
	}
	log.Println("🌈 Prevalent Colors: " + accurate.Color1 + ", " + accurate.Color2 + ", " + accurate.Color3)
	return nil
}

// Get top 3 from map comparing to new pixel
func (accurate *AccuratePrevalentColor) SortTopColors(m map[string]int, hexPixel string) {
	//make a samller map to sort with the top3 colors plus the one I want to compare with
	top := make(map[string]int)
	top[accurate.Color1] = m[accurate.Color1]
	top[accurate.Color2] = m[accurate.Color2]
	top[accurate.Color3] = m[accurate.Color3]
	//if the hex color is not in the top map, add it
	if _, ok := top[hexPixel]; !ok {
		top[hexPixel] = m[hexPixel]
	}
	//sort the top map

	keys := make([]string, 0, len(top))
	for key := range top {
		keys = append(keys, key)
	}
	sort.SliceStable(keys, func(i, j int) bool {
		return top[keys[i]] > top[keys[j]]
	})

	//set the top 3 colors
	if len(keys) > 0 {
		accurate.Color1 = keys[0]
	}
	if len(keys) > 1 {
		accurate.Color2 = keys[1]
	}
	if len(keys) > 2 {
		accurate.Color3 = keys[2]
	}
}

func (accurate *AccuratePrevalentColor) GetCalculatedPrevalentColors() (string, string, string) {
	return accurate.Color1, accurate.Color2, accurate.Color3
}
func (accurate *AccuratePrevalentColor) GetUrl() string {
	return accurate.URL
}
