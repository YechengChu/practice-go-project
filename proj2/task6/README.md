# 上传彩色图片并以黑白展示
## 目标
建立一个http服务器，上传一个彩色图片，返回该彩色的黑白图片
## 代码
server.go
```go
package main

import (
	"errors"
	"fmt"
	"image"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"path"
	"strings"

	_ "image/jpeg"
	"image/png"
)

const html = `<html>
    <head></head>
    <body>
        <form method="post" enctype="multipart/form-data">
            Please upload an image:
            <input type="file" name="image" />
            <input type="submit" />
        </form>
    </body>
</html>`

func main() {
	http.HandleFunc("/upload/", uploadHandle)    // 上传
	http.HandleFunc("/uploaded/", showPicHandle) //显示图片
	err := http.ListenAndServe(":3000", nil)
	fmt.Println(err)
} // main

// 上传图像接口
func uploadHandle(w http.ResponseWriter, req *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	req.ParseForm()
	if req.Method != "POST" {
		w.Write([]byte(html))
	} else {
		// 接收图片
		uploadFile, handle, err := req.FormFile("image")
		errorHandle(err, w)

		// 检查图片后缀
		ext := strings.ToLower(path.Ext(handle.Filename))
		if ext != ".jpg" && ext != ".png" {
			errorHandle(errors.New("Only JPG or PNG is supported"), w)
			return
			//defer os.Exit(2)
		} // if

		// 保存图片
		os.Mkdir("./uploaded/", 0777)
		saveFile, err := os.OpenFile("./uploaded/"+handle.Filename, os.O_WRONLY|os.O_CREATE, 0666)
		errorHandle(err, w)
		// fmt.Printf("The type of saveFile is: %T\n", saveFile)
		// converter(handle.Filename)
		io.Copy(saveFile, uploadFile)
		converter(handle.Filename)
		defer uploadFile.Close()
		defer saveFile.Close()
		// 上传图片成功
		w.Write([]byte("See the uploaded image: <a target='_blank' href='/uploaded/" + handle.Filename + "'>" + handle.Filename + "</a> <br>" +
			"Image to greyscale is: <a target='_blank' href='/uploaded/gray_" + handle.Filename + "'>" + "gray_" + handle.Filename))
	} // if ... else
} // uploadHandle

// 显示图片接口
func showPicHandle(w http.ResponseWriter, req *http.Request) {
	file, err := os.Open("." + req.URL.Path)
	errorHandle(err, w)

	defer file.Close()
	buff, err := ioutil.ReadAll(file)
	errorHandle(err, w)
	w.Write(buff)
} // showPicHandle

// 统一错误输出接口
func errorHandle(err error, w http.ResponseWriter) {
	if err != nil {
		w.Write([]byte(err.Error()))
	} // if
} // errorHandle

// 将图片转成黑白并保存
func converter(input string) {
	// fmt.Println(input)
	inputName, err := os.Open("./uploaded/"+input)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	} // if
	// Decode image to JPEG
	img, _, err := image.Decode(inputName)
	if err != nil {
		// handle error
		log.Fatal(err)
	} // if
	// log.Printf("Image type: %T", img)

	// Converting image to grayscale
	grayImg := image.NewGray(img.Bounds())
	for y := img.Bounds().Min.Y; y < img.Bounds().Max.Y; y++ {
		for x := img.Bounds().Min.X; x < img.Bounds().Max.X; x++ {
			grayImg.Set(x, y, img.At(x, y))
		} // for
	} // for

	// Working with grayscale image, e.g. convert to png
	// And save to uploaded file
	f, err := os.Create("./uploaded/gray_"+input)
	if err != nil {
		// handle error
		log.Fatal(err)
	} // if
	defer f.Close()

	if err := png.Encode(f, grayImg); err != nil {
		log.Fatal(err)
	} // if
} // converter
```
## 运行方法
在命令行中运行server
```
$ go run server.go
```
然后在浏览器中打开网站
```
localhost:3000/upload
```
## 运行结果展示
![Screen Shot 2020-07-21 at 19.16.41.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330207709-63d21465-82c6-4cd5-9296-f3732542b7ee.png#align=left&display=inline&height=267&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.16.41.png&originHeight=267&originWidth=591&size=32764&status=done&style=none&width=591)

在task6文件夹中打开命令行，比如使用iTerm
![Screen Shot 2020-07-21 at 19.18.36.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330323323-274ec78f-09c1-4a76-8382-ce626be425df.png#align=left&display=inline&height=112&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.18.36.png&originHeight=112&originWidth=732&size=19402&status=done&style=none&width=732)

运行server
![Screen Shot 2020-07-21 at 19.19.12.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330359557-07a8ac93-1530-429a-ab5e-a12c9829dbaf.png#align=left&display=inline&height=326&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.19.12.png&originHeight=326&originWidth=1319&size=38504&status=done&style=none&width=1319)

使用浏览器打开网页
![Screen Shot 2020-07-21 at 19.20.22.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330432946-cad16f38-82ab-4641-b316-83d0a556880f.png#align=left&display=inline&height=489&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.20.22.png&originHeight=489&originWidth=958&size=121292&status=done&style=none&width=958)

选择图片上传，并提交

![Screen Shot 2020-07-21 at 19.25.00.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330707133-21b78e7a-61ce-42af-85ab-eea068cdcee5.png#align=left&display=inline&height=253&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.25.00.png&originHeight=253&originWidth=615&size=51197&status=done&style=none&width=615)

在用户点击提交按钮后，会在server运行的文件夹内自动创建uploaded文件夹，保存用户上传的图片和转换成的黑白图片
![Screen Shot 2020-07-21 at 19.20.57.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330464236-9c6a14a1-1c7c-49f1-a263-dec48c08d6ad.png#align=left&display=inline&height=151&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.20.57.png&originHeight=151&originWidth=1312&size=35178&status=done&style=none&width=1312)

网页会显示2个链接，一个是原图链接，一个是黑白图片的链接
![image.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330598780-8419d63b-44b8-426b-9187-8676e0ab90c7.png#align=left&display=inline&height=675&margin=%5Bobject%20Object%5D&name=image.png&originHeight=675&originWidth=1316&size=827027&status=done&style=none&width=1316)
![Screen Shot 2020-07-21 at 19.23.17.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330612534-261df0fa-031d-414a-91d1-44f413260894.png#align=left&display=inline&height=683&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.23.17.png&originHeight=683&originWidth=1315&size=372857&status=done&style=none&width=1315)
若用户提交的图片不是PNG或者JPG类型，网站会报错

![Screen Shot 2020-07-21 at 19.28.27.png](https://cdn.nlark.com/yuque/0/2020/png/1609946/1595330914095-3dec6cb1-fdb1-423f-a0ef-7503d034fbba.png#align=left&display=inline&height=116&margin=%5Bobject%20Object%5D&name=Screen%20Shot%202020-07-21%20at%2019.28.27.png&originHeight=116&originWidth=943&size=21962&status=done&style=none&width=943)

---

## 参考资料
用go来搭建一个简单的图片上传网站
[https://blog.csdn.net/stpeace/article/details/82716145](https://blog.csdn.net/stpeace/article/details/82716145)

golang 上传，下载图片
[https://blog.csdn.net/BlackCarDriver/article/details/87870109](https://blog.csdn.net/BlackCarDriver/article/details/87870109)

Convert color image to grayscale
[https://riptutorial.com/go/example/31693/convert-color-image-to-grayscale](https://riptutorial.com/go/example/31693/convert-color-image-to-grayscale)

golang的图片操作：缩放图片+合成图片
[https://blog.csdn.net/mirage003/article/details/88084303](https://blog.csdn.net/mirage003/article/details/88084303)
