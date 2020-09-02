package model


type Title struct {
	*Model

	Name string `json:"Name"`
	TatleTime string `json:"title_name"`
}

func (t Title) TableName() string {
	return "Vu_title"
}
