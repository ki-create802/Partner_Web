// src/api/api.js
import axios from 'axios';

const API_BASE_URL = 'http://localhost:8082';
//const test_URL = 'http://localhost:3000/api';

// 获取聊天列表
/*alert
*/
export const getChatsList = async (userID, searchWord) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/chatLists`, {
      userID,
      searchWord,
    });
    return response.data.data.chatList;
  } catch (error) {
    console.error('获取聊天列表失败:', error);
    throw error;
  }
};

// 获取聊天记录
export const getChatMessages = async (roomId, lastMessageId) => {
  //alert("获取聊天记录lastID："+JSON.stringify(lastMessageId));
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/chatRecords`, {
      cid: roomId,
      last_rid: lastMessageId,
    });
    //alert("获取聊天记录lastID："+JSON.stringify(lastMessageId)+"获取聊天记录："+JSON.stringify(response));
    if(response.data.data==null)return [];
    return response.data.data.chatRecords;
  } catch (error) {
    console.error('Error fetching chat messages:', error);
    throw error;
  }
};

// 发送消息
export const sendMessage = async (roomId, newMessage) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/saveChatRecords`, {
      roomid: roomId,
      newmessage: newMessage,
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error sending message:', error);
    throw error;
    
  }
};

// 获取房间成员
export const getRoomMember = async (roomId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/getRoomMember`, {
      roomID: roomId,
    });
    return response.data.data.RoomMemberList;
  } catch (error) {
    console.error('Error fetching room members:', error);
    throw error;
  }
};

// 获取在会话，但不在房间中的成员
export const getChatMember = async (roomId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/getSpeakers`, {
      roomID: roomId,
    });
    return response.data.data.speakerList;
  } catch (error) {
    console.error('Error fetching room members:', error);
    throw error;
  }
};

// 添加成员
export const addMember = async (roomId, chatMemberId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/successMatch`, {
      roomID: roomId,
      userID: chatMemberId,
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error adding member:', error);
    throw error;
  }
};

// 移除成员
/* 将房间成员移除
resquest post
  roomid(number)
  RoomMemberId(number)
response
  
*/
export const removeMember = async (roomId, RoomMemberId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/leaveChatRoom`, {
      roomID: roomId,
      userID:RoomMemberId,
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error removing member:', error);
    throw error;
  }
};

//解散·房间·
export const dissolveRoom = async (roomId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/disbandChatRoom`, {
      roomID: roomId,
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error removing member:', error);
    throw error;
  }
};

/* 获取热门数据
request get
response (返回的是数组，数组每一项对象如下结构)
    roomID (Number), 
    roomName (string)
    roomIntro (string)
    roomOwnerID (Number)
    roomOwnerName (string)
    memberList (List)
        memberId,memberId...
*/
export const fetchHotData = async () => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/getTopChatRooms`);
    return response.data.data.topChatRooms;
  } catch (error) {
    console.error('Error fetching hot data:', error);
    throw error;
  }
};

/*搜索功能
request post 
    query(string)
    range(string)
response (和热门数据返回内容的格式相同是个List)
    roomID (Number), 
    roomName (string)
    roomIntro (string)
    roomOwnerID (Number)
    roomOwnerName (string)
    memberList (List)
        memberId,memberId...
*/
export const search = async (query, scope) => {
  alert("搜索词："+JSON.stringify(query));
    if(!query)query="";
    try {
      const response = await axios.post(`${API_BASE_URL}/chat/searchChatList`, {
        query, 
        range: scope
      });
      if(response.data.data==null)return [];
      else return response.data.data.searchResults;
    } catch (error) {
      alert('搜索错误:'+error);
      throw error;
    }
  };
  


/*用户搜索
request post
  searchWord(string)
respone
  UID (number)
  UName(string) 
  URemark(string)
  UImage (string) //好像已经实现了uid同统一头像的功能，所以这一项没有我感觉问题不大
*/
export const searchUsers = async (searchWord) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/searchUser`, {
      query: searchWord
    });
    return response.data.data.userList;
  } catch (error) {
    console.error('Error searching users:', error);
    throw error;
  }
};


/*登录
request post: 
    Email(string)
    Password(string)
respone:
    UID(Number),
    UName(string),
    Ukey(string),
    UEmail(string),
    URemark(string),
    UImage(string) // 返回实际的头像URL
*/
// export const login = async (email, password) => {
//   try {
//     alert("进入登录,输入为"+JSON.stringify(email)+" "+JSON.stringify(password));
//     const response = await axios.post(`${API_BASE_URL}/user/login`, {
//       Email: email,
//       Password: password,
//     });
//     return response.data.data.userInfo;
//     // return response.data;
//   } catch (error) {
//     alert("ERROR:"+error);
//     throw error;
//   }
// };
export const login = async (email, password) => {
    try {
      const response = await axios.post(`${API_BASE_URL}/user/login`, {
        Email: email,
        Password: password,
      });
      return response.data.data.userInfo;
    } catch (error) {
      console.error('Error logging in:', error);
      throw error;
    }
  };
  

//粉丝数请求
export const getFansNum = async () => {
  try {
    //本地获取userId
    const userID=JSON.parse(localStorage.getItem("user")).UID;
    const response = await axios.post(`${API_BASE_URL}/user/fansNum`, {
      userID
    });
    return response.data.data.fansNum;
  } catch (error) {
    console.error('Error logging in:', error);
    throw error;
  }
};

// 注册
export const signUp = async (username,email, password) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/user/register`, {
      UName:username,
      Email: email,
      Password: password,
    });
    if(response.data && response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error signing up:', error);
    throw error;
  }
};


/*获取个人日程信息
request get
  Uid(Number)
respone (响应数据为List,每个List内容如下,List内按照时间先后排序)
  date(具体格式看后端)
  content(string)
 */
export const api_ScheduleItems = async (uid) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/user/schedule`, {
      userID: uid ,
    });
    //alert("后端获取行程信息："+JSON.stringify(response.data.data.ScheduleList));
    return response.data.data.ScheduleList;
  } catch (error) {
    // console.error('Error fetching schedule items:', error);
    alert('Error fetching schedule items:'+error);
    throw error;
  }
};

// 更改房间状态
export const changeRoomStatus = async (roomId, status) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/setRoomStatus`, {
      params: { roomid: roomId, setStatus: status },
    });
    return response.data;
  } catch (error) {
    console.error('Error changing room status:', error);
    throw error;
  }
};

/**
 * 创建新房间
 *  {string} roomName - 房间名称
 *  {number} typeID - 房间类型ID
 *  {number} Uid - 用户ID
 *  {number} memberNum - 房间人数
 *  {string} dateTime - 预约时间
 *  {string} remark - 备注
 * return {boolean} - 是否创建成功
 */
// 新建房间
export const newRoom = async (roomName, typeID,Uid,memberNum,dateTime,remark) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/createChat`, {
      Name:roomName,
      Bid:Number(typeID),
      Uid:Uid,
      Cnumber:Number(memberNum),
      CYueDate:dateTime,
      Cremark:remark
    });
    if (response.data.code==200)return true;
    else return false;
  } catch (error) {
    alert("新建房间失败"+error);
    console.error('Error signing up:', error);
    throw error;
  }
};

//api重置密码
export const checkCode = async (email,code) => {
  alert("邮箱："+email+"验证码"+code);
  try {
    const response = await axios.post(`${API_BASE_URL}/user/verify-reset-code`, {
      email,
      code
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error signing up:', error);
    throw error;
  }
};

export const forgetPW = async (email) => {
  alert("邮箱："+email);
  try {
    const response = await axios.post(`${API_BASE_URL}/user/forgot-password`, {
      email,
    });
    if(response.data.code==200)return true;
    else if(response.data.code==202){
      alert("该邮箱未注册");
      return false;
    }
    else return false;
  } catch (error) {
    console.error('Error signing up:', error);
    throw error;
  }
};

//重新设置密码
export const resetPW = async (email,newPW) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/user/resetPassword`, {
      Email:email,
      NewPassword:newPW
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error signing up:', error);
    throw error;
  }
};


//加入会话请求
export const joinChat = async (cid,uid) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/addMember`, {
      roomID:cid,
      userID:uid
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('Error signing up:', error);
    throw error;
  }
};
  
//修改密码
export const editPW = async (uid,newPW) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/user/editPassword`, {
      userid:uid,
      newPassword:newPW
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('修改密码失败', error);
    throw error;
  }
};

//修改昵称请求
export const editNewName = async (uid,newName) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/user/editInfo`, {
      UID:uid,
      UName:newName
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('修改密码失败', error);
    throw error;
  }
};

// 上传头像请求
export const uploadAvatar = async (uid, file) => {
  try {
    // 创建 FormData 对象
    const formData = new FormData();
    formData.append("userID", uid); // 添加用户ID（注意：后端使用的是 userID）
    formData.append("avatar", file); // 添加头像文件

    // 发送 POST 请求
    const response = await axios.post(`${API_BASE_URL}/file/uploadAvatar`, formData, {
      headers: {
        "Content-Type": "multipart/form-data", // 设置请求头
      },
    });

    // 根据响应结果返回 true 或 false
    if (response.data.code == 200) return true;
    else return false;
  } catch (error) {
    console.error("上传头像失败", error);
    throw error;
  }
};

//退出房间
export const getoutOfRoom = async (uid,cid) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/??`, {
      UID:uid,
      Cid:cid
    });
    if(response.data.code==200)return true;
    else return false;
  } catch (error) {
    console.error('退出房间失败：', error);
    throw error;
  }
};
