package uploader

// ShipperAdd - response for Shipper:Add
type ShipperAdd struct {
	Data    *ShipperAddData `json:"data"`
	Success string          `json:"success"`
}

// ShipperAddData - intermediate object for Shipper:Add:Data
type ShipperAddData struct {
	ID        int    `json:"id"`
	AccessKey string `json:"access_key"`
}

// ShipperDelete - response for Shipper:Delete
type ShipperDelete struct {
	Success string `json:"success"`
}
