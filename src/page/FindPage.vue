<template>
    <body>
        <GuideBar />
        <div class="searchContainer">
            <SearchBox class="searchBox"/>
        </div>
        <div class="basicSelect">
            <button>拼车</button>
            <button>运动</button>
            <button>娱乐</button>
            <button>拼单</button>
            <button>约拍</button>
            <button>其他</button>
        </div>
        <FindListItem />
        <FindListItem />
        <FindListItem />
        <FindListItem />
        <FindListItem />
        <FindListItem />
        <FindListItem />
        <FindListItem />
        <FindListItem />
    </body>
</template>

<script>
import GuideBar from '../components/GuideBar.vue';
import SearchBox from '../components/SearchBox.vue';
import FindListItem from '../components/FindListItem.vue';

export default {
    name: 'FindPage',
    components:{
        GuideBar,
        SearchBox,
        FindListItem,
    },
    props:{
        username:{
            type:String,
            required:true,
            default: 'Guest'
        },
        userID:{
            type:Number,
            required:true,
            default:0
        }
    },
}
</script>
    


<style scoped>
.searchContainer{
    margin-top: 20px;
    display: flex;
    justify-content: center;
}
.basicSelect {
  display: flex;
  justify-content: center;
  margin-top: 20px;
  gap:20px;
}

.basicSelect button {
  background-color: #cceaff;
  border: none;
  border-radius: 4px;
  height: 35px;
  width: 60px;
  font-size: 17px;
  color: #657c8c;
}
.basicSelect button:hover{
    background-color: #657c8c;
    color: #ffffff;
}


</style>