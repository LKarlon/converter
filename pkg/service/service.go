package service

type Conversion interface {
	Convert(fileIn []byte) (fileOut []byte, err error)
}

type Service struct {
	Conversion
}

func NewService() *Service {
	return &Service{}
}
 