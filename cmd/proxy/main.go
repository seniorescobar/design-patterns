package main

import (
	"fmt"

	"github.com/seniorescobar/design-patterns/cmd/proxy/image"
)

func main() {
	bm1 := image.NewBitmap("neymar.jpg")
	drawImage(bm1)

	bm2 := image.NewLazyBitmap("coutinho.jpg")
	drawImage(bm2)
	drawImage(bm2)
}

func drawImage(img image.Image) {
	fmt.Println("About to draw image")
	img.Draw()
	fmt.Println("Done drawing image")
	fmt.Println("---")
}
