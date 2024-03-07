package weChatMiniProgram

import (
	"bytes"
	"compressURL/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/glog"
	"io"
	"log"

	"compressURL/api/weChatMiniProgram/v1"
)

func (c *ControllerV1) GetMiniProgramCode(ctx context.Context, req *v1.GetMiniProgramCodeReq) (res *v1.GetMiniProgramCodeRes, err error) {
	token, err := service.WeChatMiniProgram().GetWeChatToken(ctx)
	if err != nil {
		return nil, err
	}

	url := fmt.Sprintf("https://api.weixin.qq.com/wxa/getwxacodeunlimit?access_token=%s", token)

	// 定义要发送的 map 数据
	opts := fmt.Sprintf(`{"scene": "%s", "page": "%s"}`, req.Scene, req.Page)

	resData, err := g.Client().Post(ctx, url, opts)
	if err != nil {
		panic(err)
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			glog.Error(ctx, "调用获取微信小程序码接口失败", err)
		}
	}(resData.Body)

	contentType := resData.Header.Get("Content-Type")

	// 判断返回数据是否是 json 类型,json 类型代表响应错误
	if "application/json; encoding=utf-8" == contentType {
		body, err := io.ReadAll(resData.Body)
		if err != nil {
			return nil, err
		}

		// 解析JSON数据
		var resp v1.GetMiniProgramCodeRes
		if err := json.Unmarshal(body, &resp); err != nil {
			return nil, err
		}
		return &resp, nil
	}

	// 否则读取响应的内容，返回小程序码
	body := new(bytes.Buffer)
	_, err = body.ReadFrom(resData.Body)
	if err != nil {
		log.Fatal(err)
	}

	g.RequestFromCtx(ctx).Response.Write(body)
	return nil, nil
}
