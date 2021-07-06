package model
type Address struct {
	ID         int64  `gorm:"column:id" json:"id" form:"id"`
	Name       string `gorm:"column:name" json:"name" form:"name"`
	UserId     int64  `gorm:"column:user_id" json:"user_id" form:"user_id"`
	CountryId  int64  `gorm:"column:country_id" json:"country_id" form:"country_id"`
	ProvinceId int64  `gorm:"column:province_id" json:"province_id" form:"province_id"`
	CityId     int64  `gorm:"column:city_id" json:"city_id" form:"city_id"`
	DistrictId int64  `gorm:"column:district_id" json:"district_id" form:"district_id"`
	Address    string `gorm:"column:address" json:"address" form:"address"`
	Mobile     string `gorm:"column:mobile" json:"mobile" form:"mobile"`
	IsDefault  int64  `gorm:"column:is_default" json:"is_default" form:"is_default"`
	IsDelete   int64  `gorm:"column:is_delete" json:"is_delete" form:"is_delete"`
}
