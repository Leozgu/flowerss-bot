package tgraph

import (
	"fmt"
	"go.uber.org/zap"
	"html"
	"math/rand"
	"time"
)

var bicLink = "https://t.me/Bitczx"

func PublishHtml(sourceTitle string, title string, rawLink string, htmlContent string) (string, error) {
	//html = fmt.Sprintf(
	//	"<p>本文章由 <a href=\"https://github.com/indes/flowerss-bot\">flowerss</a> 抓取自RSS，版权归<a href=\"\">源站点</a>所有。</p><hr>",
	//) + html + fmt.Sprintf(
	//	"<hr><p>本文章由 <a href=\"https://github.com/indes/flowerss-bot\">flowerss</a> 抓取自RSS，版权归<a href=\"\">源站点</a>所有。</p><p>查看原文：<a href=\"%s\">%s - %s</p>",
	//	rawLink,
	//	title,
	//	sourceTitle,
	//)

	htmlContent = html.UnescapeString(htmlContent) + fmt.Sprintf(
		"<hr><p><a href=\"https://t.me/Bitczx\"><h1>更多内容</h1></a><br/><a href=\"\">.</a></p><p>.<a href=\"%s\">%s</p>",
		bicLink,
		title,
	//	sourceTitle,
	)
	rand.Seed(time.Now().Unix()) // initialize global pseudo random generator
	client := clientPool[rand.Intn(len(clientPool))]

	if page, err := client.CreatePageWithHTML(title+" - "+sourceTitle, sourceTitle, rawLink, htmlContent, true); err == nil {
		zap.S().Infof("Created telegraph page url: %s", page.URL)
		return page.URL, err
	} else {
		zap.S().Warnf("Create telegraph page failed, error: %s", err)
		return "", nil
	}
}