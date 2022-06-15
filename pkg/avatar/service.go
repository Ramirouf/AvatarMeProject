package main

import (
	"fmt"
	"github.com/Ramirouf/AvatarMeProject/pkg/avatar/images"
	"github.com/Ramirouf/AvatarMeProject/pkg/avatar/encoder"
	//"flag"
)
/*
// cryptoEncoder encodes information.
type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}


// imageGenerator makes images.
type imageGenerator interface {
	GenerateAndSaveImage(encodedInformation []byte) error
}


// Service contains functionalities to generate avatars... it's a generator.
type Service struct {
	//Services isolates the implementation of the interfaces
	encoder cryptoEncoder //Doesn't care how the encoder hashes the information
	//generator imageGenerator
}


// Information 
type Information struct {
	//Info to encode
	name string
}

func AvatarGenerator() *Service {
	return &Service{
		encoder: &encoder.MD5Encoder{},
		generator: images.GenerateAndSaveImage,
	}
}
//This method "GenerateAndSaveAvatar" is used to generate the avatar
//It has a receiver of type Service named "s"
//Now, "s" is a pointer to a Service, which will have this following method
// GenerateAndSaveAvatar is a constructor that generates the avatar

func (s *Service) GenerateAndSaveAvatar (information Information) error {
	//All logic is here

	//Information encoded... then it's saved in encodedBytes
	encodedBytes, err :=  s.encoder.EncodeInformation(information.name)
	if err != nil {
		return err
	}
	//Using encodedBytes, we generate the avatar
	err = s.generator.BuildAndSaveImage(encodedBytes)
	if err != nil {
		return err
	}
	fmt.Fprintln("Avatar generated and saved, related to %s", information.name)
	return nil
	
}
*/
func main() {
	/*
	var name = flag.String("name", "Ramiro", "Text to hash")
    encoder := encoder.MD5Encoder{}
    hash, _ := encoder.EncodeInformation(*name)

    fmt.Println(hash)
	images.GenerateAndSaveImage(hash)
	*/
	service1 := AvatarGenerator()
	service1.GenerateAndSaveAvatar()
}

type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}

type Service struct {
	//Services isolates the implementation of the interfaces
	encoder cryptoEncoder
	//generator imageGenerator
}

func AvatarGenerator() *Service {
	return &Service{
		encoder: &encoder.MD5Encoder{},
		//generator: images.GenerateAndSaveImage,
	}
}

func (s *Service) GenerateAndSaveAvatar () error {

	hash, _ := s.encoder.EncodeInformation("Ramiro11")
	//Using encodedBytes, we generate the avatar
	err := images.GenerateAndSaveImage(hash)
	if err != nil {
		return err
	}
	fmt.Println("Avatar generated and saved, related to %s", "Ramiro1")
	return nil
}