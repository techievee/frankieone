package testSDKService

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"

	"github.com/techievee/frankieone/customLogger/debugcore"
)

// Handle for the test sdk services
type testSDKServiceHandlers struct {
	Config *viper.Viper
	Logger debugcore.Logger
}

// swagger:operation POST /isgood Test
//
// Simple check to see if the service is running smoothly.
//
//The body that you post here should be the direct output of the SDK's GetSessionData call
//
//
// ---
// produces:
// - application/json
// parameters:
// - name: deviceCheckDetails
//   in: body
//   description: This is what the JSON that is exported from the SDK should look like. It is an array of objects that contain the details from each different provider wrapped up in the Universal SDK.
//   required: true
//   type: array
//   items:
//     "$ref": "#/definitions/DeviceCheckDetailsObjects"
// responses:
//   '200':
//     description: The data is fine. No issues, and everyone gets a puppy.
//     schema:
//       "$ref": "#/definitions/PuppyObject"
//   '500':
//     description: The system is presently unavailable, or running in a severely degraded state. Check the error message for details
//     schema:
//       "$ref": "#/definitions/ErrorObject"
//
func (tsh *testSDKServiceHandlers) IsGood(c *gin.Context) {

	devices := DeviceCheckDetailsObjectCollection{}
	var err error
	var validationErrors []string

	if err = c.ShouldBindJSON(&devices.DeviceCheckDetailsObjects); err != nil {
		c.JSON(http.StatusInternalServerError, ErrorObject{Code: 500, Message: err.Error()})
		return
	}

	v := validator.New()
	if err := v.Struct(devices); err != nil {
		errors, _ := err.(validator.ValidationErrors)
		for _, err := range errors {
			validationErrors = append(validationErrors, err.Error())
		}
	}

	// validate the DeviceCheckDetailsObject
	for _, d := range devices.DeviceCheckDetailsObjects {
		if err := v.Struct(d); err != nil {
			errors, _ := err.(validator.ValidationErrors)
			for _, err := range errors {
				validationErrors = append(validationErrors, err.Error())
			}
		}

		// Iterate every possible KVPDate for Validation
		for _, kvp := range d.ActivityData {
			if err := v.Struct(kvp); err != nil {
				errors, _ := err.(validator.ValidationErrors)
				for _, err := range errors {
					validationErrors = append(validationErrors, err.Error())
				}
			}
		}
	}

	if len(validationErrors) > 0 {
		c.JSON(http.StatusInternalServerError, ErrorObject{
			Code:    http.StatusInternalServerError,
			Message: strings.Join(validationErrors, ","),
		})
		return
	}

	response := &PuppyObject{Puppy: true}
	c.JSON(http.StatusOK, &response)
}
