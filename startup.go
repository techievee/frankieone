// Package main Frankie Financial Universal SDK Tester API (Internal Only)
//
// This API allows developers to test the Universal SDK output to ensure it looks right..
// The traditional Swagger view of this documentation can be found here:\n  - https://app.swaggerhub.com/apis-docs/FrankieFinancial/TestUniversalSDK
//
//     Schemes: https
//     Host: localhost:8081
//     Version: 1.0.3
//     Contact: Vinod Kumar jayarajan<vinodkumarjayarajan@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//
//     Security:
//
//     SecurityDefinitions:
// swagger:meta
package main

import (
	"github.com/golang/glog"
	"github.com/spf13/viper"

	"github.com/techievee/frankieone/apiServer"
	"github.com/techievee/frankieone/customLogger"
	"github.com/techievee/frankieone/customLogger/debugcore"
	"github.com/techievee/frankieone/helper"
	"github.com/techievee/frankieone/testSDKService"
)

func main() {

	var (
		config     *viper.Viper
		configPath = "./config"
		err        error
	)

	helper.ParseFlags(configPath)
	glog.Infof("Starting Tester SDK API")

	// Configuration Initialization
	// Parse flag should have loaded the config, if its not loaded then load it via the loadconfig functions
	// Load the configuration files
	if config == nil {
		if config, err = helper.LoadConfig(); err != nil || config == nil {
			glog.Error("Failed to load config") // REVIEW : Should this panic
			return
		}
	}

	// Logger Initialization
	// Logger in injected to all the service, to maintain the logs
	env := config.GetString("app.app_env")
	customLogger := customLogger.NewLogger(env, customLogger.WithServiceName("test-sdk-api"))
	customLogger.Debug("Logger successfuly configured")

	// Starting the rest API Server
	customLogger.Debug("Initializing the Rest Framework")
	restAPI := apiServer.NewRestAPI(env, config, customLogger)

	// Start testSDK Services
	customLogger.Debug("Starting Test SDK API Service")
	starttestSDKService(config, restAPI, customLogger)

	go restAPI.StartTLSServer()
	quit := make(chan bool)
	<-quit

}

// Start services one by one
func starttestSDKService(config *viper.Viper, restAPI *apiServer.APIServer, logger debugcore.Logger) {
	ts := testSDKService.NewTestSDKService(config, restAPI, logger)
	ts.SetupService()
}
