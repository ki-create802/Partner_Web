// src/api/api.js
import axios from 'axios';

const API_BASE_URL = 'http://localhost:3000/api';

// 获取聊天列表
export const getChatsList = async (uid, searchWord) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/chats`, {
      params: { uid, searchWord },
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching chat list:', error);
    throw error;
  }
};

// 获取聊天记录
export const getChatMessages = async (roomId) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/pollchats/${roomId}/messages`);
    return response.data;
  } catch (error) {
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
request get 
    Searchword(string)
    Range(string)
response (和热门数据返回内容的格式相同)
    roomID (Number), 
    roomName (string)
    roomIntro (string)
    roomOwnerID (Number)
    roomOwnerName (string)
    memberList (List)
        memberId,memberId...
*/
export const search = async (query, scope) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/search`, {
      params: { Searchword: query, Range: scope },
    });
    return response.data;
  } catch (error) {
    console.error('Error searching:', error);
    throw error;
  }
};


/*用户搜索
request get
  searchWord(string)
respone
  UID (number)
  UName(string) 
  URemark(string)
  UImage (string) //好像已经实现了uid同统一头像的功能，所以这一项没有我感觉问题不大
*/
export const searchUsers = async (searchWord) => {
  try {
    const response = await axios.get(`${API_BASE_URL}/users`, {
      params: { searchWord },
    });
    return response.data;
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
export const login = async (email, password) => {
  try {
    const response = await axios.post(`${API_BASE_URL}/login`, {
      Email: email,
      Password: password,
    });
    return response.data;
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
    const response = await axios.get(`${API_BASE_URL}/schedule`, {
      params: { Uid: uid },
    });
    return response.data;
  } catch (error) {
    console.error('Error fetching schedule items:', error);
    alert("api:获取个人信息uid="+uid+"失败");
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