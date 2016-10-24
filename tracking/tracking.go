package tracking

import (
	"fmt"
	"time"
)

// APIRequest ...
type APIRequest struct {
	UPSSecurity  UPSSecurity  `json:"UPSSecurity"`
	TrackRequest TrackRequest `json:"TrackRequest"`
}

// UPSSecurity ...
type UPSSecurity struct {
	UsernameToken      UsernameToken      `json:"UsernameToken"`
	ServiceAccessToken ServiceAccessToken `json:"ServiceAccessToken"`
}

// ServiceAccessToken ...
type ServiceAccessToken struct {
	AccessLicenseNumber string `json:"AccessLicenseNumber"`
}

// UsernameToken ...
type UsernameToken struct {
	Username string `json:"Username"`
	Password string `json:"Password"`
}

// TrackRequest ...
type TrackRequest struct {
	Request       Request `json:"Request"`
	InquiryNumber string  `json:"InquiryNumber"`
}

// Request ...
type Request struct {
	RequestOption        int                  `json:"RequestOption"`
	TransactionReference TransactionReference `json:"TransactionReference"`
}

// TransactionReference ...
type TransactionReference struct {
	CustomerContext string `json:"CustomerContext"`
}

// TrackResponse ...
type TrackResponse struct {
	Reponse  Reponse  `json:"Response"`
	Shipment Shipment `json:"Shipment"`
}

// Response ...
type Response struct {
	ResponseStatus       ResponseStatus       `json:"ResponseStatus"`
	TransactionReference TransactionReference `json:"TransactionReference"`
}

// ResponseStatus ...
type ResponseStatus struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

// Shipment ...
type Shipment struct {
	InquiryNumber InquiryNumber `json:"InquiryNumber"`
	ShipmentType  ShipmentType  `json:"ShipmentType"`
	ShipperNumber string        `json:"ShipperNumber"`
	Service       Service       `json:"Service"`
	PickupDate    string        `json:"PickupDate"`
	Package       Package       `json:"Package"`
}

// GetPickupDate
func (s Shipment) GetPickupDate() time.Time {
	return time.parse("20060102", s.PickupDate)
}

// InquiryNumber ...
type InquiryNumber struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
	Value       string `json:"Value"`
}

// ShipmentType ...
type ShipmentType struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

// Service ...
type Service struct {
	Code        string `json:"Code"`
	Description string `json:"Description"`
}

// Package ...
type Package struct {
	TrackingNumber   string            `json:"TrackingNumber"`
	Activity         Activity          `json:"Activity"`
	PackageWeight    PackageWeight     `json:"PackageWeight"`
	ReferenceNumbers []ReferenceNumber `json:"ReferenceNumber"`
}

// Activity ...
type Activity struct {
	ActivityLocation ActivityLocation `json:"ActivityLocation"`
	Status           Status           `json:"Status"`
	Date             string           `json:"Date"`
	Time             string           `json:"Time"`
}

// GetTimestamp ...
func (a Activity) GetTimestamp() time.Time {
	return time.Parse("20060102T150405", fmt.Sprintf("%sT%s", a.Date, a.Time))
}

// ActivityLocation ...
type ActivityLocation struct {
	Address Address `json:"Address"`
}

// Address ...
type Address struct {
	City              string `json:"City"`
	StateProvinceCode string `json:"StateProvinceCode"`
	CountryCode       string `json:"CountryCode"`
}

// Status ...
type Status struct {
	Type        string `json:"Type"`
	Description string `json:"Description"`
	Code        string `json:"Code"`
}

// PackageWeight ...
type PackageWeight struct {
	UnitOfMeasurement UnitOfMeasurement `json:"UnitOfMeasurement"`
	WeightString      string            `json:"Weight"`
	Weight            float64
}

// UnitOfMeasurement ...
type UnitOfMeasurement struct {
	Code string `json:"Code"`
}

// ReferenceNumber ...
type ReferenceNumber struct {
	Code  string `json:"Code"`
	Value string `json:"Value"`
}
