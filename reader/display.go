package reader

import (
	"fmt"
	"image"
	"os"

	"github.com/blacktop/go-termimg"
	"github.com/nfnt/resize"
	"golang.org/x/term"
)

var imageInterpolationMethod = resize.MitchellNetravali

func DisplayImage(img image.Image) {
	tf := termimg.QueryTerminalFeatures()

	wWPixels := tf.WindowCols * tf.FontWidth
	wHPixels := tf.WindowRows * tf.FontHeight
	windowAspectRatio := float32(wWPixels) / float32(wHPixels)

	imageAspectRatio := float32(img.Bounds().Dx()) / float32(img.Bounds().Dy())

	if imageAspectRatio > windowAspectRatio {
		img = resize.Resize(uint(tf.WindowCols * tf.FontWidth), 0, img, imageInterpolationMethod)
	} else {
		img = resize.Resize(0, uint(tf.WindowRows * tf.FontHeight), img, imageInterpolationMethod)
	}


	ti := termimg.New(img)

	err := ti.Print()
	if err != nil {
		panic(err)
	}


}


func Read(images []image.Image) {
	oldState, err := term.MakeRaw(int(os.Stdin.Fd()))
	if err != nil {
		panic(err)
	}
	defer term.Restore(int(os.Stdin.Fd()), oldState)


	currentIndex := 0



	fmt.Print(TerminalEscapeCodes.HideCursor)
	defer fmt.Print(TerminalEscapeCodes.ShowCursor)

	buf := make([]byte, 1)

	loop:
	for {
		fmt.Print(TerminalEscapeCodes.ClearScreen, TerminalEscapeCodes.MoveCursorTopLeft)

		termimg.ClearResizeCache()
		DisplayImage(images[currentIndex])


		_, err := os.Stdin.Read(buf)
		if err != nil {
			panic(err)
		}

		switch buf[0] {
		case 'j':
			currentIndex++
		case 'k':
			currentIndex--
		case 'q':
			break loop
		default:
			fmt.Printf("pressed %s\n", string(buf[0]))
			panic("unrecognized key")
		}

	}


	fmt.Print(TerminalEscapeCodes.ClearScreen, TerminalEscapeCodes.MoveCursorTopLeft)


}



