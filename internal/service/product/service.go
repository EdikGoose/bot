package product

type Service struct{}

func NewService() *Service {
	return &Service{}
}

func (s *Service) List() []Product {
	return allProducts
}

func (s *Service) Get(idx int) (*Product, error) {
	// TODO: сделать проверку и в случае невалидного индекса возвращать ошибку (обработать её тоже не забыть)
	return &allProducts[idx], nil
}
