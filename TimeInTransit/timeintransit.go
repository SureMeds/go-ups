package tnt

// APIRequest ...
type APIRequest struct {
	Security             Security             `json:"Security"`
	TimeInTransitRequest TimeInTransitRequest `json:"TimeInTransitRequest"`
}

// Security ...
type Security struct {
	UsernameToken         UsernameToken         `json:"UsernameToken,omitempty"`
	UPSServiceAccessToken UPSServiceAccessToken `json:"UPSServiceAccessToken,omitempty"`
}

// UsernameToken ...
type UsernameToken struct {
	Username string `json:"Username,omitempty"`
	Password string `json:"Password,omitempty"`
}

// UPSServiceAccessToken ...
type UPSServiceAccessToken struct {
	AccessLicenseNumber string `json:"AccessLicenseNumber,omitempty"`
}

// TimeInTransitRequest ...
type TimeInTransitRequest struct {
	Request         Request        `json:"Request,omitempty"`
	ShipFrom        Address        `json:"ShipFrom,omitempty"`
	ShipTo          Address        `json:"ShipTo,omitempty"`
	Pickup          Pickup         `json:"Pickup,omitempty"`
	ShipmentWeight  ShipmentWeight `json:"ShipmentWeight,omitempty"`
	MaximumListSize string         `json:"MaximumListSize,omitempty"`
}

// Request ...
type Request struct {
	RequestOption string `json:"RequestOption,omitempty"`
	TransactionReference
}

// TransactionReference ...
type TransactionReference struct {
	CustomerContext       string `json:"CustomerContext,omitempty"`
	TransactionIdentifier string `json:"TransactionIdentifier,omitempty"`
}

// Address ...
type Address struct {
	StateProvinceCode string `json:"StateProvinceCode,omitempty"`
	CountryCode       string `json:"CountryCode,omitempty"`
	Country           string `json:"Country,omitempty"`
	PostalCode        string `json:"PosatalCode,omitempty"`
}

// Pickup ...
type Pickup struct {
	Date string `json:"Date,omitempty"`
}

// ShipmentWeight ...
type ShipmentWeight struct {
	UnitOfMeasurement UnitOfMeasurement `json:"UnitOfMeasurement,omitempty"`
	Weight            string            `json:"Weight,omitempty"`
}

// TimeInTransitResponse ...
type TimeInTransitResponse struct {
	Response        Reponse         `json:"Response,omitempty"`
	TransitResponse TransitResponse `json:"TransitResponse,omitempty"`
}

// Response ...
type Response struct {
	ResponseStatus       ResponseStatus       `json:"ResponseStatus,omitempty"`
	TransactionReference TransactionReference `json:"TransactionReference,omitempty"`
}

// ResponseStatus ...
type ResponseStatus struct {
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
}

// TransitResponse ...
type TransitResponse struct {
	ShipFrom        Address        `json:"ShipFrom,omitempty"`
	ShipTo          Address        `json:"ShipTo,omitempty"`
	PickupDate      string         `json:"PickupDate,omitempty"`
	ShipmentWeight  ShipmentWeight `json:"ShipmentWeight,omitempty"`
	MaximumListSize string         `json:"MaximumListSize,omitempty"`
	ServiceSummary  ServiceSummary `json:"ServiceSummary,omitempty"`
	AutoDutyCode    string         `json:"AutoDutyCode,omitempty"`
	Disclaimer      string         `json:"Disclaimer,omitempty"`
}

// ServiceSummary ...
type ServiceSummary struct {
	Service          Service          `json:"Service,omitempty"`
	EstimatedArrival EstimatedArrival `json:"EstimatedArrival,omitempty"`
}

// Service ...
type Service struct {
	Code        string `json:"Code,omitempty"`
	Description string `json:"Description,omitempty"`
}

// EstimatedArrival ...
type EstimatedArrival struct {
	Arrival               Arrival `json:"Arrival,omitempty"`
	BusinessDaysInTransit string  `json:"BusinessDaysInTransit,omitempty"`
	DayOfWeek             string  `json:"DayOfWeek,omitempty"`
	CustomerCenterCutoff  string  `json:"CustomerCenterCutoff,omitempty"`
	TotalTransitDays      string  `json:"TotalTransitDays,omitempty"`
}

// Arrival ...
type Arrival struct {
	Date string `json:"Date,omitempty"`
	Time string `json:"Time,omitempty"`
}

// Pickup ...
type Pickup struct {
	Date string `json:"Date,omitempty"`
	Time string `json:"Time,omitempty"`
}
