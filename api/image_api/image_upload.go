package image_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/service/common/res"
	"gvd_server/utils/hash"
	"gvd_server/utils/jwts"
	"path"
	"strings"
	"time"
)

var ImageWhiteList = []string{
	"jpg",
	"png",
	"jpeg",
	"gif",
	"svg",
	"webp",
}

// ImageUploadView 上传图片
// @Tags 图片管理
// @Summary 上传图片
// @Description 上传图片
// @Param token header string true "token"
// @Param image formData file true "文件上传"
// @Router /api/image [post]
// @Accept multipart/form-data
// @Produce json
// @Success 200 {object} res.Response{}
func (ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		res.FailWithMsg("图片参数错误", c)
		return
	}
	_claims, _ := c.Get("claims")
	claims, _ := _claims.(*jwts.CustomClaims)
	savePath := path.Join("uploads", claims.NickName, fileHeader.Filename)
	// 白名单判断
	if !InImageWhiteList(fileHeader.Filename, ImageWhiteList) {
		res.FailWithMsg("文件非法", c)
		return
	}
	// 文件大小判断  2MB
	if fileHeader.Size > int64(2*1024*1024) {
		res.FailWithMsg("文件过大", c)
		return
	}
	// 计算文件hash
	file, _ := fileHeader.Open()
	fileHash := hash.FileMd5(file)
	// 对重复文件的判断
	var imageModel models.ImageModel
	err = global.DB.Take(&imageModel, "hash = ?", fileHash).Error
	// 没有 要上传，要入库
	// 有 只需要入库，但是入库的path需要改成和有的那个一样
	if err != nil {
		// 没有
		// 判断一下，数据库里面有没有这个路径的图片
		var count int64
		global.DB.Model(models.ImageModel{}).
			Where("path = ?", savePath).Count(&count)
		if count > 0 {
			// 存在重名的情况，那么这个时候就需要改一下文件名
			// 123.png   ->  123_1688054761.png
			// 12.png.png  ->  12.png_1688054761.png
			fileHeader.Filename = ReplaceFileName(fileHeader.Filename)
			savePath = path.Join("uploads", claims.NickName, fileHeader.Filename)
		}
		if err = c.SaveUploadedFile(fileHeader, savePath); err != nil {
			global.Log.Errorf("%s 文件保存错误 %s", savePath, err)
			res.FailWithMsg("上传图片错误", c)
			return
		}
	} else {
		// 有，修改入库的path
		savePath = imageModel.Path
	}
	// 使用这个hash对数据库里面记录的图片进行查询
	// zke        6dc...        uploads/zke/456.png
	// 李四        6dc...        uploads/zke/456.png
	// 用户删除图片的时候，发现有多个相同的hash，那就只删除记录
	imageModel = models.ImageModel{
		UserID:   claims.UserID,
		FileName: fileHeader.Filename,
		Size:     fileHeader.Size,
		Path:     savePath,
		Hash:     fileHash,
	}
	// 针对上传成功的图片写库
	if err = global.DB.Create(&imageModel).Error; err != nil {
		global.Log.Errorln(err)
		res.FailWithMsg("文件上传失败", c)
		return
	}
	res.OK(imageModel.WebPath(), "图片上传成功", c)
}

// InImageWhiteList 判断一个图片是否在白名单中
func InImageWhiteList(fileName string, whiteList []string) bool {
	// 截取文件后缀
	_list := strings.Split(fileName, ".") // xxx  1.2 xxx.png xxx.PNG  xxx.png   xxx.1.2.png  xxxx.png.exe
	if len(_list) < 2 {
		return false
	}
	suffix := strings.ToLower(_list[len(_list)-1])
	for _, s := range whiteList {
		if suffix == s {
			return true
		}
	}
	return false
}

// ReplaceFileName 修改文件名，加上时间戳
// tupian.png -> tupian_1688054761.png
// tupian.haokan.png -> tupian.haokan_1688054761.png
func ReplaceFileName(oldFileName string) string {
	// 123.png
	_list := strings.Split(oldFileName, ".")
	// [123   png] -> [123 _1688054761  png]
	lastIndex := len(_list) - 2
	var newList []string
	for i, s := range _list {
		if i == lastIndex {
			newList = append(newList, fmt.Sprintf("%s_%d", s, time.Now().Unix()))
			continue
		}
		newList = append(newList, s)
	}
	return strings.Join(newList, ".")
}
