package functions

import (
	"github.com/brianvoe/gofakeit/v6"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GenerateUser() Entity {
	return User{
		Name:  gofakeit.Name(),
		Email: gofakeit.Email(),
		Role:  gofakeit.RandomString([]string{"editor", "viewer", "administrator"}),
	}
}

func GenerateCompany() Entity {
	return Company{
		ID:            primitive.NewObjectID(),
		Name:          gofakeit.Company(),
		Email:         gofakeit.Email(),
		Phone:         gofakeit.Phone(),
		Street:        gofakeit.Street(),
		City:          gofakeit.City(),
		State:         gofakeit.State(),
		Zip:           gofakeit.Zip(),
		ShopifyKey:    gofakeit.UUID(),
		ShopifySecret: gofakeit.UUID(),
	}
}

func GenerateReturn(company interface{}) func() Entity {
	return func() Entity {
		return Return{
			CompanyID:        company.(Company).ID,
			OrderID:          gofakeit.IntRange(10000, 99999),
			OrderDate:        gofakeit.Date(),
			ItemName:         gofakeit.BeerName(),
			ItemSKU:          gofakeit.UUID(),
			Quantity:         gofakeit.IntRange(4, 10),
			ReturnedQuantity: gofakeit.IntRange(1, 4),
			ReturnReason:     gofakeit.Sentence(5),
			Charity:          gofakeit.Company(),
			Status:           gofakeit.RandomString([]string{"complete", "processing", "received", "pending", "approved"}),
		}
	}
}
