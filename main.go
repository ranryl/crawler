package main

import (
	"crawler/bases"
	"crawler/conf"
	"crawler/controllers"
	"crawler/utils"
	"io/ioutil"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"gopkg.in/yaml.v2"
)

func main() {

	currentPath, err := utils.GetCurrentPath()
	utils.PanicError(err)
	dbConfFile := currentPath + string(os.PathSeparator) +
		"config" + string(os.PathSeparator) + "local" + string(os.PathSeparator) + "database.yaml"
	dbYaml, err := ioutil.ReadFile(dbConfFile)
	utils.PanicError(err)
	dbConf := conf.DataBase{}
	err = yaml.Unmarshal(dbYaml, &dbConf)
	utils.PanicError(err)
	engine, err := xorm.NewEngine(dbConf.DbEngine, dbConf.CrawlerMaster)
	bases.Orm["default"] = engine
	utils.PanicError(err)
	route := gin.Default()
	RegisterRouter(route)

	route.Run(":8080")
}

// RegisterRouter 路由注册
func RegisterRouter(router *gin.Engine) {
	new(controllers.MachineInfoController).Router(router)
	new(controllers.CityController).Router(router)
}
