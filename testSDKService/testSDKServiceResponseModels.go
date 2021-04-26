package testSDKService

//Everyone gets a puppy if the SDK output is good
// swagger:model PuppyObject
type PuppyObject struct {

	// Everyone gets a puppy if the SDK output is good
	//
	// required: true
	// example: true
	// default: true
	Puppy bool `json:"puppy"`
}

// swagger:model ErrorObject
type ErrorObject struct {

	// Description of what went wrong (if we can tell)
	//
	// required: false
	Code int `json:"code"`

	// Description of what went wrong (if we can tell)
	//
	// example: Everything is wrong. Go fix it
	// required: false
	Message string `json:"message"`
}
