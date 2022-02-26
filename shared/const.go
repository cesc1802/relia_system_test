package shared

const (
	CmtMaxLength  = 9
	CccdMaxLength = 12
)

const (
	DBTypeUser    = 1
	DBTypeProduct = 2
)

type UploadType int

const (
	KeyCurrentUser = "current_user"
)

const (
	UserActive     = 1
	UserDeactivate = 0
)
