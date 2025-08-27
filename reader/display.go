package reader

import (
	"fmt"
	"image"

	"github.com/blacktop/go-termimg"
)

// path is for caching
func DisplayImage(img image.Image) {
	ti := termimg.New(img)
	ti.Scale(termimg.ScaleMode(2))

	err := ti.Print()
	if err != nil {
		panic(err)
	}

	fmt.Println()

}

