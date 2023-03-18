<template >
    
        <div style="width:100%;height:100%;" @contextmenu.prevent="onContextmenu">
        <div class="file-box" >
            <div v-clickDown="{ clickFun: 'dblclick', clickFunType: '1111' }" @contextmenu.prevent="onFolder">
                <img class="folder" src="../../assets/images/myfile/folder.png" alt="">
                <div class="file-name">文件夹</div>
            </div>
        </div>
        <Modal
            v-model="uploadFile"
            class-name="vertical-center-modal"
            title="上传文件">
            <Upload
                multiple
                type="drag"
                :before-upload="getFileHash"
                action="http://127.0.0.1:3000/api/user/uploadfile"
                :data="fileParam">
                <div style="padding: 20px 0">
                    <Icon type="ios-cloud-upload" size="52" style="color: #3399ff"></Icon>
                    <p>点击或将文件拖拽到这里上传</p>
                </div>
            </Upload>
        </Modal>

        <Modal
            v-model="Folder"
            title="创建文件夹"
            @on-ok="okFolder"
            @on-cancel="cancelFolder">
            <Form :model="folder" label-position="left" :label-width="100">
                
                <FormItem label="文件夹名" v-if="createF">
                    <Input v-model="folder.name"></Input>
                </FormItem>
                <FormItem label="文件夹备注" v-if="createF">
                    <Input v-model="folder.introduce"></Input>
                </FormItem>
            </Form>
        </Modal>
        </div>
   
</template>

<script>
    import Vue from 'vue';
    import SparkMD5 from 'spark-md5';
    import axios from 'axios'
    import { getUser } from '@/libs/util'
    import Contextmenu from "vue-contextmenujs";
    Vue.use(Contextmenu);
    export default {
        data () {
            return {
                uploadFile: false,
                Folder:false,
                createF:false,
                fileParam:{
                    hash:"",
                    uid:getUser("userId")
                },
                folder:{
                    uid:parseInt(getUser("userId"))
                },
                isclickAll:0,
            }
        }
        ,mounted() {
            // 先把浏览器自带的右键功能屏蔽掉
        this.$nextTick(() => {
            // 禁止右键
            document.oncontextmenu = new Function("event.returnValue=false");
        });
        // 监听页面滚动事件
        window.addEventListener('scroll', this.handleScroll, true)
        },
        beforeDestroy() {
        // 离开页面时，清除页面滚动事件
        window.removeEventListener('scroll', this.handleScroll, true)
        this.$nextTick(() => {
            // 放开对浏览器右键的禁止
            document.oncontextmenu = new Function();
        });
        },
        methods: {
            onContextmenu(event) {
            if(this.isclickAll==1){
                return;
            }
            this.$contextmenu({
                items: [
                    {
                        label: "创建文件夹",
                        onClick: () => {
                            this.Folder = true
                            this.createF = true
                        }
                    }
                ],
                event, // 鼠标事件信息
                customClass: "custom-class", // 自定义菜单 class
                zIndex: 3, // 菜单样式 z-index
                minWidth: 230 // 主菜单最小宽度
            });
            return false;
            },
            onFolder(event) {
                let that=this;
                this.isclickAll=1;
                this.$contextmenu({
                    items: [
                        {
                            label: "重命名",
                            onClick: () => {
                            this.message = "重命名";
                            console.log("重命名");
                            }
                        },
                        {
                            label: "删除",
                            onClick: () => {
                            this.message = "删除";
                            console.log("删除");
                            }
                        },
                        {
                            label: "此文件夹中上传文件",
                            onClick: () => {
                            this.uploadFile = true;
                            console.log("此文件夹中上传文件");
                            }
                        }
                    ],
                    event, // 鼠标事件信息
                    customClass: "custom-class", // 自定义菜单 class
                    zIndex: 3, // 菜单样式 z-index
                    minWidth: 230 // 主菜单最小宽度
                });
                
                setTimeout(function(){
                    that.isclickAll=0;
                },300)
                return false;
            },
            okFolder () {
                axios.post('/api/api/file/folder', this.folder).then(function (response) {
                    console.log(response);
                }).catch(function (error) {
                    console.log(error);
                });
            },
            cancelFolder () {
                this.Folder=false;
            },
            getFileHash(file){
                return new Promise( resolve => {
                    const reader = new FileReader();
                    reader.readAsArrayBuffer(file);
                    reader.onload = ev => {
                    let buffer = ev.target.result,
                        spark = new SparkMD5.ArrayBuffer(),
                        HASH,
                        suffix;
                        spark.append(buffer);
                        HASH = spark.end();
                        suffix = file.name.substring(file.name.lastIndexOf(".") + 1);
                        resolve({
                            buffer,
                            HASH,
                            suffix,
                            filename: `${HASH}.${suffix}`
                        });
                    };
                }).then(res=>{
                    this.fileParam.hash = res.HASH
                }).catch(err=>{
                    return Promise.reject(err);
                });
            },
        }
  
    }
</script>
<style lang="less">
    .file-box{
        display: flex;
    }
    .folder{
        width: 60px;
    }
    .file-name{
        text-align: center;
    }
    .vertical-center-modal{
        display: flex;
        align-items: center;
        justify-content: center;

        .ivu-modal{
            top: 0;
        }
    }
</style>
