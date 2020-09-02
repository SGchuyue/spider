package model

type Title struct {
	*Model

	Name      string `json:"name" bson:"name"`
	TatleTime string `json:"title_name" bson:"title_name"`
}

func (t Title) TableName() string {
	return "Vu_title"
}
