package models

// City struct
type City struct {
	ID       int64  `xorm:"bigint notnull pk autoincr id"`
	CityID   int64  `xorm:"bigint city_id"`
	CityName string `xorm:"varchar(30) city_name"`
	ProvID   int32  `xorm:"int prov_id"`
}

// TableName ...
func (c *City) TableName() string {
	return "td_city"
}
