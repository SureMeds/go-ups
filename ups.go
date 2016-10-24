package ups

import (
	"bytes"
	"json"
	"net/http"
	"strconv"
	"time"

	"github.com/SureMeds/utils-go/ups/TimeInTransit"
	"github.com/SureMeds/utils-go/ups/tracking"
)

// Client ...
type Client struct {
	Username      string `json:"username"`
	Password      string `json:"password"`
	AccessLicense string `json:"access_license"`
}

// Track ...
func (c Client) Track(trackingNumber string) (tracking.TrackResponse, error) {
	request := tracking.APIRequest{
		UPSSecurity: tracking.UPSSecurity{
			UsernameToken: tracking.UsernameToken{
				Username: c.Username,
				Password: c.Password},
			ServiceAccessToken: tracking.ServiceAccessToken{
				AccessLicenseNumber: c.AccessLicense}},
		TrackRequest: tracking.TrackRequest{
			Request: Request{
				RequestOption: "1",
				TransactionReference: tracking.TransactionReference{
					CustomerContext: "SureMeds shipment tracking"}},
			InquiryNumber: trackingNumber}}
	url := "https://onlinetools.ups.com/rest/Track"
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var response tracking.TrackReponse
	// throw error if json not decoded improperly
	if err = json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}

// GetTimeInTransitEstimate ...
func (c Client) GetTimeInTransitEstimate(from tnt.Address, to tnt.Address, weight float64) (tnt.TimeInTransitResponse, error) {
	request := tnt.APIRequest{
		Security: tnt.Security{
			UsernameToken: tnt.UsernameToken{
				Username: c.Username,
				Password: c.Password},
			ServiceAccessToken: tnt.ServiceAccessToken{
				AccessLicenseNumber: c.AccessLicense}},
		TimeInTransitRequest: tnt.TimeInTransitRequest{
			ShipFrom: from,
			ShipTo:   to,
			Pickup: tnt.Pickup{
				Date: time.Now().Format("20060102")},
			ShipmentWeight: tnt.ShipmentWeight{
				UnitOfMeasurement: tnt.UnitOfMeasurement{
					Code: "LB"},
				Weight: strconv.FormatFloat(weight, "f", -1, 64)}}}
	url := "https://onlinetools.ups.com/rest/TimeInTransit"
	jsonStr, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	r, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	var response tnt.TimeInTransitResponse
	// throw error if json not decoded improperly
	if err = json.NewDecoder(r.Body).Decode(&response); err != nil {
		return nil, err
	}
	return response, nil
}
