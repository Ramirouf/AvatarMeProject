package avatar

import (
	"fmt"
)

// cryptoEncoder encodes information.
type cryptoEncoder interface {
	EncodeInformation(strInformation string) (encodedInformation []byte, err error)
}

// imageGenerator makes images.
type imageGenerator interface {
	BuildAndSaveImage(encodedInformation []byte) error
}


// Service contains functionalities to generate avatars... it's a generator.
type Service struct {
	//Services isolates the implementation of the interfaces
	encoder cryptoEncoder //Doesn't matter how the encoder hashes the information
	generator imageGenerator
}


// Information 
type Information struct {
	//Info to encode
	name string
}
func AvatarGenerator() *Service {
	return &Service{
		encoder: &encoder.MD5Encoder{},
		generator: images.Pipe{},
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