package shortURL

import (
	"compressURL/internal/dao"
	"compressURL/internal/model/entity"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
)

// BatchImport 批量导入
func (s *sShortURL) BatchImport(ctx context.Context, in []entity.ShortUrl) ([]string, error) {
	// 获取一条未使用短链 code
	var shortCodeDataList []entity.ShortUrlCode
	err := dao.ShortUrlCode.Ctx(ctx).Where("status", 0).Limit(len(in)).Scan(&shortCodeDataList)
	if err != nil {
		return nil, gerror.Newf("批量获取短链 code 错误 %s", err)
	}

	var shortCodeList []string
	for i, _ := range in {
		in[i].ShortUrl = shortCodeDataList[i].Code
		shortCodeList = append(shortCodeList, shortCodeDataList[i].Code)
	}

	// 创建事务
	db := g.DB()
	if tx, err := db.Begin(ctx); err == nil {

		_, err := tx.Model("short_url").Data(in).Save()

		if err != nil {
			err := tx.Rollback()
			return nil, err
		}

		_, err = tx.Model("short_url_code").Data(g.Map{"status": 1}).WhereIn("code", shortCodeList).Update()
		if err != nil {
			err := tx.Rollback()
			return nil, err
		}

		// 提交事务
		err = tx.Commit()
		return nil, err
	}

	return nil, err
}
