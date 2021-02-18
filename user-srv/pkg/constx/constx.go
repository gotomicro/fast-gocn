package constx

const (
	UserOpenWechatWeb  = 1
	UserOpenWechatH5   = 2
	UserOpenWechatApp  = 3
	UserOpenWechatMini = 4
)

const (
	DateStandard    = "2006-01-02 15:04:05"
	DateDocIdentify = "2006-01-02-15-04-05"
)

const (
	PhoneCodePrefix  = "gocn:user:phone_%d"
	PhoneCodeLength  = 4
	PhoneCodeExpired = 300
	EmailCodePrefix  = "gocn:user:email_%d"
	EmailCodeLength  = 4
	EmailCodeExpired = 300
)

const (
	VIP_BASE   = 1
	VIP_ANNUAL = 2 // 大会会员
)

const (
	OpTypeProgram int32 = 1
	OpTypeAdmin   int32 = 2
)

const (
	DefaultMembershipConfigKey = "vip.membership"
)

const (
	WechatQrCode             = "https://wxopen.gocn.vip/qrcode?invitationCode=%d_%s"
	ColumnQrCode             = "https://wxopen.gocn.vip/pages/column/columnBook?id=%d"

	InvitationTypeNormal = 0
	InvitationTypeActivity = 1
	InvitationTypeColumn = 2
)
