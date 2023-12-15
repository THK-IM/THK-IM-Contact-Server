package app

import (
	"github.com/thk-im/thk-im-base-server/conf"
	"github.com/thk-im/thk-im-base-server/server"
	"github.com/thk-im/thk-im-contact-server/pkg/loader"
	"github.com/thk-im/thk-im-contact-server/pkg/model"
)

type Context struct {
	*server.Context
	modelMap map[string]interface{}
}

func (c *Context) UserContactModel() model.UserContactModel {
	return c.modelMap["user_contact"].(model.UserContactModel)
}

func (c *Context) Init(config *conf.Config) {
	c.Context = &server.Context{}
	c.Context.Init(config)
	c.modelMap = loader.LoadModels(c.Config().Models, c.Database(), c.Logger(), c.SnowflakeNode())
	err := loader.LoadTables(c.Config().Models, c.Database())
	if err != nil {
		panic(err)
	}
}
