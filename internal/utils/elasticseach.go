package utils

import (
	"crypto/tls"
	"net"
	"net/http"
	"schoolLifeDemo/internal/svc"
	"time"

	"github.com/elastic/go-elasticsearch/v7"
)

func GetEsClient(svcCtx *svc.ServiceContext) (*elasticsearch.Client, error) {
	cfg := elasticsearch.Config{
		Addresses: svcCtx.Config.ElasticSearch.Addresses,
		Username:  svcCtx.Config.ElasticSearch.UserName,
		Password:  svcCtx.Config.ElasticSearch.PassWord,
		Transport: &http.Transport{
			MaxIdleConnsPerHost:   10,
			ResponseHeaderTimeout: time.Second * 5,
			DialContext:           (&net.Dialer{Timeout: time.Second}).DialContext,
			TLSClientConfig: &tls.Config{
				MinVersion: tls.VersionTLS12,
			},
		},
	}
	return elasticsearch.NewClient(cfg)
}
