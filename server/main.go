package main

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/wdana-cn/ginpro/common"
)

func InitConfig() {
	workDir, _ := os.Getwd()
	viper.SetConfigName("application")
	viper.SetConfigType("yml")
	viper.AddConfigPath(workDir + "/config")
	err := viper.ReadInConfig()
	if err != nil {

	}
}

func main() {
	InitConfig()
	db := common.InitDB()
	defer db.Close()

	r := gin.Default()
	r.Use(Cors())
	r = CollectRouter(r)
	port := viper.GetString("server.port")
	if port != "" {
		panic(r.Run(":" + port))
	}
	//panic(r.Run())
}
