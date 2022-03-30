package example

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"sync"
)

// HandleError 错误处理
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}

// 定义一些全局变量
var (
	// 需要爬取的页数
	pages int
	// ImageUrls 存放图片链接的数据管道
	ImageUrls chan string
	wg        sync.WaitGroup
	// 用于监控协程
	chanTask chan string
	// 定义正则匹配图片链接的规则
	reImg = `https://uploadfile.bizhizu.cn[^"]+?(\.jpg)`
)

func mains() {
	fmt.Println("请输入需要爬取的页数:")
	fmt.Scanf("%d", &pages)
	// 初始化管道
	ImageUrls = make(chan string, 1000000)
	// 根据爬取的页数设置协程任务数
	chanTask = make(chan string, pages)
	// 协程循环爬取网页
	for i := 1; i < pages+1; i++ {
		wg.Add(1)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E7%BE%8E%E5%A5%B3/" + strconv.Itoa(i) + ".html")
	}
	// 任务统计协程，统计任务是否都完成，完成则关闭管道
	wg.Add(1)
	go CheckOK()
	// 下载协程：从管道中读取链接并下载
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}

// 爬图片链接到管道
// url是传的整页链接
func getImgUrls(url string) {
	urls := getUrls(url)
	// 遍历切片里所有链接，存入数据管道
	for _, url := range urls {
		ImageUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	wg.Done()
}

// 获取当前页图片链接
func getUrls(url string) (urls []string) {
	pageStr := GetHTML(url)
	// 正则匹配网页所有的图片urls
	re := regexp.MustCompile(reImg)
	results := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(results))

	// 由于re.FindAllStringSubmatch的特性，返回的不是纯粹的urls，所以重新提取赋值
	for _, result := range results {
		url := result[0]
		urls = append(urls, url)
	}
	return urls
}

// GetHTML 抽取根据url获取内容
func GetHTML(url string) (html string) {
	resp, err := http.Get(url)
	HandleError(err, "http.Get url")
	defer resp.Body.Close()
	// 2.读取页面内容
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "ioutil.ReadAll")
	// 字节转字符串
	html = string(pageBytes)
	return html
}

// CheckOK 任务统计协程
// 在传入每页图片urls添加这页的url到chanTask，然后检查的时候读取，无法读取数据时检查页数与处理数是否相等
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s 完成了爬取任务\n", url)
		count++
		if count == pages {
			close(ImageUrls)
			break
		}
	}
	wg.Done()
}

// DownloadImg 下载图片
func DownloadImg() {
	for url := range ImageUrls {
		filename := GetFilenameFromUrl(url)
		ok := DownloadFile(url, filename)
		if ok {
			fmt.Printf("%s 下载成功\n", filename)
		} else {
			fmt.Printf("%s 下载失败\n", filename)
		}
	}
	wg.Done()
}

// DownloadFile 下载图片，传入图片名称
func DownloadFile(url string, filename string) (ok bool) {
	resp, err := http.Get(url)
	//如果获取网页失败就抛出错误
	HandleError(err, "http.get.url")

	//开始读取网页数据
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "resp.body")

	//判断是否存在images文件夹,如果不存在就创建该文件夹
	path := "./images"
	if _, err := os.Stat(path); err != nil {
		fmt.Println("path not exists ", path)
		err := os.MkdirAll(path, 0711)
		if err != nil {
			log.Println("Error creating directory")
			log.Println(err)
			return
		}
	}
	filename = "./images/" + filename
	// 写出数据
	err = ioutil.WriteFile(filename, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}

// GetFilenameFromUrl 截取url名字
func GetFilenameFromUrl(url string) (filename string) {
	// 返回最后一个/的位置
	lastIndex := strings.LastIndex(url, "/")
	// 切出来
	filename = url[lastIndex+1:]
	return filename
}
