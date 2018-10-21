package main

import (
	"goserver/webs"
	"goserver/xml"
	"goserver"
	"goserver/gpa"
	"github.com/cihub/seelog"
)

type Module struct {
}

func (app *Module) Include(ele *xml.Element) {
	f := getFile(ele.Value)
	seelog.Info("include file:", f)
	dom, err := xml.LoadByFile(f)
	if err == nil {
		InvokeByXml(dom)
	} else {
		seelog.Error(err)
		panic("配置文件出错:" + f)
	}
}

func (app *Module) Parameter(ele *xml.Element) {
	ps := ele.AllNodes()
	for _, p := range ps {
		utils.EnvParamSet(p.Name(), p.Value)
	}
}

func (app *Module) Gpa(ele *xml.Element) {
	dsn, b := ele.AttrValue("DbDsn")
	if b && len(dsn) > 0 {
		db := gpa.InitGpa(dsn)
		put(ele, db)
	}
}

func (app *Module) Verify(ele *xml.Element) {
	vb := webs.NewVerify(ele, getGpa(ele), putFunRun)
	put(ele, vb)
	tfn, ext := ele.AttrValue("TplFunName")
	if ext {
		webs.RegisterBaseFun(tfn, vb)
		seelog.Info("添加模版函数:", tfn)
	}
}
