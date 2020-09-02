package model

type Article struct {
	*Model

	Title string `json:"title"`
	Time string `json:"time"`
	info string `json:"time"`
	Products string `json:"products"`
	CveID string `json:"CVE_ID"`
	Description string `json:"description"`
	types string `json:"types"`
	Url string `json:"url"`
	Solution string `json:"solution"`
	Patch string `json:"patch"`
	Information string `json:"information"`
	ReciveTime string `json:"recive_time"`
	ReceiptTime string `json:"receipt_time"`
	UpdateTime string `json:"update_time"`

}

func (a Article) TableName() string {
	return "Vu_article"
}
