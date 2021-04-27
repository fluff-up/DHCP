package data

const (
	OperationCodeUnknown     OperationCode = 0
	OperationCodeBootRequest OperationCode = 1
	OperationCodeBootReply   OperationCode = 2
)

const (
	descriptionOperationCodeUnknown     = "unknown"
	descriptionOperationCodeBootRequest = "request"
	descriptionOperationCodeBootReply   = "reply"
)

var (
	operationCodeToDescription = map[OperationCode]string{
		OperationCodeUnknown:     descriptionOperationCodeUnknown,
		OperationCodeBootRequest: descriptionOperationCodeBootRequest,
		OperationCodeBootReply:   descriptionOperationCodeBootReply,
	}
)

// OperationCode this is type for operation code
// See: RFC 1542 section 2.1
// The 'op' (opcode) field of the message must contain either the
// code for a BOOTREQUEST (1) or the code for a BOOTREPLY (2).
//
// Link: https://tools.ietf.org/html/rfc1542#section-2.1
type OperationCode int8

// Description return description of operation code
func (o OperationCode) Description() string {
	desc, exists := operationCodeToDescription[o]
	if !exists {
		return operationCodeToDescription[OperationCodeUnknown]
	}

	return desc
}
