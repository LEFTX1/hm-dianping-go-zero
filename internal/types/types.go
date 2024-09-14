// Code generated by goctl. DO NOT EDIT.
package types

type AddVoucherRequest struct {
	ShopId      int64  `json:"shop_id"`      // 商铺ID
	Title       string `json:"title"`        // 标题
	SubTitle    string `json:"sub_title"`    // 副标题
	Rules       string `json:"rules"`        // 规则
	PayValue    int64  `json:"pay_value"`    // 支付金额
	ActualValue int64  `json:"actual_value"` // 实际金额
	VoucherType int32  `json:"type"`         // 优惠券类型
	Status      int32  `json:"status"`       // 状态
	Stock       int32  `json:"stock"`        // 库存
	BeginTime   string `json:"begin_time"`   // 开始时间
	EndTime     string `json:"end_time"`     // 结束时间
}

type Blog struct {
	Id         int64  `json:"id"`          // 博文ID
	ShopId     int64  `json:"shop_id"`     // 商铺ID
	UserId     int64  `json:"user_id"`     // 用户ID
	Icon       string `json:"icon"`        // 图标
	Name       string `json:"name"`        // 名称
	Title      string `json:"title"`       // 标题
	Images     string `json:"images"`      // 图片
	Content    string `json:"content"`     // 内容
	Liked      int32  `json:"liked"`       // 点赞数
	Comments   int32  `json:"comments"`    // 评论数
	CreateTime string `json:"create_time"` // 创建时间
	UpdateTime string `json:"update_time"` // 更新时间
}

type BlogResponse struct {
	Success bool   `json:"success"` // 是否成功
	Data    []Blog `json:"data"`    // 博文列表
}

type EmptyRequest struct {
}

type EmptyResponse struct {
}

type IdResponse struct {
	Id int64 `json:"id"` // ID
}

type LikeBlogRequest struct {
	Id int64 `path:"id"` // 博文ID
}

type LoginFormDTO struct {
	Phone    string `json:"phone"`    // 手机号
	Code     string `json:"code"`     // 验证码
	Password string `json:"password"` // 密码
}

type LoginResult struct {
	Success bool   `json:"success"` // 是否成功
	Data    string `json:"data"`    // 返回数据
}

type QueryHotBlogRequest struct {
	Current int32 `form:"current"` // 当前页码
}

type QueryMyBlogRequest struct {
	Current int32 `form:"current"` // 当前页码
}

type QueryVoucherOfShopReply struct {
	Success bool      `json:"success"` // 是否成功
	Data    []Voucher `json:"data"`    // 优惠券列表
}

type QueryVoucherOfShopRequest struct {
	ShopId int64 `path:"shop_id"` // 商铺ID
}

type SaveBlogRequest struct {
	ShopId  int64  `json:"shop_id"` // 商铺ID
	UserId  int64  `json:"user_id"` // 用户ID
	Title   string `json:"title"`   // 标题
	Images  string `json:"images"`  // 图片
	Content string `json:"content"` // 内容
}

type SeckillVoucherRequest struct {
	Id int64 `path:"id"` // 优惠券ID
}

type SendCodeRequest struct {
	Phone string `json:"phone"` // 手机号
}

type SendCodeResult struct {
	Success bool `json:"success"` // 是否成功
}

type TestResponse struct {
	Message string `json:"message"` // 消息
}

type UserDTO struct {
	Id       int64  `json:"id"`        // 用户ID
	NickName string `json:"nick_name"` // 昵称
	Icon     string `json:"icon"`      // 图标
}

type UserDTOReply struct {
	Success bool    `json:"success"` // 是否成功
	Data    UserDTO `json:"data"`    // 用户数据
}

type Voucher struct {
	Id          int64  `json:"id"`           // 优惠券ID
	ShopId      int64  `json:"shop_id"`      // 商铺ID
	Title       string `json:"title"`        // 标题
	SubTitle    string `json:"sub_title"`    // 副标题
	Rules       string `json:"rules"`        // 规则
	PayValue    int64  `json:"pay_value"`    // 支付金额
	ActualValue int64  `json:"actual_value"` // 实际金额
	VoucherType int32  `json:"type"`         // 优惠券类型
	Status      int32  `json:"status"`       // 状态
	Stock       int32  `json:"stock"`        // 库存
	BeginTime   string `json:"begin_time"`   // 开始时间
	EndTime     string `json:"end_time"`     // 结束时间
}

type VoucherResult struct {
	Success  bool   `json:"success"`   // 是否成功
	ErrorMsg string `json:"error_msg"` // 错误信息
	Data     int64  `json:"data"`      // 数据
}
