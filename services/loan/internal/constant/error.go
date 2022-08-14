package constant

const (
	INTERNAL_SERVER_ERROR = 500

	INCORRECT_NAME            = 1001
	INCORRECT_IDENTITY_NUMBER = 1002
	INCORRECT_AMOUNT          = 1003
	INCORRECT_PURPOSE         = 1004
)

var ERROR_MAPPING = map[int]string{
	INTERNAL_SERVER_ERROR:     "Internal server error",
	INCORRECT_NAME:            "Incorrect name",
	INCORRECT_IDENTITY_NUMBER: "Incorrect identity number",
	INCORRECT_AMOUNT:          "Incorrect amount",
	INCORRECT_PURPOSE:         "Incorrect purpose",
}
