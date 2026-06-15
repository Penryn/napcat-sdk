// 代码由 napcatgen 根据 NapCat OpenAPI 4.18.6 生成；请勿手动修改。

package api

// DotHandleQuickOperationRequest 处理快速操作 请求参数。
type DotHandleQuickOperationRequest struct {
	// Context 事件上下文
	Context map[string]any `json:"context"`
	// Operation 快速操作内容
	Operation map[string]any `json:"operation"`
}

// DotOcrImageRequest 图片 OCR 识别 (内部) 请求参数。
type DotOcrImageRequest struct {
	// Image 图片路径、URL或Base64
	Image string `json:"image"`
}

// ArkShareGroupRequest 分享群 (Ark) 请求参数。
type ArkShareGroupRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// ArkSharePeerRequest 分享用户 (Ark) 请求参数。
type ArkSharePeerRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// PhoneNumber 手机号
	PhoneNumber string `json:"phone_number"`
	// UserID QQ号
	UserID string `json:"user_id,omitempty"`
}

// UnderscoreDelGroupNoticeRequest 删除群公告 请求参数。
type UnderscoreDelGroupNoticeRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// NoticeID 公告ID
	NoticeID string `json:"notice_id"`
}

// UnderscoreGetGroupNoticeRequest 获取群公告 请求参数。
type UnderscoreGetGroupNoticeRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// UnderscoreGetModelShowRequest 获取机型显示 请求参数。
type UnderscoreGetModelShowRequest struct {
	// Model 模型名称
	Model string `json:"model,omitempty"`
}

// UnderscoreMarkAllAsReadRequest 标记所有消息已读 请求参数。
type UnderscoreMarkAllAsReadRequest struct {
}

// UnderscoreSendGroupNoticeRequest 发送群公告 请求参数。
type UnderscoreSendGroupNoticeRequest struct {
	// ConfirmRequired 是否需要确认 (0/1)
	ConfirmRequired any `json:"confirm_required"`
	// Content 公告内容
	Content string `json:"content"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Image 公告图片路径或 URL
	Image string `json:"image,omitempty"`
	// IsShowEditCard 是否显示修改群名片引导 (0/1)
	IsShowEditCard any `json:"is_show_edit_card"`
	// Pinned 是否置顶 (0/1)
	Pinned any `json:"pinned"`
	// TipWindowType 弹窗类型 (默认为 0)
	TipWindowType any `json:"tip_window_type"`
	// Type 类型 (默认为 1)
	Type any `json:"type"`
}

// UnderscoreSetModelShowRequest 设置机型 请求参数。
type UnderscoreSetModelShowRequest struct {
}

// AddCustomFaceRequest 添加自定义表情 请求参数。
type AddCustomFaceRequest struct {
	// EmojiID 表情ID，未提供时传空字符串
	EmojiID any `json:"emoji_id,omitempty"`
	// File 本地表情文件路径
	File string `json:"file"`
	// FileName 文件名，未提供时从file路径取basename
	FileName string `json:"file_name,omitempty"`
	// FileSize 文件大小，未提供时读取本地文件
	FileSize any `json:"file_size,omitempty"`
	// IsMarkFace 是否商城表情
	IsMarkFace bool `json:"is_mark_face,omitempty"`
	// IsOrigin 是否原图
	IsOrigin bool `json:"is_origin,omitempty"`
	// Md5 文件MD5，未提供时读取本地文件计算
	Md5 string `json:"md5,omitempty"`
	// PackageID 表情包ID，未提供时传0
	PackageID any `json:"package_id,omitempty"`
}

// BotExitRequest 退出登录 请求参数。
type BotExitRequest struct {
}

// CanSendImageRequest 是否可以发送图片 请求参数。
type CanSendImageRequest struct {
}

// CanSendRecordRequest 是否可以发送语音 请求参数。
type CanSendRecordRequest struct {
}

// CancelGroupAlbumMediaLikeRequest 取消点赞群相册媒体 请求参数。
type CancelGroupAlbumMediaLikeRequest struct {
	// AlbumID 相册ID
	AlbumID string `json:"album_id"`
	// BatchID batch_id
	BatchID string `json:"batch_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Lloc lloc，若对整个上传操作则不填
	Lloc string `json:"lloc,omitempty"`
}

// CancelGroupTodoRequest 取消群待办 请求参数。
type CancelGroupTodoRequest struct {
	// GroupID 群号
	GroupID any `json:"group_id"`
	// MessageID 消息ID
	MessageID string `json:"message_id,omitempty"`
	// MessageSeq 消息Seq (可选)
	MessageSeq string `json:"message_seq,omitempty"`
}

// CancelOnlineFileRequest 取消在线文件 请求参数。
type CancelOnlineFileRequest struct {
	// MsgID 消息 ID
	MsgID string `json:"msg_id"`
	// UserID 用户 QQ
	UserID string `json:"user_id"`
}

// CheckURLSafelyRequest 检查URL安全性 请求参数。
type CheckURLSafelyRequest struct {
	// URL 要检查的 URL
	URL string `json:"url"`
}

// CleanCacheRequest 清理缓存 请求参数。
type CleanCacheRequest struct {
}

// CleanStreamTempFileRequest 清理流式传输临时文件 请求参数。
type CleanStreamTempFileRequest struct {
}

// ClickInlineKeyboardButtonRequest 点击内联键盘按钮 请求参数。
type ClickInlineKeyboardButtonRequest struct {
	// BotAppid 机器人AppID
	BotAppid string `json:"bot_appid"`
	// ButtonID 按钮ID
	ButtonID string `json:"button_id"`
	// CallbackData 回调数据
	CallbackData string `json:"callback_data"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// MsgSeq 消息序列号
	MsgSeq string `json:"msg_seq"`
}

// CompleteGroupTodoRequest 完成群待办 请求参数。
type CompleteGroupTodoRequest struct {
	// GroupID 群号
	GroupID any `json:"group_id"`
	// MessageID 消息ID
	MessageID string `json:"message_id,omitempty"`
	// MessageSeq 消息Seq (可选)
	MessageSeq string `json:"message_seq,omitempty"`
}

// CreateCollectionRequest 创建收藏 请求参数。
type CreateCollectionRequest struct {
	// Brief 简要描述
	Brief string `json:"brief"`
	// RawData 原始数据
	RawData string `json:"rawData"`
}

// CreateFlashTaskRequest 创建闪传任务 请求参数。
type CreateFlashTaskRequest struct {
	// Files 文件列表或单个文件路径
	Files any `json:"files"`
	// Name 任务名称
	Name string `json:"name,omitempty"`
	// ThumbPath 缩略图路径
	ThumbPath string `json:"thumb_path,omitempty"`
}

// CreateGroupFileFolderRequest 创建群文件目录 请求参数。
type CreateGroupFileFolderRequest struct {
	// FolderName 文件夹名称
	FolderName string `json:"folder_name,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Name 文件夹名称
	Name string `json:"name,omitempty"`
}

// DelGroupAlbumMediaRequest 删除群相册媒体 请求参数。
type DelGroupAlbumMediaRequest struct {
	// AlbumID 相册ID
	AlbumID string `json:"album_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Lloc 媒体ID (lloc)
	Lloc string `json:"lloc"`
}

// DeleteCustomFaceRequest 删除自定义表情 请求参数。
type DeleteCustomFaceRequest struct {
	ID    any      `json:"id,omitempty"`
	Ids   []string `json:"ids,omitempty"`
	Md5   any      `json:"md5,omitempty"`
	ResID any      `json:"res_id,omitempty"`
}

// DeleteEssenceMsgRequest 移出精华消息 请求参数。
type DeleteEssenceMsgRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// MessageID 消息ID
	MessageID any `json:"message_id,omitempty"`
	// MsgRandom 消息随机数
	MsgRandom string `json:"msg_random,omitempty"`
	// MsgSeq 消息序号
	MsgSeq string `json:"msg_seq,omitempty"`
}

// DeleteFriendRequest 删除好友 请求参数。
type DeleteFriendRequest struct {
	// FriendID 好友 QQ 号
	FriendID any `json:"friend_id,omitempty"`
	// TempBlock 是否加入黑名单
	TempBlock bool `json:"temp_block,omitempty"`
	// TempBothDel 是否双向删除
	TempBothDel bool `json:"temp_both_del,omitempty"`
	// UserID 用户 QQ 号
	UserID any `json:"user_id,omitempty"`
}

// DeleteGroupFileRequest 删除群文件 请求参数。
type DeleteGroupFileRequest struct {
	// FileID 文件ID
	FileID string `json:"file_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// DeleteGroupFolderRequest 删除群文件目录 请求参数。
type DeleteGroupFolderRequest struct {
	// Folder 文件夹ID
	Folder string `json:"folder,omitempty"`
	// FolderID 文件夹ID
	FolderID string `json:"folder_id,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// DeleteMsgRequest 撤回消息 请求参数。
type DeleteMsgRequest struct {
	// MessageID 消息ID
	MessageID any `json:"message_id"`
}

// DoGroupAlbumCommentRequest 发表群相册评论 请求参数。
type DoGroupAlbumCommentRequest struct {
	// AlbumID 相册 ID
	AlbumID string `json:"album_id"`
	// Content 评论内容
	Content string `json:"content"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Lloc 图片 ID
	Lloc string `json:"lloc"`
}

// DownloadFileRequest 下载文件 请求参数。
type DownloadFileRequest struct {
	// Base64 base64数据
	Base64 string `json:"base64,omitempty"`
	// Headers 请求头
	Headers any `json:"headers,omitempty"`
	// Name 文件名
	Name string `json:"name,omitempty"`
	// URL 下载链接
	URL string `json:"url,omitempty"`
}

// DownloadFileImageStreamRequest 下载图片文件流 请求参数。
type DownloadFileImageStreamRequest struct {
	// ChunkSize 分块大小 (字节)
	ChunkSize int64 `json:"chunk_size,omitempty"`
	// File 文件路径或 URL
	File string `json:"file,omitempty"`
	// FileID 文件 ID
	FileID string `json:"file_id,omitempty"`
}

// DownloadFileRecordStreamRequest 下载语音文件流 请求参数。
type DownloadFileRecordStreamRequest struct {
	// ChunkSize 分块大小 (字节)
	ChunkSize int64 `json:"chunk_size,omitempty"`
	// File 文件路径或 URL
	File string `json:"file,omitempty"`
	// FileID 文件 ID
	FileID string `json:"file_id,omitempty"`
	// OutFormat 输出格式
	OutFormat string `json:"out_format,omitempty"`
}

// DownloadFileStreamRequest 下载文件流 请求参数。
type DownloadFileStreamRequest struct {
	// ChunkSize 分块大小 (字节)
	ChunkSize int64 `json:"chunk_size,omitempty"`
	// File 文件路径或 URL
	File string `json:"file,omitempty"`
	// FileID 文件 ID
	FileID string `json:"file_id,omitempty"`
}

// DownloadFilesetRequest 下载文件集 请求参数。
type DownloadFilesetRequest struct {
	// FilesetID 文件集 ID
	FilesetID string `json:"fileset_id"`
}

// FetchCustomFaceRequest 获取自定义表情 请求参数。
type FetchCustomFaceRequest struct {
	// Count 获取数量
	Count any `json:"count"`
}

// FetchCustomFaceDetailRequest 获取自定义表情详情 请求参数。
type FetchCustomFaceDetailRequest struct {
	// Count 获取数量
	Count any `json:"count"`
}

// FetchEmojiLikeRequest 获取表情点赞详情 请求参数。
type FetchEmojiLikeRequest struct {
	// Cookie 分页Cookie
	Cookie string `json:"cookie"`
	// Count 获取数量
	Count any `json:"count"`
	// EmojiId 表情ID
	EmojiId any `json:"emojiId"`
	// EmojiType 表情类型
	EmojiType any `json:"emojiType"`
	// MessageID 消息ID
	MessageID any `json:"message_id"`
}

// FetchPttTextRequest 获取语音转文字结果 请求参数。
type FetchPttTextRequest struct {
	// MessageID 消息ID
	MessageID any `json:"message_id"`
}

// ForwardFriendSingleMsgRequest 转发单条消息 请求参数。
type ForwardFriendSingleMsgRequest struct {
	// GroupID 目标群号
	GroupID string `json:"group_id,omitempty"`
	// MessageID 消息ID
	MessageID any `json:"message_id"`
	// UserID 目标用户QQ
	UserID string `json:"user_id,omitempty"`
}

// ForwardGroupSingleMsgRequest 转发单条消息 请求参数。
type ForwardGroupSingleMsgRequest struct {
	// GroupID 目标群号
	GroupID string `json:"group_id,omitempty"`
	// MessageID 消息ID
	MessageID any `json:"message_id"`
	// UserID 目标用户QQ
	UserID string `json:"user_id,omitempty"`
}

// FriendPokeRequest 发送戳一戳 请求参数。
type FriendPokeRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// TargetID 目标QQ
	TargetID string `json:"target_id,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// GetAiCharactersRequest 获取AI角色列表 请求参数。
type GetAiCharactersRequest struct {
	// ChatType 聊天类型
	ChatType any `json:"chat_type"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetAiRecordRequest 获取 AI 语音 请求参数。
type GetAiRecordRequest struct {
	// Character 角色ID
	Character string `json:"character"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Text 语音文本内容
	Text string `json:"text"`
}

// GetClientkeyRequest 获取ClientKey 请求参数。
type GetClientkeyRequest struct {
}

// GetCollectionListRequest 获取收藏列表 请求参数。
type GetCollectionListRequest struct {
	// Category 分类ID
	Category string `json:"category"`
	// Count 获取数量
	Count string `json:"count"`
}

// GetCookiesRequest 获取 Cookies 请求参数。
type GetCookiesRequest struct {
	// Domain 需要获取 cookies 的域名
	Domain string `json:"domain"`
}

// GetCredentialsRequest 获取登录凭证 请求参数。
type GetCredentialsRequest struct {
	// Domain 需要获取 cookies 的域名
	Domain string `json:"domain"`
}

// GetCsrfTokenRequest 获取 CSRF Token 请求参数。
type GetCsrfTokenRequest struct {
}

// GetDoubtFriendsAddRequestRequest 获取可疑好友申请 请求参数。
type GetDoubtFriendsAddRequestRequest struct {
	// Count 获取数量
	Count int64 `json:"count"`
}

// GetEmojiLikesRequest 获取消息表情点赞列表 请求参数。
type GetEmojiLikesRequest struct {
	// Count 数量，0代表全部
	Count int64 `json:"count"`
	// EmojiID 表情ID
	EmojiID string `json:"emoji_id"`
	// EmojiType 表情类型
	EmojiType string `json:"emoji_type,omitempty"`
	// GroupID 群号，短ID可不传
	GroupID string `json:"group_id,omitempty"`
	// MessageID 消息ID，可以传递长ID或短ID
	MessageID string `json:"message_id"`
}

// GetEssenceMsgListRequest 获取群精华消息 请求参数。
type GetEssenceMsgListRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetFileRequest 获取文件 请求参数。
type GetFileRequest struct {
	// File 文件路径、URL或Base64
	File string `json:"file,omitempty"`
	// FileID 文件ID
	FileID string `json:"file_id,omitempty"`
}

// GetFilesetIDRequest 获取文件集 ID 请求参数。
type GetFilesetIDRequest struct {
	// ShareCode 分享码或分享链接
	ShareCode string `json:"share_code"`
}

// GetFilesetInfoRequest 获取文件集信息 请求参数。
type GetFilesetInfoRequest struct {
	// FilesetID 文件集 ID
	FilesetID string `json:"fileset_id"`
}

// GetFlashFileListRequest 获取闪传文件列表 请求参数。
type GetFlashFileListRequest struct {
	// FilesetID 文件集 ID
	FilesetID string `json:"fileset_id"`
}

// GetFlashFileURLRequest 获取闪传文件链接 请求参数。
type GetFlashFileURLRequest struct {
	// FileIndex 文件索引
	FileIndex int64 `json:"file_index,omitempty"`
	// FileName 文件名
	FileName string `json:"file_name,omitempty"`
	// FilesetID 文件集 ID
	FilesetID string `json:"fileset_id"`
}

// GetForwardMsgRequest 获取合并转发消息 请求参数。
type GetForwardMsgRequest struct {
	// ID 消息ID
	ID string `json:"id,omitempty"`
	// MessageID 消息ID
	MessageID string `json:"message_id,omitempty"`
}

// GetFriendListRequest 获取好友列表 请求参数。
type GetFriendListRequest struct {
	// NoCache 是否不使用缓存
	NoCache any `json:"no_cache,omitempty"`
}

// GetFriendMsgHistoryRequest 获取好友历史消息 请求参数。
type GetFriendMsgHistoryRequest struct {
	// Count 获取消息数量
	Count int64 `json:"count"`
	// DisableGetURL 是否禁用获取URL
	DisableGetURL bool `json:"disable_get_url"`
	// MessageSeq 起始消息序号
	MessageSeq string `json:"message_seq,omitempty"`
	// ParseMultMsg 是否解析合并消息
	ParseMultMsg bool `json:"parse_mult_msg"`
	// QuickReply 是否快速回复
	QuickReply bool `json:"quick_reply"`
	// ReverseOrder 是否反向排序(旧版本兼容)
	ReverseOrder bool `json:"reverseOrder"`
	// ReverseOrder2 是否反向排序
	ReverseOrder2 bool `json:"reverse_order"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// GetFriendsWithCategoryRequest 获取带分组的好友列表 请求参数。
type GetFriendsWithCategoryRequest struct {
}

// GetGroupAlbumMediaListRequest 获取群相册媒体列表 请求参数。
type GetGroupAlbumMediaListRequest struct {
	// AlbumID 相册ID
	AlbumID string `json:"album_id"`
	// AttachInfo 附加信息（用于分页）
	AttachInfo string `json:"attach_info,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupAtAllRemainRequest 获取群艾特全体剩余次数 请求参数。
type GetGroupAtAllRemainRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupDetailInfoRequest 获取群详细信息 请求参数。
type GetGroupDetailInfoRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupFileSystemInfoRequest 获取群文件系统信息 请求参数。
type GetGroupFileSystemInfoRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupFileURLRequest 获取群文件URL 请求参数。
type GetGroupFileURLRequest struct {
	// FileID 文件ID
	FileID string `json:"file_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupFilesByFolderRequest 获取群文件夹文件列表 请求参数。
type GetGroupFilesByFolderRequest struct {
	// FileCount 文件数量
	FileCount any `json:"file_count"`
	// Folder 文件夹ID
	Folder string `json:"folder,omitempty"`
	// FolderID 文件夹ID
	FolderID string `json:"folder_id,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupHonorInfoRequest 获取群荣誉信息 请求参数。
type GetGroupHonorInfoRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Type 荣誉类型
	Type string `json:"type,omitempty"`
}

// GetGroupIgnoreAddRequestRequest 获取群被忽略的加群请求 请求参数。
type GetGroupIgnoreAddRequestRequest struct {
}

// GetGroupIgnoredNotifiesRequest 获取群忽略通知 请求参数。
type GetGroupIgnoredNotifiesRequest struct {
}

// GetGroupInfoRequest 获取群信息 请求参数。
type GetGroupInfoRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupInfoExRequest 获取群详细信息 (扩展) 请求参数。
type GetGroupInfoExRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupListRequest 获取群列表 请求参数。
type GetGroupListRequest struct {
	// NoCache 是否不使用缓存
	NoCache any `json:"no_cache,omitempty"`
}

// GetGroupMemberInfoRequest 获取群成员信息 请求参数。
type GetGroupMemberInfoRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// NoCache 是否不使用缓存
	NoCache any `json:"no_cache,omitempty"`
	// UserID QQ号
	UserID string `json:"user_id"`
}

// GetGroupMemberListRequest 获取群成员列表 请求参数。
type GetGroupMemberListRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// NoCache 是否不使用缓存
	NoCache any `json:"no_cache,omitempty"`
}

// GetGroupMsgHistoryRequest 获取群历史消息 请求参数。
type GetGroupMsgHistoryRequest struct {
	// Count 获取消息数量
	Count int64 `json:"count"`
	// DisableGetURL 是否禁用获取URL
	DisableGetURL bool `json:"disable_get_url"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// MessageSeq 起始消息序号
	MessageSeq string `json:"message_seq,omitempty"`
	// ParseMultMsg 是否解析合并消息
	ParseMultMsg bool `json:"parse_mult_msg"`
	// QuickReply 是否快速回复
	QuickReply bool `json:"quick_reply"`
	// ReverseOrder 是否反向排序(旧版本兼容)
	ReverseOrder bool `json:"reverseOrder"`
	// ReverseOrder2 是否反向排序
	ReverseOrder2 bool `json:"reverse_order"`
}

// GetGroupRootFilesRequest 获取群根目录文件列表 请求参数。
type GetGroupRootFilesRequest struct {
	// FileCount 文件数量
	FileCount any `json:"file_count"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupShutListRequest 获取群禁言列表 请求参数。
type GetGroupShutListRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetGroupSignedListRequest 获取群组今日打卡列表 请求参数。
type GetGroupSignedListRequest struct {
	// GroupID 群号
	GroupID any `json:"group_id"`
}

// GetGroupSystemMsgRequest 获取群系统消息 请求参数。
type GetGroupSystemMsgRequest struct {
	// Count 获取的消息数量
	Count any `json:"count"`
}

// GetGuildListRequest 获取频道列表 请求参数。
type GetGuildListRequest struct {
}

// GetGuildServiceProfileRequest 获取频道个人信息 请求参数。
type GetGuildServiceProfileRequest struct {
}

// GetImageRequest 获取图片 请求参数。
type GetImageRequest struct {
	// File 文件路径、URL或Base64
	File string `json:"file,omitempty"`
	// FileID 文件ID
	FileID string `json:"file_id,omitempty"`
}

// GetLoginInfoRequest 获取登录号信息 请求参数。
type GetLoginInfoRequest struct {
}

// GetMiniAppArkRequest 获取小程序 Ark 请求参数。
type GetMiniAppArkRequest struct {
}

// GetMsgRequest 获取消息 请求参数。
type GetMsgRequest struct {
	// MessageID 消息ID
	MessageID any `json:"message_id"`
}

// GetOnlineClientsRequest 获取在线客户端 请求参数。
type GetOnlineClientsRequest struct {
}

// GetOnlineFileMsgRequest 获取在线文件消息 请求参数。
type GetOnlineFileMsgRequest struct {
	// UserID 用户 QQ
	UserID string `json:"user_id"`
}

// GetPrivateFileURLRequest 获取私聊文件URL 请求参数。
type GetPrivateFileURLRequest struct {
	// FileID 文件ID
	FileID string `json:"file_id"`
}

// GetProfileLikeRequest 获取资料点赞 请求参数。
type GetProfileLikeRequest struct {
	// Count 获取数量
	Count any `json:"count"`
	// Start 起始位置
	Start any `json:"start"`
	// UserID QQ号
	UserID string `json:"user_id,omitempty"`
}

// GetQunAlbumListRequest 获取群相册列表 请求参数。
type GetQunAlbumListRequest struct {
	// AttachInfo 附加信息（用于分页，从上一次返回结果中获取）
	AttachInfo string `json:"attach_info,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// GetRecentContactRequest 获取最近会话 请求参数。
type GetRecentContactRequest struct {
	// Count 获取的数量
	Count any `json:"count"`
}

// GetRecordRequest 获取语音 请求参数。
type GetRecordRequest struct {
	// File 文件路径、URL或Base64
	File string `json:"file,omitempty"`
	// FileID 文件ID
	FileID string `json:"file_id,omitempty"`
	// OutFormat 输出格式
	OutFormat string `json:"out_format"`
}

// GetRkeyRequest 获取扩展 RKey 请求参数。
type GetRkeyRequest struct {
}

// GetRkeyServerRequest 获取 RKey 服务器 请求参数。
type GetRkeyServerRequest struct {
}

// GetRobotUINRangeRequest 获取机器人 UIN 范围 请求参数。
type GetRobotUINRangeRequest struct {
}

// GetShareLinkRequest 获取文件分享链接 请求参数。
type GetShareLinkRequest struct {
	// FilesetID 文件集 ID
	FilesetID string `json:"fileset_id"`
}

// GetStatusRequest 获取运行状态 请求参数。
type GetStatusRequest struct {
}

// GetStrangerInfoRequest 获取陌生人信息 请求参数。
type GetStrangerInfoRequest struct {
	// NoCache 是否不使用缓存
	NoCache any `json:"no_cache"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// GetUnidirectionalFriendListRequest 获取单向好友列表 请求参数。
type GetUnidirectionalFriendListRequest struct {
}

// GetVersionInfoRequest 获取版本信息 请求参数。
type GetVersionInfoRequest struct {
}

// GroupPokeRequest 发送戳一戳 请求参数。
type GroupPokeRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// TargetID 目标QQ
	TargetID string `json:"target_id,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// MarkGroupMsgAsReadRequest 标记群聊已读 请求参数。
type MarkGroupMsgAsReadRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// MessageID 消息ID
	MessageID string `json:"message_id,omitempty"`
	// UserID 用户QQ
	UserID any `json:"user_id,omitempty"`
}

// MarkMsgAsReadRequest 标记消息已读 (Go-CQHTTP) 请求参数。
type MarkMsgAsReadRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// MessageID 消息ID
	MessageID string `json:"message_id,omitempty"`
	// UserID 用户QQ
	UserID any `json:"user_id,omitempty"`
}

// MarkPrivateMsgAsReadRequest 标记私聊已读 请求参数。
type MarkPrivateMsgAsReadRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// MessageID 消息ID
	MessageID string `json:"message_id,omitempty"`
	// UserID 用户QQ
	UserID any `json:"user_id,omitempty"`
}

// MoveGroupFileRequest 移动群文件 请求参数。
type MoveGroupFileRequest struct {
	// CurrentParentDirectory 当前父目录
	CurrentParentDirectory string `json:"current_parent_directory"`
	// FileID 文件ID
	FileID string `json:"file_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// TargetParentDirectory 目标父目录
	TargetParentDirectory string `json:"target_parent_directory"`
}

// NcGetPacketStatusRequest 获取Packet状态 请求参数。
type NcGetPacketStatusRequest struct {
}

// NcGetRkeyRequest 获取 RKey 请求参数。
type NcGetRkeyRequest struct {
}

// NcGetUserStatusRequest 获取用户在线状态 请求参数。
type NcGetUserStatusRequest struct {
	// UserID QQ号
	UserID string `json:"user_id"`
}

// OcrImageRequest 图片 OCR 识别 请求参数。
type OcrImageRequest struct {
	// Image 图片路径、URL或Base64
	Image string `json:"image"`
}

// ReceiveOnlineFileRequest 接收在线文件 请求参数。
type ReceiveOnlineFileRequest struct {
	// ElementID 元素 ID
	ElementID string `json:"element_id"`
	// MsgID 消息 ID
	MsgID string `json:"msg_id"`
	// UserID 用户 QQ
	UserID string `json:"user_id"`
}

// RefuseOnlineFileRequest 拒绝在线文件 请求参数。
type RefuseOnlineFileRequest struct {
	// ElementID 元素 ID
	ElementID string `json:"element_id"`
	// MsgID 消息 ID
	MsgID string `json:"msg_id"`
	// UserID 用户 QQ
	UserID string `json:"user_id"`
}

// RenameGroupFileRequest 重命名群文件 请求参数。
type RenameGroupFileRequest struct {
	// CurrentParentDirectory 当前父目录
	CurrentParentDirectory string `json:"current_parent_directory"`
	// FileID 文件ID
	FileID string `json:"file_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// NewName 新文件名
	NewName string `json:"new_name"`
}

// SendArkShareRequest 分享用户 (Ark) 请求参数。
type SendArkShareRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// PhoneNumber 手机号
	PhoneNumber string `json:"phone_number"`
	// UserID QQ号
	UserID string `json:"user_id,omitempty"`
}

// SendFlashMsgRequest 发送闪传消息 请求参数。
type SendFlashMsgRequest struct {
	// FilesetID 文件集 ID
	FilesetID string `json:"fileset_id"`
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// UserID 用户 QQ
	UserID string `json:"user_id,omitempty"`
}

// SendForwardMsgRequest 发送合并转发消息 请求参数。
type SendForwardMsgRequest struct {
	// AutoEscape 是否作为纯文本发送
	AutoEscape any `json:"auto_escape,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	Message any    `json:"message"`
	// MessageType 消息类型 (private/group)
	MessageType string `json:"message_type,omitempty"`
	// News 合并转发新闻
	News []map[string]any `json:"news,omitempty"`
	// Prompt 合并转发提示
	Prompt string `json:"prompt,omitempty"`
	// Source 合并转发来源
	Source string `json:"source,omitempty"`
	// Summary 合并转发摘要
	Summary string `json:"summary,omitempty"`
	// Timeout 自定义发送超时(毫秒)，覆盖自动计算值
	Timeout int64 `json:"timeout,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id,omitempty"`
}

// SendGroupAiRecordRequest 发送群 AI 语音 请求参数。
type SendGroupAiRecordRequest struct {
	// Character 角色ID
	Character string `json:"character"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Text 语音文本内容
	Text string `json:"text"`
}

// SendGroupArkShareRequest 分享群 (Ark) 请求参数。
type SendGroupArkShareRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// SendGroupForwardMsgRequest 发送群合并转发消息 请求参数。
type SendGroupForwardMsgRequest struct {
	// AutoEscape 是否作为纯文本发送
	AutoEscape any `json:"auto_escape,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	Message any    `json:"message"`
	// MessageType 消息类型 (private/group)
	MessageType string `json:"message_type,omitempty"`
	// News 合并转发新闻
	News []map[string]any `json:"news,omitempty"`
	// Prompt 合并转发提示
	Prompt string `json:"prompt,omitempty"`
	// Source 合并转发来源
	Source string `json:"source,omitempty"`
	// Summary 合并转发摘要
	Summary string `json:"summary,omitempty"`
	// Timeout 自定义发送超时(毫秒)，覆盖自动计算值
	Timeout int64 `json:"timeout,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id,omitempty"`
}

// SendGroupMsgRequest 发送群消息 请求参数。
type SendGroupMsgRequest struct {
	// AutoEscape 是否作为纯文本发送
	AutoEscape any `json:"auto_escape,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	Message any    `json:"message"`
	// MessageType 消息类型 (private/group)
	MessageType string `json:"message_type,omitempty"`
	// News 合并转发新闻
	News []map[string]any `json:"news,omitempty"`
	// Prompt 合并转发提示
	Prompt string `json:"prompt,omitempty"`
	// Source 合并转发来源
	Source string `json:"source,omitempty"`
	// Summary 合并转发摘要
	Summary string `json:"summary,omitempty"`
	// Timeout 自定义发送超时(毫秒)，覆盖自动计算值
	Timeout int64 `json:"timeout,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id,omitempty"`
}

// SendGroupSignRequest 群打卡 请求参数。
type SendGroupSignRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// SendLikeRequest 点赞 请求参数。
type SendLikeRequest struct {
	// Times 点赞次数
	Times any `json:"times"`
	// UserID 对方 QQ 号
	UserID string `json:"user_id"`
}

// SendMsgRequest 发送消息 请求参数。
type SendMsgRequest struct {
	// AutoEscape 是否作为纯文本发送
	AutoEscape any `json:"auto_escape,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	Message any    `json:"message"`
	// MessageType 消息类型 (private/group)
	MessageType string `json:"message_type,omitempty"`
	// News 合并转发新闻
	News []map[string]any `json:"news,omitempty"`
	// Prompt 合并转发提示
	Prompt string `json:"prompt,omitempty"`
	// Source 合并转发来源
	Source string `json:"source,omitempty"`
	// Summary 合并转发摘要
	Summary string `json:"summary,omitempty"`
	// Timeout 自定义发送超时(毫秒)，覆盖自动计算值
	Timeout int64 `json:"timeout,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id,omitempty"`
}

// SendOnlineFileRequest 发送在线文件 请求参数。
type SendOnlineFileRequest struct {
	// FileName 文件名 (可选)
	FileName string `json:"file_name,omitempty"`
	// FilePath 本地文件路径
	FilePath string `json:"file_path"`
	// UserID 用户 QQ
	UserID string `json:"user_id"`
}

// SendOnlineFolderRequest 发送在线文件夹 请求参数。
type SendOnlineFolderRequest struct {
	// FolderName 文件夹名称 (可选)
	FolderName string `json:"folder_name,omitempty"`
	// FolderPath 本地文件夹路径
	FolderPath string `json:"folder_path"`
	// UserID 用户 QQ
	UserID string `json:"user_id"`
}

// SendPacketRequest 发送原始数据包 请求参数。
type SendPacketRequest struct {
	// Cmd 命令字
	Cmd string `json:"cmd"`
	// Data 十六进制数据
	Data string `json:"data"`
	// Rsp 是否等待响应
	Rsp any `json:"rsp"`
}

// SendPokeRequest 发送戳一戳 请求参数。
type SendPokeRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	// TargetID 目标QQ
	TargetID string `json:"target_id,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// SendPrivateForwardMsgRequest 发送私聊合并转发消息 请求参数。
type SendPrivateForwardMsgRequest struct {
	// AutoEscape 是否作为纯文本发送
	AutoEscape any `json:"auto_escape,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	Message any    `json:"message"`
	// MessageType 消息类型 (private/group)
	MessageType string `json:"message_type,omitempty"`
	// News 合并转发新闻
	News []map[string]any `json:"news,omitempty"`
	// Prompt 合并转发提示
	Prompt string `json:"prompt,omitempty"`
	// Source 合并转发来源
	Source string `json:"source,omitempty"`
	// Summary 合并转发摘要
	Summary string `json:"summary,omitempty"`
	// Timeout 自定义发送超时(毫秒)，覆盖自动计算值
	Timeout int64 `json:"timeout,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id,omitempty"`
}

// SendPrivateMsgRequest 发送私聊消息 请求参数。
type SendPrivateMsgRequest struct {
	// AutoEscape 是否作为纯文本发送
	AutoEscape any `json:"auto_escape,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id,omitempty"`
	Message any    `json:"message"`
	// MessageType 消息类型 (private/group)
	MessageType string `json:"message_type,omitempty"`
	// News 合并转发新闻
	News []map[string]any `json:"news,omitempty"`
	// Prompt 合并转发提示
	Prompt string `json:"prompt,omitempty"`
	// Source 合并转发来源
	Source string `json:"source,omitempty"`
	// Summary 合并转发摘要
	Summary string `json:"summary,omitempty"`
	// Timeout 自定义发送超时(毫秒)，覆盖自动计算值
	Timeout int64 `json:"timeout,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id,omitempty"`
}

// SetCustomFaceDescRequest 修改自定义表情描述 请求参数。
type SetCustomFaceDescRequest struct {
	// Desc 新的表情描述
	Desc string `json:"desc"`
	// EmojiID 表情ID
	EmojiID any `json:"emoji_id"`
	// Md5 表情MD5
	Md5 string `json:"md5"`
	// ResID 资源ID
	ResID string `json:"res_id"`
}

// SetDiyOnlineStatusRequest 设置自定义在线状态 请求参数。
type SetDiyOnlineStatusRequest struct {
	// FaceID 图标ID
	FaceID any `json:"face_id"`
	// FaceType 图标类型
	FaceType any `json:"face_type"`
	// Wording 状态文字内容
	Wording string `json:"wording"`
}

// SetDoubtFriendsAddRequestRequest 处理可疑好友申请 请求参数。
type SetDoubtFriendsAddRequestRequest struct {
	// Approve 是否同意 (强制为 true)
	Approve bool `json:"approve"`
	// Flag 请求 flag
	Flag string `json:"flag"`
}

// SetEssenceMsgRequest 设置精华消息 请求参数。
type SetEssenceMsgRequest struct {
	// MessageID 消息ID
	MessageID any `json:"message_id"`
}

// SetFriendAddRequestRequest 处理加好友请求 请求参数。
type SetFriendAddRequestRequest struct {
	// Approve 是否同意请求
	Approve any `json:"approve,omitempty"`
	// Flag 加好友请求的 flag (需从上报中获取)
	Flag string `json:"flag"`
	// Remark 添加后的好友备注
	Remark string `json:"remark,omitempty"`
}

// SetFriendRemarkRequest 设置好友备注 请求参数。
type SetFriendRemarkRequest struct {
	// Remark 备注内容
	Remark string `json:"remark"`
	// UserID 对方 QQ 号
	UserID string `json:"user_id"`
}

// SetGroupAddOptionRequest 设置群加群选项 请求参数。
type SetGroupAddOptionRequest struct {
	// AddType 加群方式
	AddType int64 `json:"add_type"`
	// GroupAnswer 加群答案
	GroupAnswer string `json:"group_answer,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// GroupQuestion 加群问题
	GroupQuestion string `json:"group_question,omitempty"`
}

// SetGroupAddRequestRequest 处理加群请求 请求参数。
type SetGroupAddRequestRequest struct {
	// Approve 是否同意
	Approve any `json:"approve,omitempty"`
	// Count 搜索通知数量
	Count int64 `json:"count,omitempty"`
	// Flag 请求flag
	Flag string `json:"flag"`
	// Reason 拒绝理由
	Reason any `json:"reason,omitempty"`
}

// SetGroupAdminRequest 设置群管理员 请求参数。
type SetGroupAdminRequest struct {
	// Enable 是否设置为管理员
	Enable any `json:"enable,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// SetGroupAlbumMediaLikeRequest 点赞群相册媒体 请求参数。
type SetGroupAlbumMediaLikeRequest struct {
	// AlbumID 相册ID
	AlbumID string `json:"album_id"`
	// BatchID batch_id
	BatchID string `json:"batch_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Lloc lloc，若对整个上传操作则不填
	Lloc string `json:"lloc,omitempty"`
}

// SetGroupBanRequest 群组禁言 请求参数。
type SetGroupBanRequest struct {
	// Duration 禁言时长(秒)
	Duration any `json:"duration"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// SetGroupCardRequest 设置群名片 请求参数。
type SetGroupCardRequest struct {
	// Card 群名片
	Card string `json:"card,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// SetGroupKickRequest 群组踢人 请求参数。
type SetGroupKickRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// RejectAddRequest 是否拒绝加群请求
	RejectAddRequest any `json:"reject_add_request,omitempty"`
	// UserID 用户QQ
	UserID string `json:"user_id"`
}

// SetGroupKickMembersRequest 批量踢出群成员 请求参数。
type SetGroupKickMembersRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// RejectAddRequest 是否拒绝加群请求
	RejectAddRequest any `json:"reject_add_request,omitempty"`
	// UserID QQ号列表
	UserID []string `json:"user_id"`
}

// SetGroupLeaveRequest 退出群组 请求参数。
type SetGroupLeaveRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// IsDismiss 是否解散
	IsDismiss any `json:"is_dismiss,omitempty"`
}

// SetGroupNameRequest 设置群名称 请求参数。
type SetGroupNameRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// GroupName 群名称
	GroupName string `json:"group_name"`
}

// SetGroupPortraitRequest 设置群头像 请求参数。
type SetGroupPortraitRequest struct {
	// File 头像文件路径或 URL
	File string `json:"file"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// SetGroupRemarkRequest 设置群备注 请求参数。
type SetGroupRemarkRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Remark 备注
	Remark string `json:"remark"`
}

// SetGroupRobotAddOptionRequest 设置群机器人加群选项 请求参数。
type SetGroupRobotAddOptionRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// RobotMemberExamine 机器人成员审核
	RobotMemberExamine int64 `json:"robot_member_examine,omitempty"`
	// RobotMemberSwitch 机器人成员开关
	RobotMemberSwitch int64 `json:"robot_member_switch,omitempty"`
}

// SetGroupSearchRequest 设置群搜索选项 请求参数。
type SetGroupSearchRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// NoCodeFingerOpen 未知
	NoCodeFingerOpen int64 `json:"no_code_finger_open,omitempty"`
	// NoFingerOpen 未知
	NoFingerOpen int64 `json:"no_finger_open,omitempty"`
}

// SetGroupSignRequest 群打卡 请求参数。
type SetGroupSignRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// SetGroupSpecialTitleRequest 设置专属头衔 请求参数。
type SetGroupSpecialTitleRequest struct {
	// GroupID 群号
	GroupID string `json:"group_id"`
	// SpecialTitle 专属头衔
	SpecialTitle string `json:"special_title"`
	// UserID QQ号
	UserID string `json:"user_id"`
}

// SetGroupTodoRequest 设置群待办 请求参数。
type SetGroupTodoRequest struct {
	// GroupID 群号
	GroupID any `json:"group_id"`
	// MessageID 消息ID
	MessageID string `json:"message_id,omitempty"`
	// MessageSeq 消息Seq (可选)
	MessageSeq string `json:"message_seq,omitempty"`
}

// SetGroupWholeBanRequest 全员禁言 请求参数。
type SetGroupWholeBanRequest struct {
	// Enable 是否开启全员禁言
	Enable any `json:"enable,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// SetInputStatusRequest 设置输入状态 请求参数。
type SetInputStatusRequest struct {
	// EventType 事件类型
	EventType int64 `json:"event_type"`
	// UserID QQ号
	UserID string `json:"user_id"`
}

// SetMsgEmojiLikeRequest 设置消息表情点赞 请求参数。
type SetMsgEmojiLikeRequest struct {
	// EmojiID 表情ID
	EmojiID any `json:"emoji_id"`
	// MessageID 消息ID
	MessageID any `json:"message_id"`
	// Set 是否设置
	Set any `json:"set,omitempty"`
}

// SetOnlineStatusRequest 设置在线状态 请求参数。
type SetOnlineStatusRequest struct {
	// BatteryStatus 电量状态
	BatteryStatus any `json:"battery_status"`
	// ExtStatus 扩展状态
	ExtStatus any `json:"ext_status"`
	// Status 在线状态
	Status any `json:"status"`
}

// SetQQAvatarRequest 设置QQ头像 请求参数。
type SetQQAvatarRequest struct {
	// File 图片路径、URL或Base64
	File string `json:"file"`
}

// SetQQProfileRequest 设置QQ资料 请求参数。
type SetQQProfileRequest struct {
	// Nickname 昵称
	Nickname string `json:"nickname"`
	// PersonalNote 个性签名
	PersonalNote string `json:"personal_note,omitempty"`
	// Sex 性别 (0: 未知, 1: 男, 2: 女)
	Sex any `json:"sex,omitempty"`
}

// SetRestartRequest 重启服务 请求参数。
type SetRestartRequest struct {
}

// SetSelfLongnickRequest 设置个性签名 请求参数。
type SetSelfLongnickRequest struct {
	// LongNick 签名内容
	LongNick string `json:"longNick"`
}

// TestDownloadStreamRequest 测试下载流 请求参数。
type TestDownloadStreamRequest struct {
	// Error 是否触发测试错误
	Error bool `json:"error,omitempty"`
}

// TransGroupFileRequest 传输群文件 请求参数。
type TransGroupFileRequest struct {
	// FileID 文件ID
	FileID string `json:"file_id"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// TranslateEn2zhRequest 英文单词翻译 请求参数。
type TranslateEn2zhRequest struct {
	// Words 待翻译单词列表
	Words []string `json:"words"`
}

// UploadFileStreamRequest 上传文件流 请求参数。
type UploadFileStreamRequest struct {
	// ChunkData 分块数据 (Base64)
	ChunkData string `json:"chunk_data,omitempty"`
	// ChunkIndex 分块索引
	ChunkIndex int64 `json:"chunk_index,omitempty"`
	// ExpectedSha256 期望的 SHA256
	ExpectedSha256 string `json:"expected_sha256,omitempty"`
	// FileRetention 文件保留时间 (毫秒)
	FileRetention int64 `json:"file_retention"`
	// FileSize 文件总大小
	FileSize int64 `json:"file_size,omitempty"`
	// Filename 文件名
	Filename string `json:"filename,omitempty"`
	// IsComplete 是否完成
	IsComplete bool `json:"is_complete,omitempty"`
	// Reset 是否重置
	Reset bool `json:"reset,omitempty"`
	// StreamID 流 ID
	StreamID string `json:"stream_id"`
	// TotalChunks 总分块数
	TotalChunks int64 `json:"total_chunks,omitempty"`
	// VerifyOnly 是否仅验证
	VerifyOnly bool `json:"verify_only,omitempty"`
}

// UploadGroupFileRequest 上传群文件 请求参数。
type UploadGroupFileRequest struct {
	// File 资源路径或URL
	File string `json:"file"`
	// Folder 父目录 ID
	Folder string `json:"folder,omitempty"`
	// FolderID 父目录 ID (兼容性字段)
	FolderID string `json:"folder_id,omitempty"`
	// GroupID 群号
	GroupID string `json:"group_id"`
	// Name 文件名
	Name string `json:"name"`
	// UploadFile 是否执行上传
	UploadFile bool `json:"upload_file"`
}

// UploadImageToQunAlbumRequest 上传图片到群相册 请求参数。
type UploadImageToQunAlbumRequest struct {
	// AlbumID 相册ID
	AlbumID string `json:"album_id"`
	// AlbumName 相册名称
	AlbumName string `json:"album_name"`
	// File 图片路径、URL或Base64
	File string `json:"file"`
	// GroupID 群号
	GroupID string `json:"group_id"`
}

// UploadPrivateFileRequest 上传私聊文件 请求参数。
type UploadPrivateFileRequest struct {
	// File 资源路径或URL
	File string `json:"file"`
	// Name 文件名
	Name string `json:"name"`
	// UploadFile 是否执行上传
	UploadFile bool `json:"upload_file"`
	// UserID 用户 QQ
	UserID string `json:"user_id"`
}
