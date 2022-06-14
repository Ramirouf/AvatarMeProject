package main

import (
	"log"
	"flag"
	"os"
	"github.com/Ramirouf/AvatarMeProject/pkg/avatar/images"
	"github.com/Ramirouf/AvatarMeProject/pkg/avatar/encoder"
)

type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}
type Service struct {
	//Aislamos las interfaces / desacoplamos las implementaciones
	encoder cryptoEncoder //No le interesa cómo cryptoEncoder codifica el hash
}
type Information struct {
	//Info to encode
	name string
}


func main() {
    var (
	name = flag.String("name", "Ramiro", "Set the name where you want to generate an Identicon for")
    )
    flag.Parse()

    if *name == "" {
	flag.Usage()
	os.Exit(0)
    }

	// Setting the hash attribute of the identicon struct
	encoder := encoder.MD5Encoder{}
	hash, _ :=  s.encoder.EncodeInformation(*name)
    identicon := images.SetHash(hash)
	
    // Pass in the identicon, call the methods which you want to transform
    identicon = images.Pipe(identicon, images.PickColor, images.BuildGrid,
		 images.FilterOddSquares, images.BuildPixelMap)

    // we can use the identicon to insert to our drawRectangle function
    if err := images.DrawRectangle(identicon); err != nil {
	log.Fatalln(err)
    }
}