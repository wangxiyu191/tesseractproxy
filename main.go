package main

import (
	"github.com/astaxie/beego"
	"github.com/otiai10/gosseract"
	"log"
	"os/exec"
	"strconv"
	"strings"
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

	err := this.SaveToFile("image", filePath)
	if err != nil {
		log.Println(err)
	}

	cmd := exec.Command("gm", "convert", filePath, filePath+".jpg")
	cmd.Run()
	if err != nil {
		log.Println(err)
	}

	out := gosseract.Must(gosseract.Params{Src: filePath + ".jpg"})
	out = strings.Replace(out, "\n", "", -1)
	out = strings.Replace(out, "\r", "", -1)
	out = strings.Replace(out, " ", "", -1)
	this.Data["json"] = map[string]string{"Result": out, "Id": fileName}
	this.ServeJSON()
}

func main() {
	//beego.BConfig.Listen.EnableAdmin = true
	//beego.BConfig.Listen.AdminPort = 3001
	beego.SetLogger("console", "")
	beego.Router("/", &MainController{})
	if beego.AppConfig.String("imagepath") != imagePath {
		imagePath = beego.AppConfig.String("imagepath")
	}
	beego.Run()
}
