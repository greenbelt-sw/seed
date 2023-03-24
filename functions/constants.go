package functions

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

const UserCount = 20
const CompanyCount = 100
const ReturnCount = 100

type Entity interface {
}

type User struct {
	Name  string `bson:"name"`
	Email string `bson:"email"`
	Role  string `bson:"role"`
}
type Company struct {
	ID            primitive.ObjectID `bson:"_id,omitempty"`
	Name          string             `bson:"name"`
	Email         string             `bson:"email"`
	Phone         string             `bson:"phone"`
	Street        string             `bson:"street"`
	City          string             `bson:"city"`
	State         string             `bson:"state"`
	Zip           string             `bson:"zip"`
	ShopifyKey    string             `bson:"shopify_key"`
	ShopifySecret string             `bson:"shopify_secret"`
}
type Return struct {
	CompanyID        primitive.ObjectID `bson:"company_id"`
	OrderID          int                `bson:"order_id"`
	OrderDate        time.Time          `bson:"order_date"`
	ItemName         string             `bson:"item_name"`
	ItemSKU          string             `bson:"item_sku"`
	Quantity         int                `bson:"quantity"`
	ReturnedQuantity int                `bson:"returned_quantity"`
	ReturnReason     string             `bson:"return_reason"`
	Charity          string             `bson:"charity"`
	Status           string             `bson:"status"`
}
