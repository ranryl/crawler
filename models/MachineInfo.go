package models

// MachineInfo struct
type MachineInfo struct {
	ID      int64  `xorm:"bigint notnull pk autoincr id"`
	IP      string `xorm:"varchar(50) notnull ip"`
	Port    int32  `xorm:"int notnull default 0 port"`
	Name    string `xorm:"varchar(250) notnull default '' name"`
	Address string `xorm:"varchar(500) notnull default '' address"`
	Desc    string `xorm:"varchar(100) notnull default '' desc"`
	IsInner byte   `xorm:"tinyint(3) notnull default 0 is_inner"`
}

// TableName 指定表名
func (tb *MachineInfo) TableName() string {
	return "tb_machhine_info"
}
