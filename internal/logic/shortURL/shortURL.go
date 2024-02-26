package shortURL

import (
	v1 "compressURL/api/shortURL/v1"
	"compressURL/internal/dao"
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/net/context"
)

type sShortURL struct {
}

func init() {
	service.RegisterShortURL(New())
}

func New() *sShortURL {
	return &sShortURL{}
}

// CreateShortURL 创建短链
func (s *sShortURL) CreateShortURL(ctx context.Context, in entity.ShortUrl) (string, error) {
	// 获取一条未使用短链 code
	shortCodeData := entity.ShortUrlCode{}
	err := dao.ShortUrlCode.Ctx(ctx).Where("status", 0).Limit(1).Scan(&shortCodeData)
	if err != nil {
		return "", err
	}

	// 创建事务
	db := g.DB()
	if tx, err := db.Begin(ctx); err == nil {
		_, err := tx.Model("short_url").
			Data(g.Map{"shortUrl": shortCodeData.Code, "rawUrl": in.RawUrl, "expirationTime": in.ExpirationTime, "title": in.Title, "groupId": in.GroupId}).
			Save()

		if err != nil {
			return "", err
		}

		_, err = tx.Model("short_url_code").Data(g.Map{"status": 1}).Where("code", shortCodeData.Code).Update()
		if err != nil {
			// 回滚事务
			tx.Rollback()
			return "", err
		}

		// 提交事务
		tx.Commit()
	}

	return shortCodeData.Code, err
}

// GetShortURL 获取短链
func (s *sShortURL) GetShortURL(ctx context.Context, url string) (string, error) {
	one, err := dao.ShortUrl.Ctx(ctx).Where("shortUrl", url).One()
	if err != nil {
		return "", err
	}

	if one == nil {
		return "", errors.New("未查询到对应的数据")
	}

	rawUrl := fmt.Sprintf("%s", one.GMap().Get("rawUrl"))
	return rawUrl, nil
}

// GetList 短链列表
func (s *sShortURL) GetList(ctx context.Context, in v1.GetListReq, userId string) ([]entity.ShortUrl, int, error) {
	var list []entity.ShortUrl

	db := dao.ShortUrl.Ctx(ctx).OmitEmptyWhere().
		Where(dao.ShortUrl.Columns().Title, in.Title).
		Where(dao.ShortUrl.Columns().RawUrl, in.RawUrl)

	// 用户 id 存在只查询当前用户的数据
	if userId != "" {
		db = db.Where(dao.ShortUrl.Columns().UserId, userId)
	}

	total, _ := db.Count()

	err := db.OrderDesc(dao.ShortUrl.Columns().CreatedAt).Page(in.PageNo, in.PageSize).Scan(&list)

	if err != nil {
		return nil, 0, err
	}
	return list, total, nil
}

// Delete 删除短链
func (s *sShortURL) Delete(ctx context.Context, id string) error {
	res, err := dao.ShortUrl.Ctx(ctx).Where(dao.ShortUrl.Columns().Id, id).Delete()

	if num, _ := res.RowsAffected(); num == 0 {
		return gerror.New("需要删除的数据不存在！")
	}

	return err
}
