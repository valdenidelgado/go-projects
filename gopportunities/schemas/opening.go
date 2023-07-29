package schemas

import "gorm.io/gorm"

type Opening struct {
	gorm.Model
	Role     string
	Company  string
	Location string
	Link     string
	Remote   bool
	Salary   int64
}
