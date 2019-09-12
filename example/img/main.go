package main

import (
	"bytes"
	"fmt"
	"github.com/nfnt/resize"
	"image/png"
	"os"
)

func main() {
	img2txt("/Users/wangzhongjie/Desktop/Code/github/learning-practice-go/example/img/go.png", 200, []string{"@", "#", "*", "%", "+", ",", ".", " "}, "\n", "./保存的文本.txt")
}
/*先定义一个函数
参数：
  imgPath: 图片路径
  size: 生成文本后的尺寸(这个不是真实的尺寸，1代表1个像素，1个像素会被替换成1个字符，所以是字符的个数，高度是自动换算的，所以这里的size指的是“宽度”被压缩成多少像素)
  txts: 将像素处理成的字符列表
  rowend: 换行字符（因为windows和linux不同）
  output: 生成文本文件保存路径
*/
func img2txt(imgPath string, size uint, txts []string, rowend string, output string) {
	//获取图片文件
	file, err := os.Open(imgPath)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer file.Close()

	//用图片文件获取图片对象
	img, err := png.Decode(file)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	//用将宽度设置为size，然后换算出等比例的高度
	var width = size
	var height = (size * uint(img.Bounds().Dy())) / (uint(img.Bounds().Dx()))
	height = height * 6 / 10 //这里6/10是大致字符的宽高比
	newimg := resize.Resize(width, height, img, resize.Lanczos3)  //根据高宽resize图片，并得到新图片的像素值
	dx := newimg.Bounds().Dx()
	dy := newimg.Bounds().Dy()

	//创建一个字节buffer，一会用来保存字符
	textBuffer := bytes.Buffer{}

	//遍历图片每一行每一列像素
	for y := 0; y < dy; y++ {
		for x := 0; x < dx; x++ {
			colorRgb := newimg.At(x, y)
			r, g, b, _ := colorRgb.RGBA()

			//获得三原色的值，算一个平均数出来
			avg := uint8((r + g + b) / 3 >> 8)
			//有多少个用来替换的字符就将256分为多少个等分，然后计算这个像素的平均值趋紧与哪个字符，最后，将这个字符添加到字符buffer里
			num := avg / uint8(256/len(txts))
			textBuffer.WriteString(txts[num])
			fmt.Print(txts[num]) //打印出来
		}

		textBuffer.WriteString(rowend)  //一行结束，换行
		fmt.Print(rowend)
	}

	//将字符buffer的数据写入到文本文件里，结束。
	f, err := os.Create(output + ".txt")
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer f.Close()

	f.WriteString(textBuffer.String())
}