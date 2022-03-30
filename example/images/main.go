package main

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

var (
	// 页码数
	page int
	// 管道
	ImagesUrls chan string
	//	协程数
	chanTask chan string
	wg       sync.WaitGroup
	// 定义正则匹配图片链接的规则
	reImg = `https://uploadfile.bizhizu.cn[^"]+?(\.jpg)`
)

// HandleError 错误处理
func HandleError(err error, why string) {
	if err != nil {
		fmt.Println(why, err)
	}
}
func main() {
	fmt.Println("请输入需要爬取的页数")
	_, err := fmt.Scanf("%d", &page)
	if err != nil {
		return
	}
	//	初始化管道
	ImagesUrls = make(chan string, 1000000)
	//	初始化协程数
	chanTask = make(chan string, page)
	//	协程循环爬取
	for i := 1; i < page+1; i++ {
		wg.Add(1)
		str := strconv.Itoa(i)
		go getImgUrls("https://www.bizhizu.cn/shouji/tag-%E7%BE%8E%E5%A5%B3/" + str + ".html")
	}
	wg.Add(1)
	go CheckOK()
	for i := 0; i < 7; i++ {
		wg.Add(1)
		go DownloadImg()
	}
	wg.Wait()
}

// CheckOK 协程统计
func CheckOK() {
	var count int
	for {
		url := <-chanTask
		fmt.Printf("%s完成了爬取任务\n", url)
		count++
		if count == page {
			fmt.Println("所有任务完成")
			close(ImagesUrls)
			break
		}
	}
	wg.Done()
}

/**
 * @url 获取图片链接
 */
func getImgUrls(url string) {
	urls := getUrls(url)
	// 遍历切片里所有链接，存入数据管道
	for _, url := range urls {
		ImagesUrls <- url
	}
	// 标识当前协程完成
	// 每完成一个任务，写一条数据
	// 用于监控协程知道已经完成了几个任务
	chanTask <- url
	wg.Done()
}

//获取当前页面的所有图片链接
func getUrls(url string) (urls []string) {
	pageStr := GetHTML(url)
	//正则匹配网页字符串中所有图片urls
	re := regexp.MustCompile(reImg)
	result := re.FindAllStringSubmatch(pageStr, -1)
	fmt.Printf("共找到%d条结果\n", len(result))
	//重新提取值
	for _, item := range result {
		url := item[0]
		urls = append(urls, url)
	}
	return urls
}

// GetHTML 根据url获取html
func GetHTML(url string) (html string) {
	resp, err := http.Get(url)
	HandleError(err, "请求url失败")
	defer resp.Body.Close()
	pageBytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "读取页面失败")
	html = string(pageBytes)
	return html
}

// DownloadImg 下载图片
func DownloadImg() {
	for url := range ImagesUrls {
		fileName := GetFilenameFromUrl(url)
		fmt.Println("获取到的文件名:" + fileName)
		ok := DownloadFile(url, fileName)
		if ok {
			fmt.Printf("--------------------%s下载成功--------------------\n", fileName)
		} else {
			fmt.Printf("--------------------%s下载失败--------------------\n\n", fileName)
		}
	}
	wg.Done()
}

// GetFilenameFromUrl 根据url获取图片名称
func GetFilenameFromUrl(url string) (fileName string) {
	lastIndex := strings.LastIndex(url, "/")
	fileName = url[lastIndex+1:]
	return fileName
}

// DownloadFile 下载图片
func DownloadFile(url string, fileName string) (ok bool) {
	resp, err := http.Get(url)
	HandleError(err, "图片url请求失败")

	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	HandleError(err, "读取图片失败")
	path := "./images"
	if _, err := os.Stat(path); err != nil {
		fmt.Println("文件夹不存在，创建文件夹")
		err := os.MkdirAll(path, 0711)
		if err != nil {
			log.Println("创建文件夹失败")
			log.Println(err)
			return
		}
	}
	fileName = path + "/" + fileName
	//写入数据
	err = ioutil.WriteFile(fileName, bytes, 0666)
	if err != nil {
		return false
	} else {
		return true
	}
}
