package ocpi211

type ImageCategory string

const (
	ImageCategoryCharger  ImageCategory = "CHARGER"
	ImageCategoryEntrance ImageCategory = "ENTRANCE"
	ImageCategoryLocation ImageCategory = "LOCATION"
	ImageCategoryNetwork  ImageCategory = "NETWORK"
	ImageCategoryOperator ImageCategory = "OPERATOR"
	ImageCategoryOther    ImageCategory = "OTHER"
	ImageCategoryOwner    ImageCategory = "OWNER"
)

type Image struct {
	Url       string        `json:"url"`
	Thumbnail *string       `json:"thumbnail"`
	Category  ImageCategory `json:"category"`
	Type      string        `json:"type"`
	Width     *int          `json:"width"`
	Height    *int          `json:"height"`
}

type BusinessDetails struct {
	Name    string  `json:"name"`
	Website *string `json:"website"`
	Logo    *Image  `json:"logo"`
}
