package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
)

/**
1. 客户端请求一个 code
2. 客户端请求生成小程序码接口的时候携带上边获取的 code
3. 客户端获取到小程序码后,开始轮询改 code 的状态直到返回成功 (验证 code 状态的接口,成功后返回用户信息)
4. 手机微信扫码-获取小程序码携带的 code 调用接口更新该 code (调用 wx login 后直接登录 or 点击按钮确认)
*/

func (c *ControllerV1) WxMiniProgramLogin(ctx context.Context, req *v1.WxMiniProgramLoginReq) (res *v1.WxMiniProgramLoginRes, err error) {
	// 获取小程序 openId
	wxUserInfo, err := service.WeChatMiniProgram().GetOpenId(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// todo 后续微信小程序登录 Openid 存储到 id 字段
	userInfo, err := service.User().Detail(ctx, model.UserQueryInput{Id: wxUserInfo.Openid})

	if err != nil {
		return nil, err
	}

	// 用户信息不存在则创建
	if userInfo == nil {
		userId, err := service.User().Create(ctx, entity.User{
			Id:          wxUserInfo.Openid,
			AccountType: "02",
			Role:        "01",
		})
		if err != nil {
			return nil, gerror.Newf("创建微信用户失败! %s", err)
		}

		// 创建完成再次查询用户数据返回
		userInfo, _ = service.User().Detail(ctx, model.UserQueryInput{Id: userId})
	}

	info := getLoginRes(ctx, *userInfo)
	return (*v1.WxMiniProgramLoginRes)(info), nil
}
