package tgraph

import (
	"fmt"
	"go.uber.org/zap"
	"html"
	"math/rand"
	"time"
)

var （
bitLink = "https://t.me/Bitczx"
bitTitle = "币快讯"
tigTitle = "Tiger"
）

func PublishHtml(bitTitle string, title string, bitLink string, htmlContent string) (string, error) {
	//html = fmt.Sprintf(
	//	"<p>本文章由 <a href=\"https://github.com/indes/flowerss-bot\">flowerss</a> 抓取自RSS，版权归<a href=\"\">源站点</a>所有。</p><hr>",
	//) + html + fmt.Sprintf(
	//	"<hr><p>本文章由 <a href=\"https://github.com/indes/flowerss-bot\">flowerss</a> 抓取自RSS，版权归<a href=\"\">源站点</a>所有。</p><p>查看原文：<a href=\"%s\">%s - %s</p>",
	//	rawLink,
	//	title,
	//	sourceTitle,
	//)

	htmlContent = html.UnescapeString(htmlContent) + fmt.Sprintf(
		"<hr><p><a href=\"%s\"><h1>更多内容</h1></a></p>",
		bitLink,
	//	title,
	//	sourceTitle,
	)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	client := clientPool[rand.Intn(len(clientPool))]

	if page, err := client.CreatePageWithHTML(title+" - "+bitTitle, tigTitle, bitLink, htmlContent, true); err == nil {
		zap.S().Infof("已创建 telegraph  url: %s", page.URL)
		return page.URL, err
	} else {
		zap.S().Warnf("创建telegraph失败，错误: %s", err)
		return "", nil
	}
}
