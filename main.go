package main

import (
	"github.com/thk-im/thk-im-base-server/conf"
	"github.com/thk-im/thk-im-contact-server/pkg/app"
	"github.com/thk-im/thk-im-contact-server/pkg/handler"
)

func main() {
	configPath := "etc/contact_server.yaml"
	config := &conf.Config{}
	if err := conf.LoadConfig(configPath, config); err != nil {
		panic(err)
	}

	appCtx := &app.Context{}
	appCtx.Init(config)
	handler.RegisterContactApiHandlers(appCtx)

	appCtx.StartServe()
}
