package bootstrapper

import (
	//"time"

	"TaskManager/dataSource"
	"TaskManager/middleWare/logrus"
	"github.com/kataras/iris"
	"github.com/kataras/iris/middleware/logger"
	"github.com/kataras/iris/middleware/recover"
	"log"
)

type Configurator func(*Bootstrapper) error

type Bootstrapper struct{
	*iris.Application
	AppName string
	AppOwner string
}

func New(appName string, ownerName string) (*Bootstrapper) {
	return &Bootstrapper{
		Application: iris.New(),
		AppName:     appName,
		AppOwner:    ownerName,
	}
}

func(this *Bootstrapper)BootStrap() *Bootstrapper {
	this.Favicon("./web/static/favicon.ico")
	this.UseGlobal(recover.New())
	this.UseGlobal(logger.New())

	return this
}

func(this *Bootstrapper)Configure(configuratorList ...Configurator) *Bootstrapper{
	for _, configurator := range configuratorList {
		err := configurator(this)
		if err != nil{
			logrus.Logger.Fatal("configurator fail: ", configurator)
		}
	}
	return this
}

func (this *Bootstrapper) SetView(viewDir string) *Bootstrapper{
	htmlEngine :=iris.HTML(viewDir, ".html")
	this.RegisterView(htmlEngine)
	return this
}

func (this *Bootstrapper) Listen (addr string) *Bootstrapper {
	err := this.Run(iris.Addr(addr))
	if err != nil {
		log.Fatal("Listen fail: ", err)
	}
	return this
}

func (this Bootstrapper) Start() {
	//创建日志
	logrus.NewLogger()

	//监听端口
	this.Listen(":80")
	//配置
	this.BootStrap()
	this.Configure()
	//Views
	this.SetView("./web/")
	err := dataSource.NewMysqlGroup()
	if err != nil {
		logrus.Logger.Fatal("init mysql fail :", err)
	}

}

