package testSDKService

import (
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"

	"github.com/techievee/frankieone/customLogger/debugcore"
	"github.com/techievee/frankieone/helper"
)

// Init Config
var (
	Config     *viper.Viper
	ConfigPath = "./../config"
	G          *gin.Engine
	ts         *testSDKServiceHandlers
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
	ts = &testSDKServiceHandlers{
		Config: Config,
		Logger: &debugcore.NoOpsLogger{},
	}
	G.POST("/isgood", ts.IsGood)
	c := m.Run()
	os.Exit(c)
}

func TestIsGoodService(t *testing.T) {

	testCases := []struct {
		testCaseName string
		body         string
		statusCode   int
	}{
		{
			"singledevice",
			`[
				  {
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "string",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  }
				]`,
			200,
		},
		{
			"mutipledevice",
			`[
				  {
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "string",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  },
				  {
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "string1",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue": "1.23.45.123",
						"kvpType": "general.string"
					  },
					   {
						"kvpKey": "ip.address",
						"kvpValue": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  }
				]`,
			200,
		},
		{
			"invalidCheckType",
			`[
				  {
					"checkType": "DEVICE1",
					"activityType": "SIGNUP",
					"checkSessionKey": "string",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  }
				]`,
			500,
		},
		{
			"invalidActivityTpe",
			`[
				  {
					"checkType": "DEVICE",
					"activityType": "SIGNUP1",
					"checkSessionKey": "string",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  }
				]`,
			500,
		},
		{
			"uniqueSessionKeyTest",
			`[
				  {
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "unique1",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  },
{
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "unique1",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  }
				]`,
			500,
		},
		{
			"invalidKvpType",
			`[
				  {
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "unique1",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.stringinvalid"
					  }
					]
				  },
{
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "unique1",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.string"
					  }
					]
				  }
				]`,
			500,
		},
		{
			"fieldignored",
			`[
				  {
					"checkType": "DEVICE",
					"activityType": "SIGNUP",
					"checkSessionKey": "string",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpType": "general.string"
					  }
					]
				  },
					]
				  }
				]`,
			500,
		},
		{
			"activityDataIgnored",
			`[
				  {
					"checkType": "DEVICE",
					"checkSessionKey": "string",
					"activityData": [
					  {
						"kvpKey": "ip.address",
						"kvpValue1": "1.23.45.123",
						"kvpType": "general.stringinvalid"
					  }
					]
				  },
					]
				  }
				]`,
			500,
		},
		{
			"emptyarray",
			`[
				
				]`,
			500,
		},
	}
	for _, value := range testCases {

		b := strings.NewReader(value.body)
		// Perform a GET request with that handler.
		w := mockRequest(G, http.MethodPost, "/isgood", b)
		assert.Equal(t, value.statusCode, w.Code)

	}

}

func mockRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
