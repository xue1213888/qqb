package dto

type QQFriendAddData struct {
	Openid    string `json:"openid"`
	Timestamp int64  `json:"timestamp"`
}

type QQC2cMsgRejectData struct {
	Openid    string `json:"openid"`
	Timestamp int64  `json:"timestamp"`
}

type QQC2cMsgReceiveData struct {
	Openid    string `json:"openid"`
	Timestamp int64  `json:"timestamp"`
}

type QQGroupAddRobotData struct {
	GroupOpenid    string `json:"group_openid"`
	OpMemberOpenid string `json:"op_member_openid"`
	Timestamp      int64  `json:"timestamp"`
}

type QQGroupDelRobotData struct {
	GroupOpenid    string `json:"group_openid"`
	OpMemberOpenid string `json:"op_member_openid"`
	Timestamp      int64  `json:"timestamp"`
}

type QQGroupMsgRejectData struct {
	GroupOpenid    string `json:"group_openid"`
	OpMemberOpenid string `json:"op_member_openid"`
	Timestamp      int64  `json:"timestamp"`
}

type QQGroupMsgReceiveData struct {
	GroupOpenid    string `json:"group_openid"`
	OpMemberOpenid string `json:"op_member_openid"`
	Timestamp      int64  `json:"timestamp"`
}

type QQAuthor struct {
	UserOpenid string `json:"user_openid,omitempty"`
}

type QQGroupAuthor struct {
	MemberOpenid string `json:"member_openid,omitempty"`
}

type QQC2cMessageCreateData struct {
	ID        string    `json:"id"`
	Author    QQAuthor  `json:"author"`
	Content   string    `json:"content"`
	Timestamp Timestamp `json:"timestamp"`
}

type QQGroupAtMessageCreateData struct {
	ID          string        `json:"id"`
	Author      QQGroupAuthor `json:"author"`
	GroupOpenid string        `json:"group_openid"`
	Content     string        `json:"content"`
	Timestamp   Timestamp     `json:"timestamp"`
}
