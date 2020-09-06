package models

type Model struct {
	ID         int `gorm:"primary_key" json:"id"`
	CreatedAt  int `json:"created_at"`
	ModifiedAt int `json:"modified_at"`
	Deleted    int `json:"modified_at"`
}

//func (ins Model) name() string {
//	return ins.+ " " + p.lastName
//}
