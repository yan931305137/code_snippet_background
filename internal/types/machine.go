package types

import "time"

type MachineInfo struct {
	Id                  int       `json:"Id"`
	SerialNumber        string    `json:"SerialNumber"`
	Manufacturer        string    `json:"Manufacturer"`
	Model               string    `json:"Model"`
	PurchaseDate        time.Time `json:"PurchaseDate"`
	WarrantyExpiryDate  time.Time `json:"WarrantyExpiryDate"`
	Location            string    `json:"Location"`
	Status              int       `json:"Status"`
	LastMaintenanceDate time.Time `json:"LastMaintenanceDate"`
	UserId              string    `json:"UserId"`
	Comment             string    `json:"Comment"`
}
