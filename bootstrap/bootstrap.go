package bootstrap

import (
	"github.com/JA-Coding-J/wechatbot/handlers"
	"github.com/eatmoreapple/openwechat"
	// "time"
	// "strings"
	// "os"
	"log"
)


func Run() {
	//bot := openwechat.DefaultBot()
	bot := openwechat.DefaultBot(openwechat.Desktop) // 桌面模式，上面登录不上的可以尝试切换这种模式

	// 注册消息处理函数
	bot.MessageHandler = handlers.Handler

	// 注册登陆二维码回调
	bot.UUIDCallback = handlers.QrCodeCallBack

	// 生成唯一 storage_[time].json 避免冲突
	// timeString := strings.Join(strings.Fields(time.Now().Format("2006/01/02 15:04")), "-")
	// filename := "storage.json"
	// dirname := "storage"
	// r, err1 := os.Stat(filename)
	// if err1 != nil {	// 文件已存在，移动到 storage 目录下
	//	log.Printf("文件是否存在, %v, %v", r, err1)
	// 	_, err2 := os.Stat(dirname)
	// 	if err2 == nil {	//	 目录不存在，创建目录
	// 		err3 := os.Mkdir(dirname, 0750)
	// 		log.Printf("Make directory fail: %v", err3)
	// 	}
	// 	targetPath := "./" + dirname + "/" + timeString + ".json"
	// 	os.Rename(filename, targetPath)
	// 	_, err4 := os.Stat(targetPath)
	// 	if err4 != nil {
	// 		log.Printf("文件移动成功!")
	// 	} else {
	// 		log.Printf("文件移动失败, %v", err4)
	// 	}
	// }

	// 创建热存储容器对象
	// reloadStorage := openwechat.NewJsonFileHotReloadStorage(filename)

	// 执行热登录
	// err := bot.HotLogin(reloadStorage)

	// if err != nil {
	// log.Printf("HotLogin error: %v \n", err)
	err := bot.Login()
	if err != nil {
		log.Printf("login error: %v \n", err)
		return
	}
	// }
	// 阻塞主goroutine, 直到发生异常或者用户主动退出
	bot.Block()
}
