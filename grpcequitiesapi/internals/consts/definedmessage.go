package consts

const (
	InvalidCode       = "Invalid code input request "
	InvalidUpdateType = "Invalid update type input request "

	InvalidUpdateData = "Invalid data found for updating request"

	ErrorUpdateType         = "Update failed for merchant code %v"
	ErrorUpdateTypeNotFound = "Update type not found as %v"

	ErrorUpdateMember = "Update failed for members code %v, email %v"

	ErrorDataNotFoundCode = "Merchant data not found for given code %v"

	ErrorUserNotFoundCode = "User not found for given code %v"

	InvalidUserDocId    = "Invalid user code input request "
	InvalidMerchantCode = "Invalid client code input request "
	InvalidAppCode      = "Invalid client app code input request "

	UserLoginSuccess       = "User login successfully"
	UserAddedSuccess       = "User added successfully"
	TokenRegeneatedSuccess = "Token Regenerated successfully"

	MerchantAddedSuccess   = "Merchant added successfully"
	MerchantUpdatedSuccess = "Merchant updated successfully for code: %v"
	PageLimitMessage       = "Invalid page limit"
	SkipMessage            = "Invalid skip limit"

	ActiveStatus            = 1 // Active
	DeactiveStatus          = 0 // DeactiveStatus
	ArchiveStatus           = 9 // ArchiveStatus
	ParseLayoutISO   string = "2006-01-02"
	DateFormatLayout string = "02-01-2006"

	// MySQL
	DuplicateEntry           string = "Duplicate entry"
	ErrUserAlreadyExists     string = "Memeber already exists for merchant code: %v, email: %v"
	ErrMerchantAlreadyExists string = "Merchant already exists for merchant code: %v"
	// ErrUserAlreadyExists string = "User already exists for input"

	InvalidUserId                 = "invalid user id input request "
	InvalidOrderId                = "invalid order id input request "
	ErrorOrderDataNotFoundCode    = "Order Processed data not found for given userid %v"
	OrderProcessedUpdatedSuccess  = "Order Processed updated successfully for code: %v"
	ErrorOrderProcessedUpdateType = "Order Processed update failed for user_id %v, order_id %v"

	OrdrActiveStatus   = 1 // OrdrActiveStatus
	OrderPendingtatus  = 0 // OrderPendingtatus
	OrderArchiveStatus = 9 // OrderArchiveStatus
)
