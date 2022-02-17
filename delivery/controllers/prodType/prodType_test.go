package prodType

import "be/models"

type mockProdTypeLib struct{}

func (m *mockProdTypeLib) Create(proType models.ProductType) (models.ProductType, error) {
	return models.ProductType{}
}