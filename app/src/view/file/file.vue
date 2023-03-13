<template>
    <div >
        <div class="file-box" @contextmenu.prevent="onContextmenu">
            <div v-clickDown="{ clickFun: 'dblclick', clickFunType: '1111' }">
                <img class="folder" src="../../assets/images/myfile/folder.png" alt="">
                <div class="file-name">文件夹</div>
            </div>
        </div>
        <Modal
            v-model="modal1"
            class-name="vertical-center-modal"
            title="上传文件"
            @on-ok="ok"
            @on-cancel="cancel">
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
    </div>
</template>

<script>
    import Vue from 'vue';
    import SparkMD5 from 'spark-md5';
    import axios from 'axios';
    import Contextmenu from "vue-contextmenujs";
    Vue.use(Contextmenu);
    export default {
        data () {
            return {
                modal1: false,
                fileParam:{
                    hash:"",
                    uid:2
                },
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
                        this.modal1 = true;
                        console.log("此文件夹中上传文件");
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
            ok () {
                console.log(this.fileParam.hash)
                this.$modal1=false;
            },
            cancel () {
                this.$modal1=false;
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
