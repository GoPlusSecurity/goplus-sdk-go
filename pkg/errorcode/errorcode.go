package errorcode

const (
	SUCCESS                        = 1
	DATA_PENDING_SYNC              = 2
	ADDRESS_FORMAT_ERROR           = 2004
	CHAIN_NOT_SUPPORTED            = 2018
	NON_CONTRACT_ADDRESS           = 2020
	CONTRACT_INFO_NOT_FOUND        = 2021
	DAPP_NOT_FOUND                 = 2026
	ABI_NOT_FOUND                  = 2027
	ABI_FUNCTION_UNSUPPOERT        = 2028
	APP_KEY_NOT_EXIST              = 4010
	SIGNATURE_EXPIRATION           = 4011
	SIGNATURE_VERIFICATION_FAILURE = 4012
	FREQUENCY_OVER_LIMIT           = 4029
	INVALID_TOKEN                  = 4022
	TOKEN_NOT_FOUND                = 4023
	SERVER_ERROR                   = 5000
	PARAM_ERROR                    = 5006
)
