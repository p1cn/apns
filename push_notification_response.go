package apns

import ()

// The maximum number of seconds we're willing to wait for a response
// from the Apple Push Notification Service.
const TIMEOUT_SECONDS = 10

// This enumerates the response codes that Apple defines
// for push notification attempts.
var APPLE_PUSH_RESPONSES = map[uint8]string{
	0:   "NO_ERRORS",
	1:   "PROCESSING_ERROR",
	2:   "MISSING_DEVICE_TOKEN",
	3:   "MISSING_TOPIC",
	4:   "MISSING_PAYLOAD",
	5:   "INVALID_TOKEN_SIZE",
	6:   "INVALID_TOPIC_SIZE",
	7:   "INVALID_PAYLOAD_SIZE",
	8:   "INVALID_TOKEN",
	10:  "SHUTDOWN",
	255: "UNKNOWN",
}

// PushNotificationResponse contains the result of a PushNotification send.
// It will either contain a response string from Apple or an error (for
// example unable to open connection).
type PushNotificationResponse struct {
	Success          bool
	AppleResponse    *string
	Error            *error
	PushNotification *PushNotification
}

// Constructor.
func NewPushNotificationResponse(pn *PushNotification) (resp *PushNotificationResponse) {
	resp = new(PushNotificationResponse)
	resp.Success = false
	resp.PushNotification = pn
	return
}
