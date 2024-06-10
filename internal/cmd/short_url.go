package cmd

import (
	"compressURL/internal/model/entity"
	"compressURL/internal/service"
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/mssola/useragent"
	"io"
	"net/http"
)

// IPInfo 结构体用于存储IP地址的详细信息
type IPInfo struct {
	Query         string  `json:"query"`         // 查询的IP地址
	Status        string  `json:"status"`        // 查询状态，成功或失败
	Continent     string  `json:"continent"`     // 大洲名称
	ContinentCode string  `json:"continentCode"` // 大洲代码
	Country       string  `json:"country"`       // 国家名称
	CountryCode   string  `json:"countryCode"`   // 国家代码
	Region        string  `json:"region"`        // 地区或州的短代码（FIPS或ISO）
	RegionName    string  `json:"regionName"`    // 地区或州名称
	City          string  `json:"city"`          // 城市名称
	District      string  `json:"district"`      // 位置的区
	Lat           float64 `json:"lat"`           // 纬度
	Lon           float64 `json:"lon"`           // 经度
}

func getIpInfo(ctx context.Context, clientIp string) (IPInfo, error) {
	apiUrl := fmt.Sprintf("http://ip-api.com/json/%s?fields=status,message,continent,continentCode,country,countryCode,region,regionName,city,district,lat,lon,query&lang=zh-CN", clientIp)
	res, err := g.Client().Get(ctx, apiUrl)

	if err != nil {
		return IPInfo{}, gerror.Newf("failed to make request: %s", err)
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		return IPInfo{}, gerror.Newf("unexpected status code: %d", res.StatusCode)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return IPInfo{}, gerror.Newf("failed to read response body: %s", err)
	}

	var ipInfo IPInfo
	if err := json.Unmarshal(body, &ipInfo); err != nil {
		return IPInfo{}, gerror.Newf("failed to unmarshal JSON: %s", err)
	}

	if ipInfo.Status != "success" {
		return IPInfo{}, gerror.Newf("IP info retrieval failed: %s", ipInfo.Status)
	}

	return ipInfo, nil
}

func saveVisitsInfo(ctx context.Context, r *ghttp.Request, shortUrlInfo entity.ShortUrl) {
	clientIp := r.GetClientIp()

	ipInfo, err := getIpInfo(ctx, clientIp)

	if err != nil {
		g.Log().Error(ctx, "获取 ip 信息失败:", err)
	}

	curUa := r.Header.Get("User-Agent")
	uaInfo := useragent.New(curUa)

	browserName, browserVersion := uaInfo.Browser()
	engineName, engineVersion := uaInfo.Engine()

	err = service.ShortUrlVisits().Create(ctx, entity.ShortUrlVisits{
		Id:             0,
		UserId:         shortUrlInfo.UserId,
		ShortUrl:       shortUrlInfo.ShortUrl,
		RawUrl:         shortUrlInfo.RawUrl,
		UserAgent:      curUa,
		BrowserName:    browserName,
		BrowserVersion: browserVersion,
		DeviceModel:    uaInfo.Platform(),
		EngineName:     engineName,
		EngineVersion:  engineVersion,
		OsName:         uaInfo.OSInfo().Name,
		OsVersion:      uaInfo.OSInfo().Version,
		Ip:             clientIp,
		Continent:      ipInfo.Continent,
		ContinentCode:  ipInfo.ContinentCode,
		Country:        ipInfo.Country,
		CountryCode:    ipInfo.CountryCode,
		Region:         ipInfo.Region,
		RegionName:     ipInfo.RegionName,
		City:           ipInfo.City,
		District:       ipInfo.District,
		Lat:            ipInfo.Lat,
		Lon:            ipInfo.Lon,
		CreatedAt:      nil,
		UpdatedAt:      nil,
		DeletedAt:      nil,
	})

	if err != nil {
		g.Log().Error(ctx, "创建短链访问信息失败:", err)
	}
}

// 注册短链服务
func registerShortURLService(s *ghttp.Server, ctx context.Context) {
	s.BindHandler("/:id", func(r *ghttp.Request) {
		code := gstr.SubStr(r.Request.RequestURI, 1, len(r.Request.RequestURI))

		if code == "favicon.ico" {
			return
		}

		req, err := service.ShortUrl().GenOne(ctx, code)

		if err != nil {
			r.Response.Write("未获取到对应的地址，请检查链接是否正确！")
			return
		}

		go saveVisitsInfo(ctx, r, req)

		http.Redirect(r.Response.ResponseWriter, r.Request, req.RawUrl, http.StatusFound)
	})
}
