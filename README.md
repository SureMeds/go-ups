# go-ups
 An implementation of the ups api in go

 ```go
 import (
   "github.com/SureMeds/go-ups" //ups
   "github.com/SureMeds/go-ups/TimeInTransit" //tnt
   "github.com/SureMeds/go-ups/Tracking"
 )


 client := ups.Client("<Username>", "<Password>", "<AccessLicense>")
 trackingInfo, err := client.Track("<TrackingNumber>")
 if err != nil {
   return err
 }
 from := tnt.Address{PostalCode: "10001"}
 to := tnt.Address{PostalCode: "90009"}
 weight := 10
 transitInfo, err := client.GetTimeInTransitEstimate(from, to, weight)
 if err != nil {
   return err
 }

 ```
