package controller

import (
	"Partner_Web/Partner_Server/common"
	"Partner_Web/Partner_Server/model"
	"fmt"
	"sort"
	"strings"

	"time"

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
	CreateChat(c *gin.Context)      //创建聊天室
	ChatLists(c *gin.Context)       //返回聊天室聊天列表
	ChatRecords(c *gin.Context)     //返回聊天室对应聊天记录
	SaveChatRecords(c *gin.Context) //存储聊天记录

	DisbandChatRoom(c *gin.Context) //房主解散群聊
	LeaveChatRoom(c *gin.Context)   //退出群聊
	AddMember(c *gin.Context)       //添加成员到uc_match表中

	GetSpeakers(c *gin.Context)   //返回发言成员列表
	GetRoomMember(c *gin.Context) //返回成功配对成员列表

	SuccessMatch(c *gin.Context) //成功加入聊天室 添加到success_match表格中

	GetPendingChats(c *gin.Context) //获取主页'等待中'板块聊天室列表
	GetHistoryChats(c *gin.Context) //获取主页'历史'板块聊天室列表

	GetTopChatRooms(c *gin.Context) //火文推送
}

func NewChatController() InChatController {
	db := common.GetDB()
	//db.Table("user").AutoMigrate(model.User{})
	return ChatController{DB: db}
}

func (b ChatController) SearchChatList(c *gin.Context) { //未测试
	var queryString model.SearchString
	c.Bind(&queryString)

	fmt.Println("query:", queryString.SearchInfo)
	fmt.Println("range:", queryString.Block)

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
	var chatList []model.ChatSearch   //聊天室数据
	if queryString.SearchInfo != "" { //字符串不为空 匹配聊天室名称
		if IsAll { //筛选全部
			result := b.DB.Table("chatinfo").
				Where("cstate =?", 1).
				Where("cname LIKE ? OR cremark LIKE ?", "%"+queryString.SearchInfo+"%", "%"+queryString.SearchInfo+"%").
				Find(&chatList)
			if result.Error != nil {
				common.Fail(c, 500, nil, "搜索词数据库查询错误")
				return
			}
		} else { //多一个版块筛选
			result := b.DB.Table("chatinfo").
				Where("bid = ?", blockID).
				Where("cstate =?", 1).
				Where("cname LIKE ? OR cremark LIKE ?", "%"+queryString.SearchInfo+"%", "%"+queryString.SearchInfo+"%").Find(&chatList)
			if result.Error != nil {
				common.Fail(c, 500, nil, "版块搜索词数据库查询错误")
				return
			}
		}
	} else { //为空就是全部推送，再判断block
		if IsAll { //全部推送
			result := b.DB.Table("chatinfo").
				Where("cstate =?", 1).
				Find(&chatList)
			if result.Error != nil {
				common.Fail(c, 500, nil, "全部聊天室数据库查询错误")
				return
			}
		} else { //推送对应板块所有记录
			result := b.DB.Table("chatinfo").
				Where("cstate =?", 1).
				Where("bid = ?", blockID).
				Find(&chatList)
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
	fmt.Println("searchResults", searchResults)
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
		// common.Fail(c, 400, nil, "请输入相关用户关键词")
		// return
		result := b.DB.Table("user").Select("uid,uname,uremark,uimage").
			Find(&userList)
		if result.Error != nil {
			common.Fail(c, 500, nil, "用户查询数据库错误")
			return
		}
	}

	common.Success(c, gin.H{"userList": userList}, "成功返回用户列表")
}

// 房主解散群聊
func (b ChatController) DisbandChatRoom(c *gin.Context) {
	var request struct {
		RoomID int `json:"roomID"`
	}
	if err := c.Bind(&request); err != nil {
		common.Fail(c, 400, nil, "请求参数解析失败")
		return
	}

	roomID := request.RoomID

	// 更新房间状态为 0，表示房间被解散
	result := b.DB.Table("chatinfo").Where("cid = ?", roomID).Update("cstate", 0)
	if result.Error != nil {
		common.Fail(c, 500, nil, "更新房间状态失败")
		return
	}

	// 获取聊天室的所有成员
	var members []model.SuccessMatch
	if err := b.DB.Table("success_match").Where("cid = ?", roomID).Find(&members).Error; err != nil {
		common.Fail(c, 500, nil, "获取聊天室成员失败")
		return
	}

	// 更新每个成员的等级
	for _, member := range members {
		if err := b.UpdateUserLevel(member.UID); err != nil {
			common.Fail(c, 500, nil, "更新用户等级失败")
			return
		}
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

	// // 从 uc_match 表中删除用户
	// result = b.DB.Table("uc_match").Where("cid = ? AND uid = ?", roomID, userID).Delete(&model.UcMatch{})
	// if result.Error != nil {
	// 	common.Fail(c, 500, nil, "从 uc_match 表中删除用户失败")
	// 	return
	// }

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

// 返回发言成员列表
func (b ChatController) GetRoomMember(c *gin.Context) {
	var request struct {
		RoomID int `json:"roomID"`
	}
	c.Bind(&request)

	roomID := request.RoomID

	// 查询 success_match 表，获取发言成员列表
	var speakers []model.UcMatch
	result := b.DB.Table("success_match").Where("cid = ?", roomID).Find(&speakers)
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
	fmt.Println(speakerList)
	common.Success(c, gin.H{"RoomMemberList": speakerList}, "成功返回发言成员列表")
}

func (b ChatController) CreateChat(c *gin.Context) { //创建聊天室 版块传消息的类型
	var requestChat model.Room
	// c.Bind(&requestChat)

	// fmt.Println("bid", requestChat.Bid)
	if err := c.Bind(&requestChat); err != nil {
		common.Fail(c, 400, nil, "请求参数错误: "+err.Error())
		fmt.Println("请求参数错误: ", err.Error())
		return
	}

	fmt.Println("传入的消息111111", requestChat)

	// 检查必要字段
	if requestChat.Bid == 0 || requestChat.Name == "" || requestChat.Uid == 0 {
		common.Fail(c, 400, nil, "版块ID、聊天室名称或用户ID不能为空")
		fmt.Println("版块ID、聊天室名称或用户ID不能为空")
		return
	}

	currentTime := time.Now() //获取当前时间

	fmt.Println("111111", requestChat.CYueDate)
	fmt.Println("日期111111", requestChat.CYueDate)

	newChat := model.Room{
		Name:        requestChat.Name,
		Bid:         requestChat.Bid,
		Uid:         requestChat.Uid,
		Cnumber:     requestChat.Cnumber,
		Cstate:      1,
		CYueDate:    requestChat.CYueDate,
		CCreateTime: currentTime,
		Cremark:     requestChat.Cremark,
	}

	result := b.DB.Table("chatinfo").Create(&newChat)
	if result.Error != nil {
		common.Fail(c, 500, nil, "新增聊天室信息数据库错误")
		fmt.Println("新增聊天室信息数据库错误")
		return
	}

	common.Success(c, gin.H{"newChat": newChat}, "成功新建聊天室") //返回的cid没有值，但是数据库里是自增长的
}

func (b ChatController) SuccessMatch(c *gin.Context) { //添加成功匹配表
	var request struct {
		RoomID int `json:"roomID"`
		UserID int `json:"userID"`
	}
	c.Bind(&request)

	roomID := request.RoomID
	userID := request.UserID

	var count int64
	b.DB.Table("success_match").Where("cid = ? AND uid = ?", roomID, userID).Count(&count)
	if count > 0 {
		common.Fail(c, 400, nil, "用户已与该聊天室成功配对")
		return
	}

	newMatch := model.SuccessMatch{
		CID: uint(roomID),
		UID: uint(userID),
	}
	result := b.DB.Table("success_match").Create(&newMatch)
	if result.Error != nil {
		common.Fail(c, 500, nil, "添加成员失败")
		return
	}

	// 从 uc_match 表中删除相应的记录
	if err := b.DB.Table("uc_match").Where("cid = ? AND uid = ?", roomID, userID).Delete(&model.UcMatch{}).Error; err != nil {
		common.Fail(c, 500, nil, "删除 uc_match 记录失败")
		return
	}

	common.Success(c, nil, "成员添加成功")
}

func (b ChatController) ChatLists(c *gin.Context) {
	var request struct {
		UID        uint   `json:"userID"`
		SearchWord string `json:"searchWord"`
	}
	if err := c.Bind(&request); err != nil {
		common.Fail(c, 400, nil, "请求参数解析失败")
		return
	}
	// fmt.Println("UID:", request.UID)

	var cids []uint
	if err := b.DB.Table("uc_match").
		Where("uid = ?", request.UID).
		Pluck("cid", &cids).Error; err != nil {
		common.Fail(c, 500, nil, "查询失败")
		return
	}

	// 查询 chatinfo 表中 uid 等于传入的 userID 的 cid
	var additionalCids []uint
	if err := b.DB.Table("chatinfo").
		Where("uid = ?", request.UID).
		Pluck("cid", &additionalCids).Error; err != nil {
		common.Fail(c, 500, nil, "查询 chatinfo 表失败")
		return
	}

	// 从 success_match 表中获取 cid
	var successMatchCids []uint
	if err := b.DB.Table("success_match").
		Where("uid = ?", request.UID).
		Pluck("cid", &successMatchCids).Error; err != nil {
		common.Fail(c, 500, nil, "查询 success_match 表失败")
		return
	}

	// 将两个查询结果合并到 cids 列表中
	cids = append(cids, additionalCids...)
	cids = append(cids, successMatchCids...)

	fmt.Println("cids", cids)

	// 去重 cids 列表中的重复项
	uniqueCids := make(map[uint]bool)
	for _, cid := range cids {
		uniqueCids[cid] = true
	}

	// 将去重后的 cids 重新赋值
	cids = make([]uint, 0, len(uniqueCids))
	for cid := range uniqueCids {
		cids = append(cids, cid)
	}

	if len(cids) == 0 {
		common.Success(c, nil, "未找到聊天室记录")
		return
	}

	type ChatInfo struct {
		CID     uint   `gorm:"column:cid"`
		CName   string `gorm:"column:cname"`
		UID     uint   `gorm:"column:uid"`
		CRemark string `gorm:"column:cremark"`
	}
	// 从 chat_info 表中找到 cid 在 cids 列表中的行
	var chatInfos []ChatInfo
	result := b.DB.Table("chatinfo").
		Where("cstate =?", 1).
		Where("cid IN (?)", cids).
		Find(&chatInfos)
	if result.Error != nil {
		common.Fail(c, 500, nil, "查询 chatinfo 表失败")
		return
	}

	fmt.Println("chatInfos11111", chatInfos)

	// 排序逻辑
	sort.Slice(chatInfos, func(i, j int) bool {
		// 检查 searchWord 是否在 cname 或 cremark 中
		containsI := strings.Contains(chatInfos[i].CName, request.SearchWord) || strings.Contains(chatInfos[i].CRemark, request.SearchWord)
		containsJ := strings.Contains(chatInfos[j].CName, request.SearchWord) || strings.Contains(chatInfos[j].CRemark, request.SearchWord)

		// 如果 i 包含 searchWord 而 j 不包含，i 排在前面
		if containsI && !containsJ {
			return true
		}
		// 如果 j 包含 searchWord 而 i 不包含，j 排在前面
		if containsJ && !containsI {
			return false
		}
		// 如果都包含或都不包含，保持原有顺序
		return i < j
	})

	fmt.Println("chatInfos22222", chatInfos)

	// 构造返回数据
	type Response struct {
		CID     uint   `json:"id"`
		UID     uint   `json:"ownerId"`
		CName   string `json:"name"`
		CRemark string `json:"remark"`
	}
	var responses []Response
	for _, chatInfo := range chatInfos {
		responses = append(responses, Response{
			CID:     chatInfo.CID,
			UID:     chatInfo.UID,
			CName:   chatInfo.CName,
			CRemark: chatInfo.CRemark,
		})
	}
	// 返回结果
	common.Success(c, gin.H{"chatList": responses}, "查询成功")

}

func (b ChatController) ChatRecords(c *gin.Context) {
	var request struct {
		CID     int `json:"cid"`
		LastRID int `json:"last_rid"`
	}
	if err := c.Bind(&request); err != nil {
		common.Fail(c, 400, nil, "请求参数解析失败")
		return
	}
	// fmt.Println("cid", request.CID)
	// fmt.Println("Rid", request.LastRID)

	// 查询 record 表中 cid 等于 request.CID 的最大 ridPerChat
	var maxRIDPerChat struct {
		MAXNum int `gorm:"column:max_rid"`
	}
	if err := b.DB.Debug().Table("record").
		Select("MAX(ridPerChat) AS max_rid").
		Where("cid = ?", request.CID).
		First(&maxRIDPerChat).Error; err != nil {
		common.Fail(c, 500, nil, "查询最大 ridPerChat 失败")
		return
	}

	// fmt.Println("maxRIDPerChat1:", maxRIDPerChat.MAXNum)

	// 如果查询到的 MAXNum 为 0，表示没有聊天记录
	if maxRIDPerChat.MAXNum == 0 {
		common.Success(c, nil, "暂无聊天记录")
		return
	}

	// fmt.Println("maxRIDPerChat2:", maxRIDPerChat.MAXNum)

	//如果 maxRIDPerChat > request.LastRID，查询新增的记录
	if maxRIDPerChat.MAXNum > request.LastRID {
		var newRecords []model.Record
		if err := b.DB.Debug().Table("record").
			Where("cid = ? AND ridPerChat > ?", request.CID, request.LastRID).
			Find(&newRecords).Error; err != nil {
			common.Fail(c, 500, nil, "查询新增记录失败")
			return
		}

		// 构造返回数据
		var responses []gin.H
		for _, record := range newRecords {
			var user struct {
				UName  string `gorm:"column:uname"`
				UImage string `gorm:"column:uimage"`
			}
			if err := b.DB.Debug().Table("user").
				Select("uname, uimage").
				Where("uid = ?", record.UID).
				First(&user).Error; err != nil {
				common.Fail(c, 500, nil, "查询用户信息失败")
				return
			}

			// 构造返回数据
			responses = append(responses, gin.H{
				"RIDPerChat": record.RIDPerChat,
				"CID":        record.CID,
				"UID":        record.UID,
				"RType":      record.RType,
				"RContent":   record.RContent,
				"RTime":      record.RTime,
				"UName":      user.UName,  // 添加用户名
				"UImage":     user.UImage, // 添加用户头像
			})
		}

		// 4. 返回结果
		common.Success(c, gin.H{"chatRecords": responses}, "查询成功")
	} else {
		// 如果没有新增记录，返回空数组
		common.Success(c, nil, "没有新增记录")
	}

}

func (b ChatController) SaveChatRecords(c *gin.Context) { //存储聊天记录

	type NewMessageData struct {
		ID       uint   `json:"id"`
		Text     string `json:"text"`
		SenderID uint   `json:"senderId"`
		IsImage  bool   `json:"isImage"`
	}

	var req struct {
		RoomID     uint           `json:"roomid"`
		NewMessage NewMessageData `json:"newmessage"`
	}
	if err := c.Bind(&req); err != nil {
		common.Fail(c, 400, nil, "请求参数解析失败")
		return
	}

	fmt.Println("前端：", req)

	// 2. 将 isImage 转换为 rtype
	var rtype int
	if req.NewMessage.IsImage {
		rtype = 0 // 图片
	} else {
		rtype = 1 // 文字
	}

	rtime := time.Now()

	// 3. 将数据插入到数据库的 record 表中
	record := model.Record{
		CID:        req.RoomID,
		RIDPerChat: req.NewMessage.ID,
		RContent:   req.NewMessage.Text,
		UID:        req.NewMessage.SenderID,
		RType:      rtype,
		RTime:      rtime,
	}

	result := b.DB.Table("record").Create(&record)
	fmt.Println("记录：", record)
	if result.Error != nil {
		common.Fail(c, 500, nil, "插入新聊天记录失败")
		fmt.Println("插入record表失败12345")
		return
	}

	common.Success(c, gin.H{"record": record}, "成功插入聊天室记录")
}

func (b ChatController) GetPendingChats(c *gin.Context) { //获取主页'等待中'板块聊天室列表
	var request struct {
		UserID uint `json:"userID"`
	}
	if err := c.Bind(&request); err != nil {
		common.Fail(c, 400, nil, "请求参数解析失败")
		return
	}

	userID := request.UserID

	// 获取未确认的聊天室（用户作为成员但未成功匹配）
	var ucChats []model.ChatInfo
	if err := b.DB.Table("chatinfo").
		Joins("JOIN uc_match ON chatinfo.cid = uc_match.cid").
		Where("uc_match.uid = ? AND chatinfo.cstate = 1", userID).
		Where("NOT EXISTS (SELECT 1 FROM success_match WHERE success_match.cid = chatinfo.cid AND success_match.uid = ?)", userID).
		Find(&ucChats).Error; err != nil {
		common.Fail(c, 500, nil, "查询未确认聊天室失败")
		return
	}

	// 获取已经成功匹配的聊天室（用户作为成员且已成功匹配）
	var successChats []model.ChatInfo
	if err := b.DB.Table("chatinfo").
		Joins("JOIN success_match ON chatinfo.cid = success_match.cid").
		Where("success_match.uid = ? AND chatinfo.cstate = 1", userID).
		Find(&successChats).Error; err != nil {
		common.Fail(c, 500, nil, "查询成功匹配聊天室失败")
		return
	}

	// 获取用户作为群主的聊天室
	var ownerChats []model.ChatInfo
	if err := b.DB.Table("chatinfo").
		Where("uid = ? AND cstate = 1", userID).
		Find(&ownerChats).Error; err != nil {
		common.Fail(c, 500, nil, "查询群主聊天室失败")
		return
	}

	// 合并三个列表
	var pendingChats []model.ChatInfo
	pendingChats = append(pendingChats, ucChats...)
	pendingChats = append(pendingChats, successChats...)
	pendingChats = append(pendingChats, ownerChats...)

	common.Success(c, gin.H{"pendingChats": pendingChats}, "成功返回等待中聊天室列表")
}

func (b ChatController) GetHistoryChats(c *gin.Context) { //获取历史聊天室列表
	var request struct {
		UserID uint `json:"userID"`
	}
	if err := c.Bind(&request); err != nil {
		common.Fail(c, 400, nil, "请求参数解析失败")
		return
	}

	userID := request.UserID

	// 查询用户作为成员的历史聊天室
	var memberChats []model.ChatInfo
	if err := b.DB.Table("chatinfo").
		Joins("JOIN success_match ON chatinfo.cid = success_match.cid").
		Where("success_match.uid = ? AND chatinfo.cstate = 0", userID).
		Find(&memberChats).Error; err != nil {
		common.Fail(c, 500, nil, "查询成员历史聊天室失败")
		return
	}

	// 查询用户作为群主的历史聊天室
	var ownerChats []model.ChatInfo
	if err := b.DB.Table("chatinfo").
		Where("uid = ? AND cstate = 0", userID).
		Find(&ownerChats).Error; err != nil {
		common.Fail(c, 500, nil, "查询群主历史聊天室失败")
		return
	}

	// 合并两个列表
	var historyChats []model.ChatInfo
	historyChats = append(historyChats, memberChats...)
	historyChats = append(historyChats, ownerChats...)

	common.Success(c, gin.H{"historyChats": historyChats}, "成功返回历史聊天室列表")
}

// 更新用户等级
func (b ChatController) UpdateUserLevel(uid uint) error {
	// 获取用户作为成员的历史聊天室列表
	var memberChats []model.ChatInfo
	if err := b.DB.Table("chatinfo").
		Joins("JOIN success_match ON chatinfo.cid = success_match.cid").
		Where("success_match.uid = ? AND chatinfo.cstate = 0", uid).
		Find(&memberChats).Error; err != nil {
		return err
	}

	// 获取用户作为群主的历史聊天室列表
	var ownerChats []model.ChatInfo
	if err := b.DB.Table("chatinfo").
		Where("uid = ? AND cstate = 0", uid).
		Find(&ownerChats).Error; err != nil {
		return err
	}

	// 合并两个列表并获取总数
	count := len(memberChats) + len(ownerChats)

	// 根据活动次数计算用户等级
	var level int
	switch {
	case count >= 30:
		level = 4 // 搭子达人
	case count >= 20:
		level = 3 // 高级搭子行家
	case count >= 10:
		level = 2 // 中级搭子爱好者
	case count >= 5:
		level = 1 // 初级搭子体验者
	default:
		level = 0 // 搭子探索者
	}

	// 更新或插入用户等级
	userLevel := model.UserLevel{
		UID:   uid,
		Level: level,
	}
	result := b.DB.Table("user_level").
		Where("uid = ?", uid).
		Assign(map[string]interface{}{"level": level}).
		FirstOrCreate(&userLevel)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (b ChatController) GetTopChatRooms(c *gin.Context) {
	// 获取当前时间
	currentTime := time.Now()
	// 计算一天前的时间
	oneDayAgo := currentTime.Add(-24 * time.Hour)

	// 查询最近一天内聊天记录数最多的十个聊天室
	var topChatRooms []struct {
		CID         uint      `gorm:"column:cid"`
		CName       string    `gorm:"column:cname"`
		Bid         int       `gorm:"column:bid"`
		Uid         uint      `gorm:"column:uid"`
		CNumber     int       `gorm:"column:cnumber"`
		CState      int       `gorm:"column:cstate"`
		CYueDate    string    `gorm:"column:c_yue_date"`
		CCreateTime time.Time `gorm:"column:c_create_time"`
		CRemark     string    `gorm:"column:cremark"`
		RecordCount int       `gorm:"column:record_count"`
	}

	result := b.DB.Table("record").
		Select("chatinfo.cid, chatinfo.cname, chatinfo.bid, chatinfo.uid, chatinfo.cnumber, chatinfo.cstate, chatinfo.c_yue_date, chatinfo.c_create_time, chatinfo.cremark, COUNT(record.cid) as record_count").
		Joins("JOIN chatinfo ON record.cid = chatinfo.cid").
		Where("record.rtime BETWEEN ? AND ?", oneDayAgo, currentTime).
		Group("record.cid").
		Order("record_count DESC").
		Limit(10).
		Find(&topChatRooms)

	if result.Error != nil {
		common.Fail(c, 500, nil, "查询聊天室记录失败")
		return
	}

	// 构造返回数据
	var response []map[string]interface{}
	for _, chatRoom := range topChatRooms {
		// 查询房主信息
		var roomOwner model.User
		result := b.DB.Table("user").
			Where("uid = ?", chatRoom.Uid).
			First(&roomOwner)
		if result.Error != nil {
			common.Fail(c, 500, nil, "查询房主信息失败")
			return
		}

		// 查询 success_match 表，获取成员列表
		var successMatches []struct {
			UID uint `gorm:"column:uid"`
		}
		result = b.DB.Table("success_match").
			Select("uid").
			Where("cid = ?", chatRoom.CID).
			Find(&successMatches)
		if result.Error != nil {
			common.Fail(c, 500, nil, "查询成员列表失败")
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
				common.Fail(c, 500, nil, "查询成员信息失败")
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
		roomInfo := map[string]interface{}{
			"roomID":        chatRoom.CID,
			"roomName":      chatRoom.CName,
			"roomIntro":     chatRoom.CRemark,
			"roomOwnerID":   chatRoom.Uid,
			"roomOwnerName": roomOwner.UName,
			"roomOwnerImg":  roomOwner.UImage,
			"recordCount":   chatRoom.RecordCount,
			"memberList":    memberList, // 添加成员列表
		}
		response = append(response, roomInfo)
	}

	common.Success(c, gin.H{"topChatRooms": response}, "成功返回最近一天内聊天记录数最多的十个聊天室")
}
