package product

type Service struct {

}

func NewService()*Service{
	return &Service{}
}

func (s *Service)List()[]Product{
	return allProducts
}

func (s *Service)Get(id int)(*Product,error){
	return &allProducts[id],nil
}