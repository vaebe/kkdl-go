package shortURL

import (
	"compressURL/internal/dao"
	"compressURL/utility"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
)

// BatchCreateShortUrlCode 批量生成短链 code
func (s *sShortURL) BatchCreateShortUrlCode(ctx context.Context, num int) error {
	g.Log().Info(ctx, "开始生成短链")
	// 记录插入数量
	insertDataNum := 0

	// 写一个循环成功插入指定数量的数据后退出
	for {
		shortUrl := utility.IntToBase62(utility.GenerateRandomNumber(9999999999, 1))

		_, err := dao.ShortUrlCode.Ctx(ctx).Insert(g.Map{"Code": shortUrl, "Status": 0})
		if err == nil {
			insertDataNum++

			if insertDataNum > num {
				g.Log().Info(ctx, "生成短链完成", shortUrl)
				break
			}
		} else {
			g.Log().Info(ctx, "生成短链重复", shortUrl)
		}
	}

	return nil
}
