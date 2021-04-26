package testSDKService

import (
	"github.com/spf13/viper"

	"github.com/techievee/frankieone/apiServer"
	"github.com/techievee/frankieone/customLogger/debugcore"
)

// SDK Service structure
type testSDKService struct {
	Config                 *viper.Viper
	testSDKServiceHandlers *testSDKServiceHandlers
	RestAPI                *apiServer.APIServer
	Logger                 debugcore.Logger
}

func NewTestSDKService(config *viper.Viper, restAPI *apiServer.APIServer, logger debugcore.Logger) *testSDKService {

	testSDKServiceHandlers := &testSDKServiceHandlers{
		Config: config,
		Logger: logger,
	}
	// Create a new controller for the Product
	return &testSDKService{
		Config:                 config,
		testSDKServiceHandlers: testSDKServiceHandlers,
		RestAPI:                restAPI,
		Logger:                 logger,
	}
}

// Bootstrap the service
func (ts *testSDKService) SetupService() {

	ts.Logger.Debug("Test SDK Service Starting")
	ts.LoadRoutes()

}

// Load all the routes that is related to this service
func (ts *testSDKService) LoadRoutes() {

	ts.Logger.Debug("Setting up routes for Test SDK Services")

	testSDKRoutes := ts.RestAPI.GinFramework.Group("/")

	// TestSDK Service Routes
	// All the routes were added here
	// swagger:route POST /isgood Test
	//
	// Check that the output of the universal SDK is fine
	//
	// Simple check to see if the service is running smoothly.
	//
	// The body that you post here should be the direct output of the SDK's GetSessionData call.
	//
	//     Consumes:
	//     - application/json
	//
	//     Produces:
	//     - application/json
	//
	//     Security:
	//
	//     Responses:
	//       200: PuppyObject
	//       500: ErrorObject
	testSDKRoutes.POST("/isgood", ts.testSDKServiceHandlers.IsGood)
	ts.Logger.Debug("Routes were successfully configured")
}
