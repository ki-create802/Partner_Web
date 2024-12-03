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

// 头像URL映射
const imageUrls = {
    1: 'https://example.com/images/avatar1.png',
    2: 'https://example.com/images/avatar2.png'
};

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
            UImage: imageUrls[user.imageId] // 返回实际的头像URL
        });
    } else {
        res.status(401).json({ message: 'Invalid email or password.' });
    }
});

//搜索
const searchResults = [
    { roomID: 1, roomName: '拼车', roomIntro: '寻找拼车伙伴' ,roomOwnerID:1,roomOwnerName:'房主1',memberList:[]},
    { roomID: 2, roomName: '运动', roomIntro: '一起打篮球' ,roomOwnerID:2,roomOwnerName:'房主2',memberList:[]},
    { roomID: 3, roomName: '娱乐', roomIntro: '看电影' ,roomOwnerID:3,roomOwnerName:'房主3',memberList:[]},
    { roomID: 4, roomName: '拼单', roomIntro: '拼单购物' ,roomOwnerID:4,roomOwnerName:'房主4',memberList:[]},
    { roomID: 5, roomName: '约拍', roomIntro: '寻找摄影师' ,roomOwnerID:5,roomOwnerName:'房主5',memberList:[]},
    { roomID: 6, roomName: '其他', roomIntro: '其他活动' ,roomOwnerID:6,roomOwnerName:'房主6',memberList:[]}
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
    console.log("接受搜索信息为,Uid：",Uid);
    // 简单的搜索逻辑
    res.status(200).json(scheduleItems);
});

//
// 模拟的热门数据
const hotData = [
    { roomID: 1, roomName: '拼车', roomIntro: '寻找拼车伙伴', roomOwnerID: 1, roomOwnerName: '房主1', memberList: [], images: [] },
    { roomID: 2, roomName: '运动', roomIntro: '一起打篮球', roomOwnerID: 2, roomOwnerName: '房主2', memberList: [], images: [] },
    { roomID: 3, roomName: '娱乐', roomIntro: '看电影', roomOwnerID: 3, roomOwnerName: '房主3', memberList: [], images: [] },
    { roomID: 4, roomName: '拼单', roomIntro: '拼单购物', roomOwnerID: 4, roomOwnerName: '房主4', memberList: [], images: [] },
    { roomID: 5, roomName: '约拍', roomIntro: '寻找摄影师', roomOwnerID: 5, roomOwnerName: '房主5', memberList: [], images: [] },
    { roomID: 6, roomName: '其他', roomIntro: '其他活动', roomOwnerID: 6, roomOwnerName: '房主6', memberList: [], images: [] }
  ];
  
  // 热门数据接口
  app.get('/api/hot', (req, res) => {
    res.status(200).json(hotData);
  });


// 模拟的用户数据库
const usersInfo = [
    { UID: 1, UName: "testUser1", Ukey: "testUkey1", UEmail: "user1@example.com", URemark: "testURemark1", UImage: 'https://example.com/images/avatar1.png' },
    { UID: 2, UName: "testUser2", Ukey: "testUkey2", UEmail: "user2@example.com", URemark: "testURemark2", UImage: 'https://example.com/images/avatar2.png' },
  ];
  
  // 用户列表接口
  app.get('/api/users', (req, res) => {
    console.log("请求用户列表");
    res.status(200).json(usersInfo);
  });

// 启动服务器
app.listen(port, () => {
    console.log(`Server is running on http://localhost:${port}`);
});