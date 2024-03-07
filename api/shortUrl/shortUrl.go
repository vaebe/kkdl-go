// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package shortUrl

import (
	"context"

	"compressURL/api/shortUrl/v1"
)

type IShortUrlV1 interface {
	BatchExport(ctx context.Context, req *v1.BatchExportReq) (res *v1.BatchExportRes, err error)
	BatchImport(ctx context.Context, req *v1.BatchImportReq) (res *v1.BatchImportRes, err error)
	Create(ctx context.Context, req *v1.CreateReq) (res *v1.CreateRes, err error)
	Delete(ctx context.Context, req *v1.DeleteReq) (res *v1.DeleteRes, err error)
	GetList(ctx context.Context, req *v1.GetListReq) (res *v1.GetListRes, err error)
	GetUrl(ctx context.Context, req *v1.GetUrlReq) (res *v1.GetUrlRes, err error)
	TemplateDownload(ctx context.Context, req *v1.TemplateDownloadReq) (res *v1.TemplateDownloadRes, err error)
}
