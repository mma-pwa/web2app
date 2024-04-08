package utils

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/image/draw"
	"image"
	"image/jpeg"
	"image/png"
	"io"
	"math/rand"
	"net"
	"os"
	"reflect"
	"regexp"
	"strings"
	"time"
	"web2app/global"
)

func CopyFile(src, dst string) (err error) {
	// 打开源文件
	in, err := os.Open(src)
	if err != nil {
		return
	}
	defer in.Close()

	// 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		return
	}
	defer out.Close()

	// 执行复制操作
	_, err = io.Copy(out, in)
	if err != nil {
		return
	}

	// 关闭文件可能会产生错误
	err = out.Close()
	if err != nil {
		return
	}

	return
}

func Base64Decode(str string) string {
	//var dst []byte

	dst, err := base64.StdEncoding.DecodeString(str)
	if err != nil {
		global.GVA_LOG.Error(err.Error())
		return ""
	}
	return string(dst)
}
func IsPNGImage(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = png.Decode(file)
	if err != nil {
		return false
	}
	return true
}

func IsJpgImage(filename string) bool {
	file, err := os.Open(filename)
	if err != nil {
		return false
	}
	defer file.Close()

	_, err = jpeg.Decode(file)
	if err != nil {
		return false
	}

	return true
}
func PasswordHash(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	return string(bytes), err
}
func ValidateEmail(email string) bool {
	re := regexp.MustCompile(`^[a-zA-Z0-9.!#$%&'*+/=?^_` + "`" + `{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$`)
	return re.MatchString(email)
}
func PasswordVerify(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil

}

func RandCreator(l int) string {
	str := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	strList := []byte(str)

	result := []byte{}
	i := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i < l {
		new := strList[r.Intn(len(strList))]
		result = append(result, new)
		i = i + 1
	}
	return string(result)
}

// Base64Encode base64加密
func Base64Encode(src []byte) []byte {
	base64Table := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/"
	var coder = base64.NewEncoding(base64Table)
	return []byte(coder.EncodeToString(src))
}
func RemoveRepByMap(slc []string) []string {
	result := []string{}         //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

// 从Gin中获取真实的IP地址
func GetRealIPFromGin(c *gin.Context) string {
	xff := c.GetHeader("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xff, `,`)[0])
	if ip == "" {
		ip = strings.TrimSpace(c.GetHeader("X-Real-Ip"))
		if ip == "" {
			ip = c.ClientIP()
		}
	}
	return ip
}

func IsIp(ip string) bool {
	ipAddr := net.ParseIP(ip)
	if ipAddr == nil {
		return false
	}
	return true
}

// 去除重复字符串和空格Ip限定
func removeDuplicates(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		if !IsIp(string(a[i])) {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

func removeDuplicatesDns(a []string) (ret []string) {
	a_len := len(a)
	for i := 0; i < a_len; i++ {
		if (i > 0 && a[i-1] == a[i]) || len(a[i]) == 0 {
			continue
		}
		ret = append(ret, a[i])
	}
	return
}

// IsNil 判断是否为空指针
func IsNil(i interface{}) bool {
	vi := reflect.ValueOf(i)
	if vi.Kind() == reflect.Ptr {
		return vi.IsNil()
	}
	return false
}

func IsValidDomain(domain string) bool {
	domainRegex := regexp.MustCompile(`^[a-zA-Zd-]+(.[a-zA-Zd-]+)*.[a-zA-Z]{2,}$`)
	return domainRegex.MatchString(domain)
}

func ImgScale(sourceImg, dstImg string, width, height int) error {
	// 打开原始图片
	srcFile, err := os.Open(sourceImg)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	srcImg, err := png.Decode(srcFile)
	if err != nil {
		return err
	}

	// 计算缩放后的尺寸
	bounds := srcImg.Bounds()
	if bounds.Max.X <= 512 || bounds.Max.Y <= 512 {
		//return errors.New("上传的图片大小必须大于或者等于512X512")
	}
	// 创建一个新的图层用于缩放
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	// 进行缩放
	draw.BiLinear.Scale(m, m.Bounds(), srcImg, srcImg.Bounds(), draw.Over, nil)

	// 将缩放后的图片保存到文件
	dstFile, err := os.Create(dstImg)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	err = jpeg.Encode(dstFile, m, nil)
	if err != nil {
		return err
	}
	return nil
}
func GenerateCode() string {
	rand.Seed(time.Now().UnixNano())
	digits := "0123456789"
	code := ""
	for i := 0; i < 6; i++ {
		code += string(digits[rand.Intn(len(digits))])
	}
	return code
}
func ImgJpgScale(sourceImg, dstImg string, width, height int) error {
	// 打开原始图片
	srcFile, err := os.Open(sourceImg)
	if err != nil {
		return err
	}
	defer srcFile.Close()
	srcImg, err := jpeg.Decode(srcFile)
	if err != nil {
		return err
	}

	// 计算缩放后的尺寸
	//bounds := srcImg.Bounds()

	// 创建一个新的图层用于缩放
	m := image.NewRGBA(image.Rect(0, 0, width, height))

	// 进行缩放
	draw.BiLinear.Scale(m, m.Bounds(), srcImg, srcImg.Bounds(), draw.Over, nil)

	// 将缩放后的图片保存到文件
	dstFile, err := os.Create(dstImg)
	if err != nil {
		return err
	}
	defer dstFile.Close()

	err = jpeg.Encode(dstFile, m, nil)
	if err != nil {
		return err
	}
	return nil
}
