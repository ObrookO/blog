package controllers

import (
	"blog/models"
	"blog/tool"

	"github.com/ObrookO/go-utils"
)

type ToolController struct {
	BaseController
}

func (c *ToolController) Prepare() {
	c.BaseController.Prepare()
	c.Data["showTB"] = false
}

// Md5 MD5加/解密页面
func (c *ToolController) Md5() {
	c.Layout = "layouts/master.html"
	c.TplName = "tool/md5.html"
	c.LayoutSections = map[string]string{
		"Script": "tool/md5_script.html",
	}
}

// CalculateMd5 MD5加/解密
func (c *ToolController) CalculateMd5() {
	raw := c.GetString("rawData")
	opt, _ := c.GetInt("opt")

	if tool.GetCharAmount(raw) == 0 {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "请输入原始数据"}
		c.ServeJSON()
		return
	}

	if !utils.ObjInIntSlice(opt, []int{1, 2}) {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "请选择正确的操作类型"}
		c.ServeJSON()
		return
	}

	// 加密
	if opt == 1 {
		// 插入记录
		encrypted := utils.Md5Str(raw)
		if !models.IsMd5RecordExists(map[string]interface{}{"raw_data": raw}) {
			models.AddMd5Record(models.Md5{
				RawData:     raw,
				EncryptData: encrypted,
			})
		}

		c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK", Data: encrypted}
		c.ServeJSON()
		return
	}

	// 解密
	record, _ := models.GetOneMd5Record(map[string]interface{}{"encrypt_data": raw})
	if record.Id == 0 {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "解密失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK", Data: record.RawData}
	c.ServeJSON()
	return
}

// Aes AES加/解密页面
func (c *ToolController) Aes() {
	c.Layout = "layouts/master.html"
	c.TplName = "tool/aes.html"
	c.LayoutSections = map[string]string{
		"Script": "tool/aes_script.html",
	}
}

// CalculateAes AES加解密
func (c *ToolController) CalculateAes() {
	defer func() {
		if r := recover(); r != nil {
			c.Data["json"] = &JSONResponse{Code: 500000, Msg: "操作失败，请检查输入是否正确"}
			c.ServeJSON()
		}
	}()

	raw := c.GetString("rawData")
	aes := c.GetString("aesKey")
	opt, _ := c.GetInt("opt")

	if tool.GetCharAmount(raw) == 0 {
		c.Data["json"] = &JSONResponse{Code: 400000, Msg: "请输入原始数据"}
		c.ServeJSON()
		return
	}

	if !utils.ObjInIntSlice(len([]byte(aes)), []int{16, 24, 32}) {
		c.Data["json"] = &JSONResponse{Code: 400001, Msg: "AES KEY必须为16位、24位、32位"}
		c.ServeJSON()
		return
	}

	if !utils.ObjInIntSlice(opt, []int{1, 2}) {
		c.Data["json"] = &JSONResponse{Code: 400002, Msg: "请选择正确的操作类型"}
		c.ServeJSON()
		return
	}

	// 加密
	if opt == 1 {
		encrypted, err := utils.AesEncrypt(raw, aes)
		if err != nil {
			c.Data["json"] = &JSONResponse{Code: 400003, Msg: "加密失败"}
			c.ServeJSON()
			return
		}

		c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK", Data: encrypted}
		c.ServeJSON()
		return
	}

	// 解密
	decrypted, err := utils.AesDecrypt(raw, aes)
	if err != nil {
		c.Data["json"] = &JSONResponse{Code: 400004, Msg: "解密失败"}
		c.ServeJSON()
		return
	}

	c.Data["json"] = &JSONResponse{Code: 200, Msg: "OK", Data: decrypted}
	c.ServeJSON()
	return
}

// Base64 Base64转码页面
func (c *ToolController) Base64() {
	c.Layout = "layouts/master.html"
	c.TplName = "tool/base64.html"
	c.LayoutSections = map[string]string{
		"Script": "tool/base64_script.html",
	}
}

// RandomStr 生成随机字符串页面
func (c *ToolController) RandomStr() {
	c.Layout = "layouts/master.html"
	c.TplName = "tool/random.html"
	c.LayoutSections = map[string]string{
		"Script": "tool/random_script.html",
	}
}

// StrLength 计算字符串长度页面
func (c *ToolController) StrLength() {
	c.Layout = "layouts/master.html"
	c.TplName = "tool/length.html"
	c.LayoutSections = map[string]string{
		"Script": "tool/length_script.html",
	}
}
