package testSDKService

// swagger:enum enumKVPType
type enumKVPType string

const (
	STRING  enumKVPType = "general.string"
	INTEGER enumKVPType = "general.integer"
	FLOAT   enumKVPType = "general.float"
	BOOL    enumKVPType = "general.bool"
)

//Individual key-value pair
// swagger:model KeyValuePairObject
type KeyValuePairObject struct {

	// Name of the data
	//
	// required: false
	// example: ip.address
	KvpKey string `json:"kvpKey"`

	// Value of the data
	//
	// example: 1.23.45.123
	KvpValue string `json:"kvpValue"`

	KvpType enumKVPType `json:"kvpType" validate:"oneof='general.string' 'general.integer' 'general.float' 'general.bool'"`
}

// swagger:enum checkType
type checkType string

const (
	DEVICE    checkType = "DEVICE"
	BIOMETRIC checkType = "BIOMETRIC"
	COMBO     checkType = "COMBO"
)

// swagger:enum activityType
type activityType string

const (
	SIGNUP       activityType = "SIGNUP"
	LOGIN        activityType = "LOGIN"
	PAYMENT      activityType = "PAYMENT"
	CONFIRMATION activityType = "CONFIRMATION"
)

// Contains any or all details we want to pass on to the device or biometric checking service as part of an activity or transaction. A transaction isn't just a payment, but can represent a number of different interaction types. See below for more
// swagger:model DeviceCheckDetailsObjects
type DeviceCheckDetailsObject struct {

	//Describes the type of check service we need to verify with. Choices are:
	//DEVICE: Services that will be checking device characteristics
	//BIOMETRIC: Services that will be checking biomentric characteristics
	//COMBO: If you're using a service that combines both device and biometric information, use this.
	//
	CheckType checkType `json:"checkType" binding:"required" validate:"oneof='DEVICE' 'BIOMETRIC' 'COMBO'"`

	//The type of activity we're checking. Choices are
	//SIGNUP: Used when an entity is signing up to your service
	//LOGIN: Used when an already registered entity is logging in to your service
	//PAYMENT: Used when you wish to check that all is well for a payment
	//CONFIRMATION: User has confirmed an action and you wish to double check they're still legitimate
	//You can also supply vendor specific activityTypes if you know them. To do this, make the first character an underscore _.
	//So for example, to use BioCatch's LOGIN_3 type, you can send "_LOGIN_3" as a value. Note, if you do this, there is no error checking on the Frankie side, and thus if you supply an incorrect value, the call will fail
	//
	ActivityType activityType `json:"activityType" binding:"required" validate:"oneof='SIGNUP' 'LOGIN' 'PAYMENT' 'CONFIRMATION'"`

	//The unique session based ID that will be checked against the service.
	//Service key must be unique or an error will be returned.
	//
	CheckSessionKey string `json:"checkSessionKey" binding:"required"`

	//A collection of loosely typed Key-Value-Pairs, which contain arbitrary data to be passed on to the verification services.
	//The API will verify that:
	//
	//the list of "Keys" provided are unique to the call (no double-ups)
	//that the Value provided matches the Type specified.
	//Should the verification fail, the error message returned will include information for each KVP pair that fails.
	//
	ActivityData []KeyValuePairObject `json:"activityData" validate:"required,min=1"`
}

// swagger:model DeviceCheckDetailsObjectCollection
type DeviceCheckDetailsObjectCollection struct {
	// Array of DeviceCheckDetailsObjects
	//
	//
	// required: true
	// swagger:model DeviceCheckDetailsObjectCollection
	DeviceCheckDetailsObjects []DeviceCheckDetailsObject `json:"DeviceCheckDetailsObjectCollection" binding:"required" validate:"required,min=1,unique=CheckSessionKey"`
}
