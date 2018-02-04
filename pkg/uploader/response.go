package uploader

type ShipperAdd struct {
	Data    *ShipperAddData `json:"data"`
	Success string          `json:"success"`
}

type ShipperAddData struct {
	ID        int    `json:"id"`
	AccessKey string `json:"access_key"`
}

// {
// 	"data":{
// 		"id":2,
// 		"access_key":"5842f0b6-7289-47c2-8251-cb77c313b6ca",
// 		"created_at":"0001-01-01T00:00:00Z",
// 		"updated_at":"0001-01-01T00:00:00Z"
// 		},
// 	"success":"true"
// }
