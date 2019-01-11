package admin

import (
	"github.com/gin-gonic/gin"
	"webBlog/model"
	"net/http"
	"webBlog/helper"
	"strconv"
)

type Nav struct {
}

func (t *Nav) Index (c *gin.Context){
	navs := model.GetNavTree()
	c.HTML(http.StatusOK, "admin/nav/index.html",gin.H{
		"navs":navs,
	})
}
type NavData struct {
	Name string `form:"name" binding:"required"`
	Alias string `form:"alias" binding:"required"`
	Pid int `form:"pid"`
	Url string `form:"url" binding:"required"`
	Sort int `form:"sort"`
}
func (t *Nav) Add (c *gin.Context){
	if helper.IsGet(c) {
		//获取所有的顶级分类
		var navs []*model.Nav
		model.DB.Where("pid = ?", "0").Order("sort desc").Find(&navs)
		errorMsg := helper.GetFlash(c, "errorMsg")
		c.HTML(http.StatusOK, "admin/nav/add.html",gin.H{
			"errorMsg":errorMsg,
			"navs":navs,
		})
	} else if helper.IsPost(c) {
		var navData NavData
		if err := c.ShouldBind(&navData); err == nil{
			nav := model.Nav{
				Name:navData.Name,
				Alias:navData.Alias,
				Pid:navData.Pid,
				Url:navData.Url,
				Sort:navData.Sort,
			}
			err := model.DB.Create(&nav).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "添加成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "添加失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/navs/add")
	}
}

func (t *Nav) Edit (c *gin.Context){
	//获取当前
	id := c.Param("id")
	var nav model.Nav
	err := model.DB.Where("id = ?", id).First(&nav).Error
	if err != nil {
		c.Redirect(http.StatusFound, "/admin/nav/index")
		return
	}
	if helper.IsGet(c) {
		//获取所有的顶级分类
		var navs []*model.Nav
		model.DB.Where("pid = ?", "0").Order("sort desc").Find(&navs)
		errorMsg := helper.GetFlash(c, "errorMsg")

		c.HTML(http.StatusOK, "admin/nav/edit.html",gin.H{
			"errorMsg":errorMsg,
			"navs":navs,
			"nav":nav,
		})
	} else if helper.IsPost(c) {
		var navData NavData
		if err := c.ShouldBind(&navData); err == nil{
			nav.Name = navData.Name
			nav.Alias = navData.Alias
			nav.Url = navData.Url
			nav.Sort = navData.Sort
			nav.Pid = navData.Pid
			err := model.DB.Save(&nav).Error
			if err == nil {
				helper.SetFlash(c, "errorMsg", "修改成功 ！")
			} else {
				helper.SetFlash(c, "errorMsg", "修改失败 ！")
			}
		} else {
			helper.SetFlash(c, "errorMsg", "缺少参数 ！")
		}
		c.Redirect(http.StatusFound, "/admin/navs/edit/"+id)
	}
}

func (t *Nav) Del (c *gin.Context){
	id := c.Param("id")
	var nav model.Nav
	err := model.DB.Where("id = ?", id).First(&nav).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该分类，删除失败")
		return
	}
	model.DB.Delete(&nav)
	model.DB.Where("pid = ?", id).Delete(&model.Nav{})
	helper.ReturnJson(c, 1, "删除成功")
	return
}

func (t *Nav) ChangeOrder (c *gin.Context){
	id := c.PostForm("id")
	sort := c.PostForm("sort")
	if id =="" || sort == "" {
		helper.ReturnJson(c, 0, "参数错误")
		return
	}
	var nav model.Nav
	err := model.DB.Where("id = ?", id).First(&nav).Error
	if err != nil {
		helper.ReturnJson(c, 0, "无该分类，设置失败")
		return
	}
	nav.Sort, err = strconv.Atoi(sort)
	if err != nil {
		helper.ReturnJson(c, 0, "非法排序值")
		return
	}
	model.DB.Save(&nav)
	helper.ReturnJson(c, 1, "修改成功")
}