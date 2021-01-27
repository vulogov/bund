package bund

import (
	"fmt"

	"github.com/vulogov/Bund/internal/banner"
	"github.com/vulogov/Bund/internal/conf"
)

func Version() {
	Init()
	banner.Banner(fmt.Sprintf("[ theBund %v ]", conf.BVersion))
	banner.Table()
}
