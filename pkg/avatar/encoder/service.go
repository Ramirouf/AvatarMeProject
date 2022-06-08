package encoder

import "crypto/md5"

type MD5Encoder struct {} //Type declaration for MD5Encoder
/*
Other way:

func (e *MD5Encoder) Encode(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return fmt.Sprintf("%x", h.Sum(nil))
}
*/
func (e *MD5Encoder) EncodeInformation(strInformation string) (encodedInformation []byte, err error) {
	hash := md5.Sum([]byte(strInformation))
	return hash[:], nil
}