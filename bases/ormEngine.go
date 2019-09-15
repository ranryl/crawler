package bases

import (
	"github.com/go-xorm/xorm"
)

// Orm map
var Orm map[string]*xorm.Engine = make(map[string]*xorm.Engine)

// SetEngine 设置db engine
func SetEngine(key string, e *xorm.Engine) {
	Orm[key] = e
}

// GetEngine 获取engine
func GetEngine(keys ...string) *xorm.Engine {
	if len(keys) == 0 {
		return Orm["default"]
	}
	return Orm[keys[0]]
}
