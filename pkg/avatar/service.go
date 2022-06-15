package main

import (
	"fmt"
	"github.com/Ramirouf/AvatarMeProject/pkg/avatar/images"
	"github.com/Ramirouf/AvatarMeProject/pkg/avatar/encoder"
)

func main() {

	info1 := Information{name: "Ramiro111111"}
	service1 := AvatarGenerator()
	service1.GenerateAndSaveAvatar(info1)
}

// cryptoEncoder encodes information.
type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}
// imageGenerator makes images.
type imageGenerator interface {
	GenerateAndSaveImageIdenticon(encodedInformation []byte) error
}
// Service contains functionalities to generate avatars... it's a generator.

type Service struct {
	//Services isolates the implementation of the interfaces
	encoder cryptoEncoder
	generator imageGenerator
}

func AvatarGenerator() *Service {
	return &Service{
		encoder: &encoder.MD5Encoder{},
		//generator: images.GenerateAndSaveImageIdenticon,
		generator: &images.Identicon{},
	}
}
// Information struct, which contains the information to encode.

type Information struct {
	name string
}

//This method "GenerateAndSaveAvatar" is used to generate the avatar
//It has a receiver of type Service named "s"
//Now, "s" is a pointer to a Service, which will have this following method
// GenerateAndSaveAvatar is a constructor that generates the avatar
func (s *Service) GenerateAndSaveAvatar (info Information) error {

	hash, err := s.encoder.EncodeInformation(info.name)
	if err != nil {
		return err
	}
	//Using encodedBytes, we generate the avatar
	err = s.generator.GenerateAndSaveImageIdenticon(hash)
	if err != nil {
		return err
	}
	fmt.Println("Avatar generated and saved, related to", info.name)
	return err
}

