// src/api/api.js
import axios from 'axios';

const API_BASE_URL = 'http://localhost:8082';
//const test_URL = 'http://localhost:3000/api';

// 获取聊天列表
/*
*/
export const getChatsList = async (userID, searchWord) => {
  alert("uid,search="+JSON.stringify(userID)+JSON.stringify(searchWord));
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/chatLists`, {
      userID,
      searchWord,
    });
    // const response = await axios.get(`${test_URL}/api/chats`, {
    //   uid,
    //   searchWord,
    // });
    alert("response.data.data:"+JSON.stringify(response.data.data));
    return response.data.data.chatList;
  } catch (error) {
    alert('Error fetching chat list:'+error);
    throw error;
  }
};

// 获取聊天记录
export const getChatMessages = async (roomId,lastMessageId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/chat/chatRecords`, {
      cid : roomId,
      last_rid : lastMessageId,
    });
    //alert("response.data.data.chatRecords:"+JSON.stringify(response.data.data.chatRecords));
    return response.data.data.chatRecords;
  } catch (error) {
    //alert("error"+error);
    console.error('Error fetching chat messages:', error);
    throw error;
  }
};

// 发送消息
export const sendMessage = async (roomId, newMessage) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/sendchats/`, {
      params: { roomid: roomId, newmessage: newMessage },
    });
    return response.data;
  } catch (error) {
    console.error('Error sending message:', error);
    throw error;
  }
};

// 获取房间成员
export const getRoomMember = async (roomId) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/getRoomMember`, {
      params: { roomid: roomId },
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching room members:', error);
    throw error;
  }
};

// 添加成员
export const addMember = async (roomId, chatMemberId) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/addMember`, {
      roomid: roomId,
      chatMemberId,
    });
    return response.data;
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
    const response = await axios.post(`${API_BASE_URL}/removeMember`, {
      roomid: roomId,
      RoomMemberId,
    });
    return response.data;
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
    const response = await axios.get(`${API_BASE_URL}/hot`);
    return response.data;
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
    alert("后端返回用户俩表："+JSON.stringify(response));
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
      alert("进入登录");
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
    alert("后端粉丝数"+JSON.stringify(response.data.data));
    return response.data.data.fansNum;
    
  } catch (error) {
    console.error('Error logging in:', error);
    throw error;
  }
};

// 注册
export const signUp = async (email, password) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/signup`, {
      Email: email,
      Password: password,
    });
    return response.data;
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
    alert(JSON.stringify(response));
    alert("chat新建："+JSON.stringify({roomName, typeID,Uid,memberNum,dateTime,remark}));
    if (response.data.code==200)return true;
    else return false;
  } catch (error) {
    
    alert("新建房间失败");
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

//忘记密码
export const forgetPW = async (email) => {
  alert("邮箱："+email);
  try {
    const response = await axios.post(`${API_BASE_URL}/user/forgot-password`, {
      email,
    });
    if(response.data.code==200)return true;
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