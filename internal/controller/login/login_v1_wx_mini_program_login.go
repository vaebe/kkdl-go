package login

import (
	"compressURL/api/login/v1"
	"compressURL/internal/model"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
)

func isUserRegistered(ctx context.Context, openId string) (bool, error) {
	info, err := service.User().GetUserInfoByWxId(ctx, openId)

	// 存在错误信息直接返回
	if info != nil && err != nil {
		return false, err
	}
	return info != nil, nil
}

// 向客户端发送 ws 消息
func sendWsMessage(data v1.WxMiniProgramLoginRes, userCode string) (err error) {
	ws, ok := wsClientMap[userCode]
	if !ok {
		return gerror.New("ws 链接不存在!")
	}

	info := gjson.NewWithTag(data, "json", true)
	err = info.Set("dataType", "loginSuccess")
	if err != nil {
		return gerror.Newf("序列化数据错误! %s", err)
	}

	err = ws.WriteMessage(1, []byte(info.String()))
	if err != nil {
		return gerror.Newf("发送 ws 消息失败! %s", err)
	}
	return err
}

func (c *ControllerV1) WxMiniProgramLogin(ctx context.Context, req *v1.WxMiniProgramLoginReq) (res *v1.WxMiniProgramLoginRes, err error) {
	// 获取小程序 openId
	wxUserInfo, err := service.WeChatMiniProgram().GetOpenId(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	// 验证用户是否已经注册
	userRegistered, err := isUserRegistered(ctx, wxUserInfo.Openid)
	if err != nil {
		return nil, gerror.Newf("验证用户是否注册失败! %s", err)
	}

	// 用户不存在创建微信用户
	if userRegistered == false {
		_, err = service.User().Create(ctx, entity.User{
			WxId:        wxUserInfo.Openid,
			AccountType: "02",
			Role:        "01",
		})
		if err != nil {
			return nil, gerror.Newf("创建微信用户失败! %s", err)
		}
	}

	// 生成 token
	userInfo, token, tokenExpire, err := service.Login().UserLogin(ctx, model.LoginInput{WxId: wxUserInfo.Openid, AccountType: "02"})
	if err != nil {
		return nil, err
	}

	res = &v1.WxMiniProgramLoginRes{
		Token:       token,
		TokenExpire: tokenExpire,
		UserInfo:    userInfo,
	}

	err = sendWsMessage(*res, req.UserCode)

	if err != nil {
		return nil, err
	}

	return res, nil
}
