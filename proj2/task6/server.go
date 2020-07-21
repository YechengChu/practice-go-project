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
