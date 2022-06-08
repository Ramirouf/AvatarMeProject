package avatar

import (
	"fmt"
	"hash/fnv"
)
const (
	defaultWidth = 200
	defaultLength = 200
	strider = 10
)
/*
func DefaultAvatarGeneration (name string) string {
	return fmt.Sprintf("http://www.gravatar.com/avatar/%x?s=%d",
		hash(name), defaultWidth)
}
*/

// DefaultAvatarGenerator provides a Generator that generates an image hashing the Information
// with MD5 and saving the image with a ColorEngine
func DefaultAvatarGenerator () *GeneratorOne {
	return &GeneratorOne{
		encoder: &encoder.MD5Encoder{},
		generator: images.NewDrawer(images.NewDefaultColorEngine(), defaultLength, defaultWidth),
	}
}
func CustomGenerator(cryptoEncoder cryptoEncoder, generator imageGenerator) *GeneratorOne {
	return &GeneratorOne{
		encoder: cryptoEncoder,
		generator: generator,
	}
}
// cryptoEncoder encodes information
type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}
// imageGenerator makes images
type imageGenerator interface {
	BuildAndSaveImage (encodedInformation []byte) error
}
// Service contains functionalities to generate avatars... it's a generator
type Service struct {
	//Aislamos las interfaces / desacoplamos las implementaciones
	encoder cryptoEncoder //No le interesa cómo cryptoEncoder codifica el hash
	generator imageGenerator
}
// Information 
type Information struct {
	//Info to encode
	Name string
}

FUNC (S *GeneratorOne) GenerateAndSaveAvatar (information Information) error {
	//All logic is here
	encodedBytes, err :=  s.encoder.EncodeInformation(information.name)
	if err != nil {
		return err
	}
	err = s.generator.BuildAndSaveImage(encodedBytes)
	if err != nil {
		return err
	}
	fmt.Fprintln("Avatar generated and saved, related to %s", information.name)
	return nil
	
}