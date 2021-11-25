package messages

var (
	ErrorAuthFailed             = "code:401,message:authentication failed"
	ErrorFailedToParseJSON      = "code:400,message:failed to parse sent JSON request"
	ErrorFailedToRetreiveFromDB = "code:500,message:Internal Server error"
)
