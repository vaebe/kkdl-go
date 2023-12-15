package login

import (
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
	"io"
	"net/http"
)

// WeChatTokenMap 解析接口响应
type WeChatTokenMap struct {
	AccessToken string `json:"access_token"`
	ExpiresIn   int64  `json:"expires_in"`
}

// 设置微信 token 缓存
func setWeChatToken(ctx context.Context, req WeChatTokenMap) {
	err := g.Redis().SetEX(ctx, "weChatToken", req.AccessToken, req.ExpiresIn)
	if err != nil {
		glog.Error(ctx, "redis 缓存 weChatToken 失败", err)
	}
}

// 获取缓存的 weChatToken
func getWeChatToken(ctx context.Context) string {
	weChatToken, err := g.Redis().Get(ctx, "weChatToken")
	if err != nil {
		glog.Error(ctx, "获取 redis 缓存 weChatToken 失败", err)
		return ""
	}

	return weChatToken.String()
}

// GetWeChatToken 获取微信 api token
func (s *sLogin) GetWeChatToken(ctx context.Context) (string, error) {
	// 缓存的 weChatToken 存在直接返回
	token := getWeChatToken(ctx)
	if token != "" {
		return token, nil
	}

	// 获取微信小程序配置
	miniProgramsConf := g.Cfg().MustGet(ctx, "weChat.miniPrograms").Map()

	url := fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/token?grant_type=client_credential&appid=%s&secret=%s", miniProgramsConf["appid"], miniProgramsConf["secret"])
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			glog.Error(ctx, "调用获取微信 token 接口失败", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	// 解析JSON数据
	var resp WeChatTokenMap
	if err := json.Unmarshal(body, &resp); err != nil {
		return "", err
	}

	setWeChatToken(ctx, resp)

	return resp.AccessToken, nil
}
