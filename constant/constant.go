package constant

import (
	"net/http"

	"github.com/more-than-code/messaging/pb"
	"google.golang.org/grpc/codes"
)

const (
	UserRoleCommon   = 0
	UserRoleRetailer = 1
	UserRoleEditor   = 2
	UserRoleAdmin    = 3
)

const (
	PurposeRegistration        = 0
	PurposeLogin               = 1
	PurposeResettingPassword   = 2
	PurposeBindingPhoneOrEmail = 3
)

const (
	SmsSubject             = "「毛孩街」驗證碼"
	SmsRegistration        = "「毛孩街」註冊的驗證碼(5分鐘內有效)"
	SmsLogin               = "「毛孩街」登錄的驗證碼(5分鐘內有效)"
	SmsResettingPassword   = "「毛孩街」重置密碼的驗證碼(5分鐘內有效)"
	SmsBindingPhoneOrEmail = "「毛孩街」綁定手機或郵箱的驗證碼(5分鐘內有效)"
)

const (
	ExHttpStatusAuthenticationFailure = 460
	ExHttpStatusBusinessError         = 461
)

const (
	CodeExistingPhoneOrEmail    = 1000
	CodeNonexistingPhoneOrEmail = 1001

	CodeAuthenticationFailure = 1002
	CodeUnauthorized          = 1003
	CodeNotAllowed            = 1004

	CodeExpiredVerificationCode              = 1100
	CodeWrongVerificationCode                = 1101
	CodeTooFrequentlySendingVerificationCode = 1102
	CodeNeedingResendingVerificationCode     = 1103
	CodeMaximumAttemptsOnVerificationCode    = 1104

	CodeInvalidParameters    = 1200
	CodeWrongEmailOrPassword = 1201
	CodeDeletedUser          = 1202

	CodeDatabaseOperationFailure = 2000
	CodeDatabaseDocumentNotFound = 2001
	CodeDataUnmarshalingFailure  = 2002
	CodeDataMarshalingFailure    = 2003

	// checkin
	CodeCheckInRecordExisted        = 4000
	CodeCheckInInvalidTime          = 4001
	CodeCheckInInvalidGeolocation   = 4002
	CodeCheckInInvalidCounter       = 4003
	CodeCheckInOutOfGeoRangeCounter = 4004

	// order, points
	CodeInsufficientPoints    = 4100
	CodeInsufficientInventory = 4101

	CodeServiceInternalServerError = 5000
	CodeServiceNotImplemented      = 5001
	CodeServiceBadGateway          = 5002
	CodeServiceServiceUnavailable  = 5003
	CodeServiceGatewayTimeout      = 5004
)

const (
	MsgExistingPhoneOrEmail    = "phone or email already exists"
	MsgNonexistingPhoneOrEmail = "phone or email not exists"
	MsgAuthenticationFailure   = "authenticating user failure"
	MsgUnauthorized            = "unauthorized"
	MsgNotAllowed              = "not allowed"

	MsgExpiredVerificationCode              = "expired verification code"
	MsgWrongVerificationCode                = "wrong verification code"
	MsgDeletedUser                          = "deleted user"
	MsgTooFrequentlySendingVerificationCode = "too frequently sending code"
	MsgNeedingResendingVerificationCode     = "needing resending verification code"
	MsgMaximumAttemptsOnVerificationCode    = "maximum attempts on verification code"

	MsgInvalidParamters     = "invalid parameter(s)"
	MsgWrongEmailOrPassword = "wrong email or password"

	MsgDatabaseOperationFailure = "database operation failure"
	MsgDatabaseDocumentNotFound = "document not found in database"
	MsgDataUnmarshalingFailure  = "data unmarshaling failure"
	MsgDataMarshalingFailure    = "data marshaling failure"

	// checkin
	MsgCheckInRecordExisted        = "check-in record existed"
	MsgCheckInInvalidTime          = "check-in invalid time"
	MsgCheckInInvalidGeolocation   = "check-in invalid geolocation"
	MsgCheckInInvalidCounter       = "check-in invalid counter"
	MsgCheckInOutOfGeoRangeCounter = "check-in out of geo range"

	// voucher ownership
	MsgInsufficientPoints    = "insufficient points"
	MsgInsufficientInventory = "insufficient inventory"
)

func TranslateToHttpStatusCode(code int) int {
	switch code {
	case int(codes.Unavailable):
		return http.StatusServiceUnavailable
	case int(CodeDatabaseDocumentNotFound):
		return http.StatusNotFound
	}

	if code < 5000 {
		return ExHttpStatusBusinessError
	}

	return http.StatusInternalServerError
}

func TranslateToServiceCode(code int) int {
	switch code {
	case int(codes.Unavailable):
		return CodeServiceServiceUnavailable
	default:
		return code
	}
}

func TranslateVerificationCodeValidationStatusToServiceCode(status pb.VerificationCodeValidationStatus) int {
	switch status {
	case pb.VerificationCodeValidationStatus_VERIFICATION_CODE_VALIDATION_STATUS_EXPIRED:
		return CodeExpiredVerificationCode
	case pb.VerificationCodeValidationStatus_VERIFICATION_CODE_VALIDATION_STATUS_INVALID:
		return CodeWrongVerificationCode
	case pb.VerificationCodeValidationStatus_VERIFICATION_CODE_VALIDATION_STATUS_MAXIMUM_ATTEMPTS:
		return CodeMaximumAttemptsOnVerificationCode
	default:
		return int(status)
	}
}

type RedisChannel string

const (
	RcCreateComment     = "createComment"
	RcUpdateUserBasics  = "updateUserBasics"
	RcUpdateTransaction = "update_transaction"
)

type RedisKey string

const (
	RkCreateCommentPetIds  = "createCommentPetIds"
	RkUpdateUserBasicsIds  = "UpdateUserBasicsIds"
	RkUpdateTransactionIds = "UpdateTransactionIds"
)
