package lang

import (
	"fmt"
	"strings"

	"github.com/gofiber/fiber/v2"
)

func GetLangContext(c *fiber.Ctx) LocaleContext {
	langHeader := c.Get("X-Lang")
	rawlangCode := ""
	if len(langHeader) > 0 && langHeader != "en" {
		rawlangCode = langHeader
	}
	langCode := GetLangCode2(rawlangCode)
	return LocaleContext{
		Code: langCode, RawCode: rawlangCode, UsesTranslation: len(rawlangCode) > 0}
}
func GetLangCode2(code string) string {
	code = strings.ReplaceAll(code, "-", "_")
	code = strings.ToLower(code)
	codes := []string{"es", "fr", "de"}
	for _, c := range codes {
		if c == code || strings.Contains(code, fmt.Sprintf("%v_", c)) {
			return c
		}
	}
	return ""
}

type LocaleContext struct {
	Code            string `json:"code"`
	RawCode         string `json:"rawcode"`
	UsesTranslation bool   `json:"uses_translation"`
}

func (c LocaleContext) GetLanguageName() string {
	switch c.Code {
	case "":
		return "English"
	case "es":
		return "Spanish"
	case "fr":
		return "French"
	case "de":
		return "German"
	default:
		return "English"

	}
}
