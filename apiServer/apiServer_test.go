package apiServer

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"

	"github.com/techievee/frankieone/customLogger/debugcore"
	"github.com/techievee/frankieone/helper"
)

// Init Config
var (
	Config     *viper.Viper
	ConfigPath = "./../config"
	G          *gin.Engine
	ApiServer  *APIServer
)

func TestMain(m *testing.M) {

	var err error

	helper.ParseFlags(ConfigPath)

	Config, err = helper.LoadConfig()
	if err != nil || Config == nil {
		println("Failed to load config")
		os.Exit(1)
	}

	G = gin.Default()
	ApiServer = NewRestAPI("test", Config, &debugcore.NoOpsLogger{})

	c := m.Run()
	os.Exit(c)
}

func TestStartAPI(t *testing.T) {

	ApiServer.StartTLSServer()

}
