package util

import (
	"go.uber.org/zap"
	"net/http"
	"time"

	"github.com/leozgu/flowerss-bot/config"
	"golang.org/x/net/proxy"
)

var (
	HttpClient *http.Client
)

func clientInit() {
	httpTransport := &http.Transport{}
	HttpClient = &http.Client{Transport: httpTransport, Timeout: 2 * time.Second}
	// set proxy
	if config.Socks5 != "" {
		zap.S().Infow("enable proxy",
			"socks5", config.Socks5,
		)

		dialer, err := proxy.SOCKS5("tcp", config.Socks5, nil, proxy.Direct)
		if err != nil {
			zap.S().Fatal("Error creating dialer, aborting.")
		}
		httpTransport.Dial = dialer.Dial
	}
}
