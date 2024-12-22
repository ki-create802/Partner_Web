const express = require('express');
const cors = require('cors');
const app = express();
const port = 3000; // 你可以选择其他端口

// 使用CORS中间件允许跨域请求
app.use(cors());

// 解析JSON请求体
app.use(express.json());

// 模拟的用户数据库
const users = [
    { email: 'user@example.com', password: 'password', imageId: 1 }
];


// 登录接口
app.post('/api/login', (req, res) => {
    const { Email, Password } = req.body;

    // 简单的验证逻辑
    const user = users.find(u => u.email === Email && u.password === Password);

    if (user) {
        res.status(200).json({  
            UID: 1,
            UName: "testUser",
            Ukey: "testUkey",
            UEmail: "user@example.com",
            URemark: "testURemark",
            UImage: "images/avatar1.png" // 返回实际的头像URL
        });
    } else {
        res.status(401).json({ message: 'Invalid email or password.' });
    }
});

//搜索
const searchResults = [
    { roomID: 1, roomName: '拼车', roomIntro: '寻找拼车伙伴', roomOwnerID: 1, roomOwnerName: '房主1', memberList: [], images: ['/images/roomuser1.jpg', '/images/roomuser2.jpg'] },
    { roomID: 2, roomName: '运动', roomIntro: '一起打篮球', roomOwnerID: 2, roomOwnerName: '房主2', memberList: [], images: ['/images/roomuser1.jpg'] },
    { roomID: 3, roomName: '娱乐', roomIntro: '看电影', roomOwnerID: 3, roomOwnerName: '房主3', memberList: [], images: ['/images/roomuser3.jpg', '/images/roomuser1.jpg'] },
    { roomID: 4, roomName: '拼单', roomIntro: '拼单购物', roomOwnerID: 4, roomOwnerName: '房主4', memberList: [], images: ['/images/roomuser3.jpg'] },
    { roomID: 5, roomName: '约拍', roomIntro: '寻找摄影师', roomOwnerID: 5, roomOwnerName: '房主5', memberList: [], images: ['/images/roomuser2.jpg'] },
    { roomID: 6, roomName: '其他', roomIntro: '其他活动', roomOwnerID: 6, roomOwnerName: '房主6', memberList: [], images: ['/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg'] }
];
// 搜索接口
app.get('/api/search', (req, res) => {
    const { Searchword,Range } = req.query;
    console.log("接受搜索信息为：Searchword:",Searchword);
    console.log("接受搜索信息为：Range:",Range);
    // 简单的搜索逻辑
    res.status(200).json(searchResults);
});


//获取个人日程
const scheduleItems =[
    { date: '10.01', content: '就是有事' },
    { date: '10.05', content: '行程1' },
    { date: '10.10', content: '行程2' },
    { date: '10.01', content: '行程3' },
    { date: '10.05', content: '行程4' },
    { date: '10.10', content: '行程5' }
  ];
app.get('/api/schedule', (req, res) => {
    const { Uid } = req.query;
    console.log("获取个人行程,Uid：",Uid);
    // 简单的搜索逻辑
    res.status(200).json(scheduleItems);
});

//
// 模拟的热门数据
const hotData = [
    { roomID: 1, roomName: '拼车', roomIntro: '寻找拼车伙伴', roomOwnerID: 1, roomOwnerName: '房主1', memberList: [], images: ['/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg','/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg'] },
    { roomID: 2, roomName: '运动', roomIntro: '一起打篮球', roomOwnerID: 2, roomOwnerName: '房主2', memberList: [],images: ['/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg']},
    { roomID: 3, roomName: '娱乐', roomIntro: '看电影', roomOwnerID: 3, roomOwnerName: '房主3', memberList: [], images: ['/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg']},
    { roomID: 4, roomName: '拼单', roomIntro: '拼单购物', roomOwnerID: 4, roomOwnerName: '房主4', memberList: [], images: ['/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg']},
    { roomID: 5, roomName: '约拍', roomIntro: '寻找摄影师', roomOwnerID: 5, roomOwnerName: '房主5', memberList: [],images: ['/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg']},
    { roomID: 6, roomName: '其他', roomIntro: '其他活动', roomOwnerID: 6, roomOwnerName: '房主6', memberList: [],images: ['/images/roomuser2.jpg', '/images/roomuser2.jpg','/images/roomuser3.jpg']}
  ];
  
  // 热门数据接口
  app.get('/api/hot', (req, res) => {
    res.status(200).json(hotData);
  });


// 模拟的用户数据库
const usersInfo = [
    { UID: 1, UName: "testUser1", URemark: "testURemark1", UImage: '/images/avatar1.png' },
    { UID: 2, UName: "testUser2", URemark: "testURemark2", UImage: '/images/avatar2.png' },
  ];
  
  // 用户列表接口
  app.get('/api/users', (req, res) => {
    console.log("请求用户列表");
    res.status(200).json(usersInfo);
  });
 


// 模拟的聊天室列表数据
let chatRooms = [
    { id: 1, name: '聊天室1',ownerImage:'images/avatar1.png'},
    { id: 2, name: '聊天室2',ownerImage:'images/avatar2.png'},
    { id: 3, name: '聊天室3',ownerImage:'images/roomuser1.jpg'},
    { id: 4, name: '聊天室4',ownerImage:'images/roomuser2.jpg'},
    { id: 5, name: '聊天室5',ownerImage:'images/roomuser3.jpg'},
    { id: 6, name: '聊天室6',ownerImage:'images/avatar1.png' },
    { id: 7, name: '聊天室7',ownerImage:'images/avatar1.png' },
    { id: 8, name: '聊天室8',ownerImage:'images/avatar1.png' },
    { id: 9, name: '聊天室9' ,ownerImage:'images/avatar1.png'},
    { id: 10, name: '聊天室10',ownerImage:'images/avatar1.png' },
    { id: 11, name: '聊天室11',ownerImage:'images/avatar1.png' },
    { id: 12, name: '聊天室12',ownerImage:'images/avatar1.png' },
    { id: 13, name: '聊天室13',ownerImage:'images/avatar1.png' },
    { id: 14, name: '聊天室14',ownerImage:'images/avatar1.png' },
    { id: 15, name: '聊天室15' ,ownerImage:'images/avatar1.png'},
    { id: 16, name: '聊天室16' ,ownerImage:'images/avatar1.png'},
    { id: 17, name: '聊天室17' ,ownerImage:'images/avatar1.png'},
    { id: 18, name: '聊天室18' ,ownerImage:'images/avatar1.png'},
    { id: 19, name: '聊天室19' ,ownerImage:'images/avatar1.png'},
];

// 聊天室列表接口
app.get('/api/chats', (req, res) => {
    const {uid,searchWord} = req.query; // 获取 uid 参数
    console.log('请求聊天列表，uid：', uid); // 打印 uid 参数
    console.log('请求SearchWord：',searchWord);
    res.status(200).json(chatRooms);
});


//获取头像·
app.use(express.static('public'));

app.get('/api/user/:uid/avatar', (req, res) => {
    const uid = req.params.uid;
    const user = usersInfo.find(u => u.UID == uid);

    if (user) {
        res.status(200).json({ avatarUrl: user.UImage });
    } else {
        res.status(404).json({ message: 'User not found.' });
    }
});


//模拟聊天记录
const chatMessages = {
    1: [
      { id: 1, text: 'Hello!', senderId: 23 ,senderName:'机器人1号' ,senderAvatarSrc:'images/avatar1.png',isImage:false, imageSrc:''},
      { id: 2, text: 'Hi there!', senderId: 432 ,senderName:'机器人2号',senderAvatarSrc:'images/avatar2.png',isImage:false, imageSrc:''},
    ],
    2: [
      { id: 1, text: 'Welcome!', senderId: 2342,senderName:'机器人1号' ,senderAvatarSrc:'images/avatar1.png',isImage:false, imageSrc:''},
      { id: 2, text: 'Thanks!', senderId: 232 ,senderName:'机器人2号',senderAvatarSrc:'images/avatar2.png',isImage:false, imageSrc:''},
    ],
    // 其他聊天室的历史消息
  };
  
// 获取聊天室消息接口
// 长轮询接口，用于实时获取新消息
app.get('/api/pollchats/:chatId/messages', (req, res) => {
  const chatId = req.params.chatId;
  const lastMessageId = req.query.lastMessageId || 0; // 获取客户端传来的最后一条消息的ID

  // 模拟长轮询，等待新消息
  const checkForNewMessages = () => {
      const messages = chatMessages[chatId] || [];
      const newMessages = messages.filter(msg => msg.id > lastMessageId);

      if (newMessages.length > 0) {
          res.status(200).json(newMessages);
          console.log("发送信息：",newMessages);
      } else {
          // 如果没有新消息，等待一段时间后再次检查
          setTimeout(checkForNewMessages, 2000); // 每2秒检查一次
      }
  };

  checkForNewMessages();
});

//发送聊天信息
app.get('/api/sendchats/', (req, res) => {
    console.log('发送聊天消息');
    const { roomid, newmessage } = req.query;
    const newMessageData = JSON.parse(newmessage); // 将 JSON 字符串转换为对象
  
    console.log('接受信息:', newMessageData);
  
    // 将新消息添加到对应的聊天室消息列表中
    if (!chatMessages[roomid]) {
      chatMessages[roomid] = [];
    }
    chatMessages[roomid].push(newMessageData);
  
    res.status(200).json({ message: 'Message sent successfully' });
});

//设置房间状态
app.get('/api/setRoomStatus/', (req, res) => {
  const {roomid,setStatus } = req.query;
  console.log('设置房间：',roomid,'的聊天状态为：',setStatus);
  if(setStatus=="解散"){
    chatRooms=chatRooms.filter(room => room.id !== parseInt(roomid, 10))
    console.log('成功解散房间，解散后房间信息为：',chatRooms);
  }
  res.status(200).json({ success: true });
});

//模拟chatMember
const chatMember=[
{ id:23 , name:'member1' , avatar:'images/avatar1.png'},
{ id:33 , name:'member2' , avatar:'images/avatar2.png'},
{ id:43 , name:'member3' , avatar:'images/roomuser3.jpg'},
];

//获取在会话中但不在项目中的成员列表
app.get('/api/getChatMember/', (req, res) => {
  const {roomid} = req.query;
  console.log('获取房间：',roomid,'的会话成员列表为：',chatMember);
  res.status(200).json(chatMember);
});

//添加会话成员到项目中
app.post('/api/addMember', (req, res) => {
  const { roomid,chatMemberId } = req.body;
  console.log("房间",roomid,"  将成员",chatMemberId,"加入到会话中");
  res.status(200).json({message:"成功将会话成员添加到项目中"});
});

//获取房间成员
app.get('/api/getRoomMember/', (req, res) => {
  const {roomid} = req.query;
  console.log('获取房间：',roomid,'的房间成员列表为：',chatMember);
  res.status(200).json(chatMember);
});

//移除项目成员
app.post('/api/removeMember', (req, res) => {
  const { roomid,roomMemberId } = req.body;
  console.log("房间",roomid,"  将成员",roomMemberId,"从房间移除");
  res.status(200).json({message:"成功将会话成员从房间移除"});
});

// 启动服务器
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});