package model

type Model struct {
	ID         uint32 `gorm:"primary_key" json:"id"`
	CreateBy   string `json:"create-by"`
	Modifiedby string `json:"modified-by"`
	CreateOn   uint32 `json:"created_on"`
	ModifiedOn uint32 `json:"modified-on"`
	DeletedOn  uint32 `json:"deleted_on"`
	IsDel      uint8  `json:"is_del"`
}
