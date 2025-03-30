package services

type HashingService struct{}

func NewHashingService() *HashingService {
	return &HashingService{}
}

// TODO: Implement hashing functionality
func (hs *HashingService) Hash(dataToHash string) string {
	return dataToHash
}
