<style lang="less">
  @import './login.less';
</style>

<template>
  <div class="login">
    <div class="login-con">
      <Card icon="log-in" title="欢迎登录" :bordered="false">
        <div class="form-con">
          <login-form @on-success-valid="handleSubmits"></login-form>
          <p class="login-tip">输入任意用户名和密码即可</p>
        </div>
      </Card>
    </div>
  </div>
</template>

<script>
import LoginForm from '_c/login-form'
import { mapActions } from 'vuex'
import { setToken, getToken, setUser } from "@/libs/util";
import axios from "@/libs/api.request";
export default {
  components: {
    LoginForm
  },
  methods: {
    // ...mapActions([
    //   'handleLogin',
    //   'getUserInfo'
    // ]),
    handleSubmits ({ userName, password }) {
      let _self = this;
      axios
        .request({
          method: "post",
          url: "api/common/login-zhe",
          data: {
            userName: userName,
            password: password,
          },
        })
        .then((res) => {
          console.log(res)
          setUser("USER_M", JSON.stringify(res.data));
          // _self.$router.push({
          //   name: _self.$config.homeName,
          // });
        })
        .catch((res) => {
          console.log(res);
          alert("登陆失败!");
        });
      // this.handleLogin({ userName, password }).then(res => {
      //   this.getUserInfo().then(res => {
      //     this.$router.push({
      //       name: this.$config.homeName
      //     })
      //   })
      // })
    }
  }
}
</script>

<style>

</style>
