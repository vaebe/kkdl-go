package login

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"

	"compressURL/api/login/v1"
)

var wsClientMap = make(map[string]*ghttp.WebSocket)

func (c *ControllerV1) Ws(ctx context.Context, req *v1.WsReq) (res *v1.WsRes, err error) {
	r := g.RequestFromCtx(ctx)
	ws, err := r.WebSocket()
	if err != nil {
		glog.Error(ctx, err)
		r.Exit()
	}

	// 用户 code 不存在返回错误信息并关闭连接
	if req.UserCode == "" {
		if err = ws.WriteMessage(1, []byte(`{"data": "用户code不能为空!"}`)); err != nil {
			return nil, err
		}

		err = ws.Close()
		if err != nil {
			return nil, err
		}
	}

	wsClientMap[req.UserCode] = ws
	if err = ws.WriteMessage(1, []byte(`{"status": "ok"}`)); err != nil {
		return nil, err
	}

	for {
		_, _, _ = ws.ReadMessage()
	}
}
