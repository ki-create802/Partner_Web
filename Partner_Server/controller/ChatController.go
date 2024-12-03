package controller

import (
	"Partner_Web/Partner_Server/common"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
)

type ChatController struct {
	DB *gorm.DB
}

type InChatController interface {
	//SearchChatList(c *gin.Context) //搜索返回匹配聊天室
	SearchUser(c *gin.Context) //搜索返回匹配用户
}

func NewChatController() InChatController {
	db := common.GetDB()
	//db.Table("user").AutoMigrate(model.User{}) 聊天室的数据结构还没整
	return ChatController{DB: db}
}

//func (b ChatController) SearchChatList(c *gin.Context) {
//	var queryString model.SearchString
//	c.ShouldBind(&queryString)
//
//	if queryString.SearchInfo != "" {//字符串不为空 匹配聊天室名称
//		var chatList []model.ChatInfo  //还没补充搜索完返回的聊天室数据类型
//		result := b.DB.Table("chatinfo").where("cname LIKE ? OR cremark LIKE ?","%"+queryString.SearchInfo+"%","%"+queryString.SearchInfo+"%").Find(& chatList)
//		if result.Error != nil {
//			common.Fail(c,,nil,"")
//		}
//
//		if(len()==0){
//			common.Success(c,nil,"未查询相关记录")
//		}
//
//	}
//	else{
//		common.Fail(c,422, nil, "搜索字符串为空")
//		return
//	}
//
//}

func (ChatController) SearchUser(c *gin.Context) {

}
