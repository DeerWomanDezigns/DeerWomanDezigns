package models

import (
	"time"
)

type Order struct {
	ID                     int           `dynamo:"order_id" json:"id"`
	Type                   int           `dynamo:"order_type" json:"type"`
	SellerEmail            string        `dynamo:"seller_email" json:"sellerEmail"`
	BuyerEmail             string        `dynamo:"buyer_email" json:"buyerEmail"`
	BuyerName              string        `dynamo:"buyer_name" json:"buyerName"`
	BuyerAddressFirstLine  string        `dynamo:"buyer_addr_first" json:"buyerAddressFirstLine"`
	BuyerAddressSecondLine string        `dynamo:"buyer_addr_second" json:"buyerAddressSecondLine"`
	BuyerCity              string        `dynamo:"buyer_city" json:"buyerCity"`
	BuyerState             string        `dynamo:"buyer_state" json:"buyerState"`
	BuyerZip               string        `dynamo:"buyer_zip" json:"buyerZip"`
	Status                 string        `dynamo:"order_status" json:"status"`
	BuyerFormattedAddress  string        `dynamo:"buyer_formatted_address" json:"buyerFormattedAddress"`
	BuyerCountryIso        string        `dynamo:"buyer_country_iso" json:"buyerCountryIso"`
	PaymentMethod          string        `dynamo:"payment_method" json:"paymentMethod"`
	PaymentEmail           string        `dynamo:"payment_email" json:"paymentEmail"`
	SellerMessage          string        `dynamo:"seller_message" json:"sellerMessage"`
	BuyerMessage           string        `dynamo:"buyer_message" json:"buyerMessage"`
	PaymentMessage         string        `dynamo:"payment_message" json:"paymentMessage"`
	IsPaid                 bool          `dynamo:"is_paid" json:"isPaid"`
	IsShipped              bool          `dynamo:"is_Shipped" json:"isShipped"`
	CreateTime             time.Time     `dynamo:"create_time" json:"createTime"`
	UpdateTime             time.Time     `dynamo:"update_time" json:"updateTime"`
	GiftMessage            string        `dynamo:"gift_message" json:"giftMessage"`
	GrandTotal             Cost          ``
	Subtotal               Cost          ``
	TotalPrice             Cost          ``
	TotalShippingCost      Cost          ``
	TotalTaxCost           Cost          ``
	TotalVat               Cost          ``
	Discount               Cost          ``
	GiftWrapCost           Cost          ``
	Shipments              []Shipment    ``
	Transactions           []Transaction ``
}

type Cost struct {
	AmountCents  int    ``
	Divisor      int    ``
	CurrencyCode string ``
}

type Shipment struct {
	ShipmentID               int       ``
	ShipmentNotificationTime time.Time ``
	ShipmentCarrierName      string    ``
	ShipmentTrackingCode     string    ``
}

type Transaction struct {
	ID                int         ``
	Title             string      ``
	Description       string      ``
	SellerID          int         ``
	BuyerID           int         ``
	CreateTime        time.Time   ``
	PaidTime          time.Time   ``
	ShippedTime       time.Time   ``
	Quantity          int         ``
	ListingImageID    int         ``
	ReceiptID         int         ``
	IsDigital         bool        ``
	FileData          string      ``
	ListingID         int         ``
	Type              string      ``
	ProductID         int         ``
	Price             Cost        ``
	ShippingCost      Cost        ``
	Variations        []Variation ``
	ShippingProfileID int         ``
	MinProcessingDays int         ``
	MaxProcessingDays int         ``
	ShippingMethod    string      ``
	ShippingUpgrade   string      ``
	ExpectedShipDate  string      ``
	ActualShipDate    string      ``
}

type Variation struct {
	ProperyID      int    ``
	ValueID        int    ``
	FormattedName  string ``
	FormattedValue string ``
}

type User struct {
	ID        string `dynamo:"user_id" json:"id"`
	FirstName string `dynamo:"first_name" json:"firstName"`
	LastName  string `dynamo:"last_name" json:"lastName"`
	Address   string `dynamo:"address" json:"address"`
	Email     string `dynamo:"email" json:"email"`
}
