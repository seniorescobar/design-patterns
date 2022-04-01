package image

import "fmt"

type Image interface {
	Draw()
}

type Bitmap struct {
	filename string
}

func NewBitmap(filename string) *Bitmap {
	fmt.Println("Loading bitmap from disk")

	return &Bitmap{
		filename: filename,
	}
}

func (b *Bitmap) Draw() {
	fmt.Println("Drawing bitmap", b.filename)
}

type LazyBitmap struct {
	filename string
	bitmap   *Bitmap
}

func NewLazyBitmap(filename string) *LazyBitmap {
	return &LazyBitmap{
		filename: filename,
	}
}

func (b *LazyBitmap) Draw() {
	if b.bitmap == nil {
		b.bitmap = NewBitmap(b.filename)
	}

	b.bitmap.Draw()
}
