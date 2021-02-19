package generate_sign_img

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"os"
	"time"
)

var fonts = []string{
	"jfcs.ttf",
	"yqk.ttf",
	"zql.ttf",
	"haku.ttf",
	//太模糊了, 注释掉
	"lfc.ttf",
	"bzcs.ttf",
	"qmt.ttf",
}

//name 签名名称
//out 生成文件路径,包含文件名与后缀
func GenerateSignImg(name string, out string) error {
	rand.Seed(time.Now().UnixNano())
	font := fonts[rand.Intn(len(fonts))]
	fmt.Println(font)
	resp, err := http.Get(fmt.Sprintf(
		"https://www.yishuzi.cn/qianming/image.png?fsize=50&font=%s&text=%s&mirror=no&color=111&vcolor=111&bgcolor=fff&alpha=no&output=png&spacing=4&shadow=no&transparent=no&icon=no&iconic=&top_spacing=5&left_spacing=6&icon_size=48",
		font,
		name,
	))
	if err != nil {
		return fmt.Errorf("请求链接失败: %s", err.Error())
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("读取数据失败: %s", err.Error())
	}
	if err = ioutil.WriteFile(out, bytes, os.ModePerm); err != nil {
		return fmt.Errorf("保存文件失败: %s", err.Error())
	}
	return nil
}
