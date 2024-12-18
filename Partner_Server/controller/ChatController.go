package controller

import (
	"Partner_Web/Partner_Server/common"
	"Partner_Web/Partner_Server/model"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	//"golang.org/x/crypto/bcrypt"
)

type ChatController struct {
	DB *gorm.DB
}

type InChatController interface {
	SearchChatList(c *gin.Context)  //搜索返回匹配聊天室
	SearchUser(c *gin.Context)      //搜索返回匹配用户
	DisbandChatRoom(c *gin.Context) //房主解散群聊
	LeaveChatRoom(c *gin.Context)   //退出群聊
	AddMember(c *gin.Context)       //添加成员到uc_match表中
	GetSpeakers(c *gin.Context)     //返回发言成员列表
}

func NewChatController() InChatController {
	db := common.GetDB()
	//db.Table("user").AutoMigrate(model.User{})
	return ChatController{DB: db}
}

func (b ChatController) SearchChatList(c *gin.Context) { //未测试
	var queryString model.SearchString
	c.Bind(&queryString)

	//获取板块信息
	var IsAll bool
	var blockID int

	IsAll = true
	if queryString.Block != "全部" { //判断是否选择了版块，版块
		IsAll = false //不是全部板块

		var BID_Info struct { //定义一个结构体接收
			BID int `gorm:"column:bid"`
		}
		result := b.DB.Table("block").
			Select("bid").Where("bname = ?", queryString.Block).
			First(&BID_Info)
		if result.Error != nil {
			common.Fail(c, 500, nil, "数据库查询板块id错误")
			return
		}
		blockID = BID_Info.BID
		fmt.Println("id:   ", blockID)
		fmt.Println("block_name: ", queryString.Block)
	}

	//获取聊天室信息
	var chatList []model.ChatInfo     //聊天室数据
	if queryString.SearchInfo != "" { //字符串不为空 匹配聊天室名称
		if IsAll { //筛选全部
			result := b.DB.Table("chatinfo").
				Where("cname LIKE ? OR cremark LIKE ?", "%"+queryString.SearchInfo+"%", "%"+queryString.SearchInfo+"%").
				Find(&chatList)
			if result.Error != nil {
				common.Fail(c, 500, nil, "搜索词数据库查询错误")
				return
			}
		} else { //多一个版块筛选
			result := b.DB.Table("chatinfo").
				Where("bid = ?", blockID).
				Where("cname LIKE ? OR cremark LIKE ?", "%"+queryString.SearchInfo+"%", "%"+queryString.SearchInfo+"%").Find(&chatList)
			if result.Error != nil {
				common.Fail(c, 500, nil, "版块搜索词数据库查询错误")
				return
			}
		}
	} else { //为空就是全部推送，再判断block
		if IsAll { //全部推送
			result := b.DB.Table("chatinfo").Find(&chatList)
			if result.Error != nil {
				common.Fail(c, 500, nil, "全部聊天室数据库查询错误")
				return
			}
		} else { //推送对应板块所有记录
			result := b.DB.Table("chatinfo").Where("bid = ?", blockID).Find(&chatList)
			if result.Error != nil {
				common.Fail(c, 500, nil, "版块数据库查询错误")
				return
			}
		}
	}

	if len(chatList) == 0 {
		common.Success(c, nil, "未查询到相关记录")
		return
	}

	var searchResults []map[string]interface{}
	for _, chat := range chatList {
		//user里查找每个RoomOwnerID 的头像 姓名
		var roomOwner model.User
		result := b.DB.Table("user").
			Where("uid = ?", chat.RoomOwnerID).
			First(&roomOwner)
		if result.Error != nil {
			common.Fail(c, 500, nil, "搜索房主数据库查询错误")
			return
		}

		// 查询 success_match 表，获取成员列表
		var successMatches []struct {
			UID uint `gorm:"column:uid"`
		}
		result = b.DB.Table("success_match").
			Select("uid").
			Where("cid = ?", chat.RoomID).
			Find(&successMatches)
		if result.Error != nil {
			common.Fail(c, 500, nil, "搜索匹配用户数据库查询错误")
			return
		}

		// 构建成员列表
		var memberList []map[string]interface{}
		for _, match := range successMatches {
			// 查询 user 表，获取成员信息
			var user model.User
			result = b.DB.Table("user").
				Where("uid = ?", match.UID).
				First(&user)
			if result.Error != nil {
				common.Fail(c, 500, nil, "查询成员信息数据库查询错误")
				return
			}

			// 构建成员信息
			member := map[string]interface{}{
				"memberID":  user.UID,
				"memberImg": user.UImage,
			}
			memberList = append(memberList, member)
		}

		// 构建聊天室信息
		room := map[string]interface{}{
			"roomID":        chat.RoomID,
			"roomName":      chat.RoomName,
			"roomIntro":     chat.RoomIntro,
			"roomOwnerID":   chat.RoomOwnerID,
			"roomOwnerName": roomOwner.UName,
			"roomOwnerImg":  roomOwner.UImage,
			"memberList":    memberList,
		}
		searchResults = append(searchResults, room)
	}
	common.Success(c, gin.H{"searchResults": searchResults}, "成功返回聊天室列表")
}

func (b ChatController) SearchUser(c *gin.Context) {
	var queryString model.SearchString
	c.Bind(&queryString)

	fmt.Println("query:", queryString.SearchInfo)

	var userList []model.User
	if queryString.SearchInfo != "" { //搜索词不为空
		result := b.DB.Table("user").Select("uid,uname,uremark,uimage").
			Where("uname LIKE ? OR uid LIKE ?", "%"+queryString.SearchInfo+"%", "%"+queryString.SearchInfo+"%").
			Find(&userList)
		if result.Error != nil {
			common.Fail(c, 500, nil, "用户查询数据库错误")
			return
		}
	} else {
		common.Fail(c, 400, nil, "请输入相关用户关键词")
		return
	}

	common.Success(c, gin.H{"userList": userList}, "成功返回用户列表")
}

// 房主解散群聊
func (b ChatController) DisbandChatRoom(c *gin.Context) {
	var request struct {
		RoomID int `json:"roomID"`
	}
	c.Bind(&request)

	roomID := request.RoomID

	// 更新房间状态为 0，表示房间被删除
	result := b.DB.Table("chatinfo").Where("cid = ?", roomID).Update("cstate", 0)
	if result.Error != nil {
		common.Fail(c, 500, nil, "更新房间状态失败")
		return
	}

	// 从 success_match 表中删除所有成员
	result = b.DB.Table("success_match").Where("cid = ?", roomID).Delete(&model.SuccessMatch{})
	if result.Error != nil {
		common.Fail(c, 500, nil, "从 success_match 表中删除所有成员失败")
		return
	}

	// 从 uc_match 表中删除所有成员
	result = b.DB.Table("uc_match").Where("cid = ?", roomID).Delete(&model.UcMatch{})
	if result.Error != nil {
		common.Fail(c, 500, nil, "从 uc_match 表中删除所有成员失败")
		return
	}

	common.Success(c, nil, "房间已成功解散")
}

// 成员退出群聊
func (b ChatController) LeaveChatRoom(c *gin.Context) {
	var request struct {
		RoomID int `json:"roomID"`
		UserID int `json:"userID"`
	}
	c.Bind(&request)

	roomID := request.RoomID
	userID := request.UserID

	// 从 success_match 表中删除用户
	result := b.DB.Table("success_match").Where("cid = ? AND uid = ?", roomID, userID).Delete(&model.SuccessMatch{})
	if result.Error != nil {
		common.Fail(c, 500, nil, "从 success_match 表中删除用户失败")
		return
	}

	// 从 uc_match 表中删除用户
	result = b.DB.Table("uc_match").Where("cid = ? AND uid = ?", roomID, userID).Delete(&model.UcMatch{})
	if result.Error != nil {
		common.Fail(c, 500, nil, "从 uc_match 表中删除用户失败")
		return
	}

	common.Success(c, nil, "用户已成功退出群聊")
}

// 添加成员到uc_match表中
func (b ChatController) AddMember(c *gin.Context) {
	var request struct {
		RoomID int `json:"roomID"`
		UserID int `json:"userID"`
	}
	c.Bind(&request)

	roomID := request.RoomID
	userID := request.UserID

	// 检查用户是否已经在 uc_match 表中
	var count int64
	b.DB.Table("uc_match").Where("cid = ? AND uid = ?", roomID, userID).Count(&count)
	if count > 0 {
		common.Fail(c, 400, nil, "用户已在聊天室中")
		return
	}

	// 添加成员到 uc_match 表
	newMatch := model.UcMatch{
		CID: uint(roomID),
		UID: uint(userID),
	}
	result := b.DB.Table("uc_match").Create(&newMatch)
	if result.Error != nil {
		common.Fail(c, 500, nil, "添加成员失败")
		return
	}

	common.Success(c, nil, "成员添加成功")
}

// 返回发言成员列表
func (b ChatController) GetSpeakers(c *gin.Context) {
	var request struct {
		RoomID int `json:"roomID"`
	}
	c.Bind(&request)

	roomID := request.RoomID

	// 查询 uc_match 表，获取发言成员列表
	var speakers []model.UcMatch
	result := b.DB.Table("uc_match").Where("cid = ?", roomID).Find(&speakers)
	if result.Error != nil {
		common.Fail(c, 500, nil, "获取发言成员列表失败")
		return
	}

	// 构建发言成员信息
	var speakerList []map[string]interface{}
	for _, speaker := range speakers {
		// 查询 user 表，获取成员信息
		var user model.User
		result = b.DB.Table("user").Where("uid = ?", speaker.UID).First(&user)
		if result.Error != nil {
			common.Fail(c, 500, nil, "查询成员信息失败")
			return
		}

		// 构建成员信息
		speakerInfo := map[string]interface{}{
			"userID":    user.UID,
			"userName":  user.UName,
			"userImage": user.UImage,
		}
		speakerList = append(speakerList, speakerInfo)
	}

	common.Success(c, gin.H{"speakerList": speakerList}, "成功返回发言成员列表")
}
