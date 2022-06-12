package encoder

import "crypto/md5"

type MD5Encoder struct {} //Type declaration for MD5Encoder

func (e *MD5Encoder) EncodeInformation(strInformation string) (encodedInformation []byte, err error) {
	hash := md5.Sum([]byte(strInformation))
	return hash[:], nil
}

