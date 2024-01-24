package common

import (
	v1 "compressURL/api/common/v1"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"golang.org/x/net/context"
)

func (s *sCommon) UploadFile(ctx context.Context, file *ghttp.UploadFile) (out v1.UploadFileRes, err error) {
	filesReader, _ := file.Open()

	qiNiuConfig := g.Cfg().MustGet(ctx, "qiNiuConfig").Map()

	mac := qbox.NewMac(qiNiuConfig["access"].(string), qiNiuConfig["secret"].(string))
	putPolicy := storage.PutPolicy{
		Scope: qiNiuConfig["bucket"].(string),
	}
	upToken := putPolicy.UploadToken(mac)

	cfg := storage.Config{
		Region:        &storage.ZoneHuadongZheJiang2, // 空间对应的机房
		UseHTTPS:      true,                          // 是否使用https域名
		UseCdnDomains: false,                         // 上传是否使用CDN上传加速
	}

	// 构建表单上传的对象
	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}

	// 可选配置
	putExtra := storage.PutExtra{
		Params: map[string]string{
			"x:name": "ddl file",
		},
	}

	key := fmt.Sprintf("kkdl/%s/%s-%s", gtime.Now().Format("Y-m-d"), guid.S(), file.Filename)
	err = formUploader.Put(context.Background(), &ret, upToken, key, filesReader, file.Size, &putExtra)
	if err != nil {
		return v1.UploadFileRes{}, err
	}

	fileUrl := qiNiuConfig["baseUrl"].(string) + key
	glog.Info(ctx, "文件访问路径:", fileUrl)

	return v1.UploadFileRes{
		Url:  fileUrl,
		Name: file.Filename,
	}, err
}
