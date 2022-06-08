package images

import (
	"image"
	"image/color"
	"image/png"
	"os"
)

//This file should use BuildAndSaveImage

const (
	informationLength = 16
	filename = "avatarImagen"
)

type ColorTransformer interface {
	ByteToColor (colorByte byte)(color.Color, error)
}

type Drawer struct {
	//These two are private structures
	colorTransformer ColorTransformer
	imageWidth, imageLength, stride int
}
// NewDrawer is a constructor for Drawer
func NewDrawer (colorTransformer ColorTransformer, imageLength, imageWidth int) *Drawer {
	return &Drawer{
		colorTransformer: colorTransformer,
		imageLength: imageLength,
		imageWidth: imageWidth,
		stride: imageWidth / informationLength,
	}
}

func (d *Drawer) BuildAndSaveImage (encodedInformation []byte) error {
	//This is the logic to build the image
	//We need to create a new image
	if len(encodedInformation) != informationLength {
		return fmt.Errorf("Invalid encoded information length")
	}
	//encodedColor, err := d.colorTransformer.ByteToColor(encodedInformation[0])
	encodedColor, err := d.bytesArrayToColorArray (encodedInformation)
	if err != nil {
		return err
	}
	avatarImage := d.createImage(encodedColor)
	f, err := os.Create(filename)

	if err != nil {
		return err
	}
	defer f.Close()
	err = png.Encode(f, avatarImage)
	//We need to save the image
	return err
}

