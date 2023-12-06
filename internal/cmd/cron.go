package cmd

import (
	"compressURL/internal/dao"
	"compressURL/utility"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcron"
	"golang.org/x/net/context"
)

func GenerateShortLinkScheduledTask(ctx context.Context) (err error) {

	// 启动一个单例定时任务
	_, err = gcron.AddSingleton(ctx, "0 0 3 * * *", func(ctx context.Context) {
		g.Log().Info(ctx, "开始生成短链")
		// 记录插入数量
		insertDataNum := 0

		// 写一个循环成功插入指定数量的数据后退出
		for {
			shortUrl := utility.IntToBase62(utility.GenerateRandomNumber(9999999999, 1))

			_, err := dao.ShortUrlCode.Ctx(ctx).Insert(g.Map{"Code": shortUrl, "Status": 0})
			if err != nil {
				g.Log().Info(ctx, "生成短链重复", shortUrl)
				return
			}

			insertDataNum++

			if insertDataNum > 999 {
				break
			}
		}

	}, "GenerateShortLinkScheduledTask")
	return err
}
