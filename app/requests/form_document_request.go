package requests

type FormDocumentRequest struct {
	DocumentID int64   `json:"ID"`
	Font       string  `json:"Font"`
	FontSize   float64 `json:"FontSize"`
	Interval   float64 `json:"Interval"`
}
