package bidder

type Context struct{}

type DisplayFormat struct {
	Width  uint16 `json:"w"`
	Height uint16 `json:"h"`
}
type Display struct {
	PlacementPosition uint16        `json:"pos"`
	Interstitial      uint8         `json:"instl"`
	Format            DisplayFormat `json:"displayfmt"`
}
type Placement struct {
	TagId   string  `json:"tagid"`
	Display Display `json:"display"`
}
type Spec interface {
}
type Item struct {
	Id               string    `json:"id"`
	Qty              uint16    `json:"qty"`
	BidFloor         float32   `json:"flr"`
	BidFloorCurrency string    `json:"flrcur"`
	Spec             Placement `json:"spec"`
}

type BidRequest struct {
	Id          string   `json:"id"`
	AuctionType uint16   `json:"at"`
	Currency    []string `json:"cur"`
	Seat        []string `json:"seat"`
	Context     Context  `json:"context"`
	Item        []Item   `json:"item"`
}
