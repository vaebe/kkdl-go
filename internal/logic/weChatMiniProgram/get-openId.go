package weChatMiniProgram

import (
	v1 "compressURL/api/weChatMiniProgram/v1"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
	"io"
	"net/http"
)

// GetOpenId 获取微信小程序用户openId
func (s *sWeChatMiniProgram) GetOpenId(ctx context.Context, Code string) (resInfo *v1.GetOpenIdRes, err error) {
	// 获取微信小程序配置
	miniProgramsConf := g.Cfg().MustGet(ctx, "weChat.miniPrograms").Map()
	url := fmt.Sprintf("https://api.weixin.qq.com/sns/jscode2session?grant_type=client_credential&appid=%s&secret=%s&js_code=%s&grant_type=authorization_code", miniProgramsConf["appid"], miniProgramsConf["secret"], Code)
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			glog.Error(ctx, "调用获取微信 openid 接口失败", err)
		}
	}(res.Body)

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	// 解析JSON数据
	err = json.Unmarshal(body, &resInfo)

	if resInfo.ErrCode != 0 {
		return nil, gerror.Newf("获取用户openId失败:%s", resInfo.ErrMsg)
	}
	return resInfo, err
}
