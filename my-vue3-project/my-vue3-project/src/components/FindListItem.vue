<template>
    <body>
        <div class="body">
            <left>
                <a>房间号：{{  }}</a>
                <a>{{  }}的聊天室</a>
            </left>
            <right>
                <img src="src/assets/logo.png"/>
                <img />
                <img />
            </right>
        </div>
    </body>
</template>

<script>
export default {
    name: 'FindListItem',
}
</script>
    


<style scoped>
body{
    height: 100px;
    background-color:#cfe9fc ;
    border-radius: 16px;
    padding: 15px;
}
.body{
    display: flex;
    flex-direction: row;
    justify-content: space-between; /* 子元素之间的间距 */
}
left{
    height: 100px;
    background-color: aquamarine;
    display: flex;
    flex-direction: column;
    align-items: flex-start; /* 左对齐 */
}
right{
    background-color: rgb(188, 58, 58);
    display: flex;
    flex-direction: row;
    align-items: center;
}
img{
    width: 50px; /* 设置图片的宽度 */
    height: 50px; /* 设置图片的高度 */
    border-radius: 50%;
    background-color: aquamarine;

}
</style>