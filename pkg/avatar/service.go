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
func hash(s string) uint32 {
        h := fnv.New32a()
        h.Write([]byte(s))
        return h.Sum32()
}
func main() {
	fmt.Println(hash("test"))
	fmt.Println(hash("tesl"))

}
*/
/*
func DefaultAvatarGeneration (name string) string {
	return fmt.Sprintf("http://www.gravatar.com/avatar/%x?s=%d",
		hash(name), defaultWidth)
}
*/
func DefaultAvatarGeneration () *GeneratorOne {
	return &GeneratorOne{
		encoder: &encoder.MD5Encoder{},
		generator: images.NewDrawer(images.NewDefaultColorEngine(), defaultLength, defaultWidth),
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
	fmt.Fprintln("Avatar generated and saved, related to %s", information.name)")
	return nil
	
}