package errhandle

type ErrCode int

//go:generate stringer -type=ErrCode
const (
	// NoError 表示没有发生错误。
	NoError ErrCode = iota

	// InnerError 表示内部错误，发生了系统内部的意外错误。
	InnerError

	// TokenError 表示 Token 错误，可能是因为 Token 无效或过期等问题。
	TokenError

	// ParamsError 表示参数错误，客户端发送的请求参数不正确。
	ParamsError

	// UserExists 表示用户已存在，用于注册或创建用户时，如果用户已经存在，则返回该错误。
	UserExists

	// UserNonExists 表示用户不存在，用于识别用户是否存在的情况。
	UserNonExists

	// PasswordInvalid 表示密码无效，可能是因为密码格式不正确或者密码错误等问题。
	PasswordInvalid

	// PermissionDenied 表示权限被拒绝，用户没有足够的权限执行某个操作。
	PermissionDenied

	// SexError 表示性别错误，用于表示性别字段的值不符合预期的格式或范围。
	SexError

	// NameExists 表示名称已存在，用于注册或创建某个实体时，如果名称已经存在，则返回该错误。
	NameExists

	// EmailFormatError 表示邮箱格式错误，用于表示邮箱字段的值不符合预期的格式。
	EmailFormatError

	// PasswordFormatError 表示密码格式不符合要求。
	PasswordFormatError

	// FileSizeTooLarge 上传的文件不符合大小限制
	FileSizeTooLarge

	// FileTypeNoMatch 上传文件的类型不匹配
	FileTypeNoMatch

	// VerifyCodeError 验证码错误
	VerifyCodeError

	// EmailIsDifferent 两次邮箱不一致
	EmailIsDifferent

	// MenuExists  菜单已经存在
	MenuExists

	// MenuNotExists  菜单不存在
	MenuNotExists

	// TagExists   标签已经存在
	TagExists

	// TagNotExists 标签不存在
	TagNotExists

	// ImageExists   图片已经存在
	ImageExists

	// ImageNotExists 图片不存在
	ImageNotExists

	// OtherError 其他错误
	OtherError
)
