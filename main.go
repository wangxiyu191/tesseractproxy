package main

import (
	"github.com/astaxie/beego"
	"github.com/otiai10/gosseract"
	"log"
	"os/exec"
	"strconv"
	"time"
)

var (
	imagePath = "/tmp/"
)

type MainController struct {
	beego.Controller
}

func (this *MainController) Post() {
	fileName := strconv.Itoa(int(time.Now().UnixNano()))
	filePath := imagePath + fileName //无扩展名的文件路径
	//log.Println(filePath)

	err := this.SaveToFile("image", filePath+".gif")
	if err != nil {
		log.Println(err)
	}

	cmd := exec.Command("convert", filePath+".gif", filePath+".jpg")
	cmd.Run()
	if err != nil {
		log.Println(err)
	}

	out := gosseract.Must(gosseract.Params{Src: filePath + ".jpg"})
	this.Data["json"] = map[string]string{"Result": out[0:4], "Id": fileName}
	this.ServeJson()
}

func main() {
	//beego.EnableAdmin = true
	//beego.AdminHttpPort = 3001
	beego.SetLogger("console", "")
	beego.Router("/", &MainController{})
	if beego.AppConfig.String("imagepath") != imagePath {
		imagePath = beego.AppConfig.String("imagepath")
	}
	beego.Run()
}
