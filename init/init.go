package init

import (
	"github.com/go-xuan/mongox"

	"github.com/go-xuan/quanx/configx"
	"github.com/go-xuan/utilx/errorx"
)

func init() {
	errorx.Panic(Init())
}

func Init() error {
	var err error
	if err = configx.LoadConfigurator(&mongox.Configs{}); err == nil && mongox.Initialized() {
		return nil
	} else if err = configx.LoadConfigurator(&mongox.Config{}); err == nil && mongox.Initialized() {
		return nil
	}
	return errorx.Wrap(err, "init mongo failed")
}
