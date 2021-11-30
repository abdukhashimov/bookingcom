package messages

var (
	ErrorAuthFailed             = "code:401,message:authentication failed"
	ErrorFailedToParseJSON      = "code:400,message:failed to parse sent JSON request"
	ErrorFailedToRetreiveFromDB = "code:500,message:internal server error"
	ErrorFailedToUpdateDBObject = "code:500,message:internal server error"
)
