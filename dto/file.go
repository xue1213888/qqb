package dto

// FileToCreate 上传文件的接口
type FileToCreate struct {
	FileType   int    `json:"file_type,omitempty"`    // 文件类型，1 图片，2 视频，3 音频，4 文件
	Url        string `json:"url,omitempty"`          // 文件的 url
	SrvSendMsg bool   `json:"srv_send_msg,omitempty"` // 是否发送消息
}

type FileResponse struct {
	FileUUID string `json:"file_uuid,omitempty"` // 文件的 uuid
	FileInfo string `json:"file_info,omitempty"` // 文件的信息
	Ttl      int    `json:"ttl,omitempty"`       // 文件的有效期
}
