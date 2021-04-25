package main

import (
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/spf13/viper"

	"github.com/techievee/frankieone/frankieoneHelper"
	"github.com/techievee/frankieone/frankieoneLogger"
)

var AppFramework *gin.Engine

// Global Variable used by all the other packages

func main() {

	var (
		config     *viper.Viper
		configPath = "./config"
		err        error
	)

	frankieoneHelper.ParseFlags(configPath)
	glog.Infof("Starting Tester SDK API")

	// Configuration Initialization
	// Parse flag should have loaded the config, if its not loaded then load it via the loadconfig functions
	// Load the configuration files
	if config == nil {
		if config, err = frankieoneHelper.LoadConfig(); err != nil || config == nil {
			glog.Error("Failed to load config") // REVIEW : Should this panic
			return
		}
	}

	// Logger Initialization
	// Logger in injected to all the service, to maintain the logs
	env := config.GetString("app.app_env")
	frankieoneLogger := frankieoneLogger.NewLogger(env, frankieoneLogger.WithServiceName("test-sdk-api"))
	frankieoneLogger.Debug("FrankieoneLogger successfuly configured")

	// Starting the API framework for serving the Prodcut
	frankieoneLogger.Debug("Initializing the Rest Framework")
	// //restAPI := apiServer.NewRestAPI(env, config, frankieoneLogger)

	frankieoneLogger.Debug("Starting Products API Service")
	// startProductsService(config, db, restAPI, frankieoneLogger)

	//go restAPI.StartTLSServer()
	quit := make(chan bool)
	<-quit

}

/*func startProductsService(config *viper.Viper, db *database.DB, restAPI *apiServer.APIServer, logger debugcore.Logger) {
	ps := productService.NewProductService(config, db, restAPI, logger)
	ps.SetupService()
}
*/
