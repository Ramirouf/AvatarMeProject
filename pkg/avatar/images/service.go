package images

import (
	"image"
	"image/color"
	"github.com/llgcode/draw2d/draw2dimg"
    //Delete the following:
    //"crypto/md5"
)

type Identicon struct {
	hash	[]byte //16 byte array containing the hash returned by the encoder
	color   [3]byte //3 byte array containing Red, Green and Blue
	grid       []byte // New property to hold the grid
	gridPoints []GridPoint // Filtered points in the grid
	pixelMap   []DrawingPoint // pixelMap for drawing

}

// setHash is like a setter... for the hash property
func setHash(input []byte) Identicon {
    return Identicon{
        hash: input,
    }
}

/*
Now we need to define a color for our Identicon, we want the identicon 
to always be the same when we generate it again with the same input. 
The trick is to take the first 3 bytes of the hash, this way we get 
the same color over and over again. The first value will be Red, 
the second value Green and the third value will be Blue.
This way we generate a valid RGB value and we can use it together with the color package in go.
*/
func pickColor(identicon Identicon) Identicon {
    // first we make a byte array with length 3
    rgb := [3]byte{}
    // next we copy the first 3 values from the hash to the rgb array
    copy(rgb[:], identicon.hash[:3])
    // we then assign it to the color value
    identicon.color = rgb
    // return the modified identicon
    return identicon
}

//Now we have an Identicon which holds a color and the hash

func buildGrid(identicon Identicon) Identicon {
    // Create empty grid
    grid := []byte{}
    // Loop over the hash from the identicon
    // Increment with 3 (Chunk the array in 3 parts)
    // this ensures we wont get array out of bounds error and will retrieve exactly 5 chunks of 3
    for i := 0; i < len(identicon.hash) && i+3 <= len(identicon.hash)-1; i += 3 {
        // Create a placeholder for the chunk
        chunk := make([]byte, 5)
        // Copy the items from the old array to the new array
        copy(chunk, identicon.hash[i:i+3])
        chunk[3] = chunk[1] // mirror the second value in the chunk
        chunk[4] = chunk[0] // mirror the first value in the chunk
        grid = append(grid, chunk...) // append the chunk to the grid
    }
    identicon.grid = grid // set the grid property on the identicon
    return identicon // finally return the modified identicon
}

type GridPoint struct {
	value byte
	index int
}

func filterOddSquares(identicon Identicon) Identicon {
    grid := []GridPoint{} // create a placeholder to hold the values of the loop
    for i, code := range identicon.grid { // loop over the grid
	if code%2 == 0 { // check if the value is odd or not
            // create a new Gridpoint where we save the value and the index in the grid
	    point := GridPoint{
		value: code,
		index: i,
	    }
                // append the item to the new grid
	    grid = append(grid, point)
	}
    }
    // set the property
    identicon.gridPoints = grid
    return identicon // return the modified identicon
}

type Point struct {
    x, y int
}

type DrawingPoint struct {
    topLeft     Point
    bottomRight Point
}

func buildPixelMap(identicon Identicon) Identicon {
    drawingPoints := []DrawingPoint{} // define placeholder for drawingpoints

    // Closure, this function returns a Drawingpoint
    pixelFunc := func(p GridPoint) DrawingPoint {
        // This is the formula, we use the index from the gridpoint to calculate the horizontal dimension
        horizontal := (p.index % 5) * 50
        // This is the formula, we use the index from the gridpoint to calculate the vertical dimension
        vertical := (p.index / 5) * 50
        // this is the topleft point with x and the y
        topLeft := Point{horizontal, vertical}
        // the bottom right point is just the topleft point +50 because 1 block in the grid is 50x50
        bottomRight := Point{horizontal + 50, vertical + 50}

        return DrawingPoint{ // We then return the drawingpoint
	    topLeft,
	    bottomRight,
        }
    }

    for _, gridPoint := range identicon.gridPoints {
        // for every gridPoint we calculate the drawingpoints and we add them to the array
        drawingPoints = append(drawingPoints, pixelFunc(gridPoint))
    }
    identicon.pixelMap = drawingPoints // set the drawingpoint value on the identicon
    return identicon // return the modified identicon
}

func rect(img *image.RGBA, col color.Color, x1, y1, x2, y2 float64) {
    gc := draw2dimg.NewGraphicContext(img) // Prepare new image context
    gc.SetFillColor(col) // set the color
    gc.MoveTo(x1, y1) // move to the topleft in the image
    // Draw the lines for the dimensions
    gc.LineTo(x1, y1)
    gc.LineTo(x1, y2)
    gc.MoveTo(x2, y1) // move to the right in the image
    // Draw the lines for the dimensions
    gc.LineTo(x2, y1)
    gc.LineTo(x2, y2)
    // Set the linewidth to zero
    gc.SetLineWidth(0)
    // Fill the stroke so the rectangle will be filled
    gc.FillStroke()
}

func drawRectangle(identicon Identicon) error {
    // We create our default image containing a 250x250 rectangle
    var img = image.NewRGBA(image.Rect(0, 0, 250, 250))
    // We retrieve the color from the color property on the identicon
    col := color.RGBA{identicon.color[0], identicon.color[1], identicon.color[2], 255}

    // Loop over the pixelmap and call the rect function with the img, color and the dimensions
    for _, pixel := range identicon.pixelMap {
	rect(
            img,
            col,
            float64(pixel.topLeft.x),
            float64(pixel.topLeft.y),
            float64(pixel.bottomRight.x),
            float64(pixel.bottomRight.y),
        )
    }
    // Finally save the image to disk
    return draw2dimg.SaveToPngFile("avatar.png", img)
}

type Apply func(Identicon) Identicon

func pipe(identicon Identicon, funcs ...Apply) Identicon {
    for _, applyer := range funcs {
	identicon = applyer(identicon)
    }
    return identicon
}

type buildImage interface {
    pipe(Identicon) Identicon
}
type saveBuiltImage interface {
    SaveBuiltImage(Identicon) error
}
type generateAndSaveImageStruct struct {
    buildImage buildImage
    saveBuiltImage saveBuiltImage

}

/*
Structure that has a method GenerateAndSaveImage
*/
/*
func (g *GenerateAndSaveImageStruct) GenerateAndSaveImage(identicon Identicon) error {
    identicon = g.buildImage.Pipe(identicon)
    return g.saveBuiltImage.SaveBuiltImage(identicon)
}
*/
//Necesito una interfaz que genere y guarde la imagen

// GenerateAndSaveImage should be a method for an struct
// GenerateAndSaveImage receives a hash, saves the image and returns an error
func (i *Identicon) GenerateAndSaveImageIdenticon(encodedInformation []byte) error {
    identicon := setHash(encodedInformation)
    identicon = pipe(identicon, pickColor, buildGrid, filterOddSquares, buildPixelMap)
    return drawRectangle(identicon)
}

/*
From the main service.go file we execute the function 
GenerateAndSaveImage... Sending the hash, and it returns 
an error, which we should handle.
*/




