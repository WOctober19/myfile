import Vue from 'vue'
import router from '../router'
 
const clickDown = Vue.directive('clickDown', {
 
    inserted(el, binding, vnode) {
        let clickTimer = null
          
        // 单击
        // el.addEventListener('click', () => {
        //   if (clickTimer) {
        //     window.clearTimeout(clickTimer);
        //     clickTimer = null;
        //   }
        //   clickTimer = setTimeout(() => {
    
        //     vnode.context[binding.value.clickFun](binding.value.clickFunType);
        //   }, 300);
        // })
    
        // 双击
        el.addEventListener('dblclick', () => {
          if (clickTimer) {
            window.clearTimeout(clickTimer);
            clickTimer = null;
          }
          // this.$router.push({ path:'/home'  })
          console.log(111)
          // vnode.context[binding.value.dblclickFu](binding.value.dblclickFuType);
        })
      },
});
 
export default { clickDown }