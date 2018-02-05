package uploader

type ShipperAdd struct {
	Data    *ShipperAddData `json:"data"`
	Success string          `json:"success"`
}

type ShipperAddData struct {
	ID        int    `json:"id"`
	AccessKey string `json:"access_key"`
}

type ShipperDelete struct {
	Success string `json:"success"`
}
