package formatters

import (
	m "stretches-common-api/models"

	"stretches-common-api/publicid"
)

type ArticleFormattedItem struct {
	Pid         uint32              `json:"pid"`
	Content     string              `json:"content"`
	Title       string              `json:"title"`
	Pages       []PageFormattedItem `json:"pages"`
	DurationMin int                 `json:"duration_min"`
}

func FormatArticles(items []m.Article, langCode string) []ArticleFormattedItem {
	res := []ArticleFormattedItem{}
	for _, v := range items {
		res = append(res, FormatArticle(v, langCode))
	}
	return res
}
func FormatArticle(v m.Article, langCode string) ArticleFormattedItem {
	res := ArticleFormattedItem{
		Pid:     publicid.Obfuscate32bit(v.ID),
		Content: v.GetContent(langCode),
		Title:   v.GetTitle(langCode),

		DurationMin: v.DurationMin,
	}

	res.Pages = FormatPages(v.Pages, langCode)
	return res
}

type PageFormattedItem struct {
	Pid     uint32 `json:"pid"`
	Title   string `json:"title"`
	Content string `json:"content"`
	Idx     uint32 `json:"idx"`
}

func FormatPages(items []m.Page, langCode string) []PageFormattedItem {
	res := []PageFormattedItem{}
	for _, v := range items {
		res = append(res, FormatPage(v, langCode))
	}
	return res
}
func FormatPage(v m.Page, langCode string) PageFormattedItem {
	res := PageFormattedItem{
		Pid:     publicid.Obfuscate32bit(v.ID),
		Content: v.GetContent(langCode),
		Title:   v.GetTitle(langCode),
		Idx:     v.Idx,
	}
	return res
}
