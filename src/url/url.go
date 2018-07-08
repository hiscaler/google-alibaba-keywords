package url

import (
	"github.com/go-ozzo/ozzo-config"
	conf "config"
)

var cfg *config.Config
var prefix string

func init() {
	cfg = conf.Instance()
	if cfg.GetBool("debug") {
		prefix = cfg.GetString("api.prefix.debug")
	} else {
		prefix = cfg.GetString("api.prefix.prod")
	}
}
func KeywordIndex() string {
	return prefix + cfg.GetString("api.keyword.index")
}

func KeywordSubmit() string  {
	return prefix + cfg.GetString("api.keyword.submit")
}

func ProductSubmit() string {
	return prefix + cfg.GetString("api.product.submit")
}
