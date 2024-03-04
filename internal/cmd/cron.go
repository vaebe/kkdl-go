package cmd

import (
	"compressURL/internal/service"
	"github.com/gogf/gf/v2/os/gcron"
	"github.com/gogf/gf/v2/os/glog"
	"golang.org/x/net/context"
)

func GenerateShortLinkScheduledTask(ctx context.Context) (err error) {
	// 启动一个单例定时任务
	_, err = gcron.AddSingleton(ctx, "0 0 3 * * *", func(ctx context.Context) {
		_ = service.ShortURL().BatchCreateShortUrlCode(ctx, 1000)
		glog.Info(ctx, "定时任务生成短链 code \r\n")
	}, "GenerateShortLinkScheduledTask")
	return err
}
