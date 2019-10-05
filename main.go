package main

import (
	"crawler/bases"
	"crawler/conf"
	"crawler/controllers"
	"crawler/utils"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func main() {

	basePath := "config" + string(os.PathSeparator)

	// 加载db配置文件及初始化
	dbConfFile, err := utils.GetFilePath(basePath + "database.yaml")
	utils.PanicError(err)
	dbConf := conf.DataBase{}
	err = utils.BindYamlConf(&dbConf, dbConfFile)
	utils.PanicError(err)
	engine, err := xorm.NewEngine(dbConf.DbEngine, dbConf.CrawlerMaster)
	bases.Orm["default"] = engine
	utils.PanicError(err)

	// 加载cache配置文件及初始化连接池
	cacheFile, err := utils.GetFilePath(basePath + "cache.yaml")
	utils.PanicError(err)
	cacheConf := conf.Cache{}
	err = utils.BindYamlConf(&cacheConf, cacheFile)
	utils.PanicError(err)
	_ = utils.NewRedisPool(cacheConf.Master)

	// 加载app配置文件
	appFile, err := utils.GetFilePath(basePath + "app.yaml")
	utils.PanicError(err)
	appConf := conf.App{}
	err = utils.BindYamlConf(&appConf, appFile)
	utils.PanicError(err)
	// 加载log配置文件
	logFile, err := utils.GetFilePath(basePath + "log.yaml")
	utils.PanicError(err)
	logConf := conf.Log{}
	err = utils.BindYamlConf(&logConf, logFile)
	utils.PanicError(err)
	curPath, err := utils.GetWdPath()
	utils.PanicError(err)

	route := gin.New()
	route.Use(utils.GinLogger(curPath, logConf))
	RegisterRouter(route)
	route.Use(gin.Logger())
	route.GET("/", func(c *gin.Context) {
		log.Println("print log")
	})
	route.Run(appConf.ServerPort)
}

// RegisterRouter 路由注册
func RegisterRouter(router *gin.Engine) {
	new(controllers.MachineInfoController).Router(router)
	new(controllers.CityController).Router(router)
}
