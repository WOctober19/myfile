import Vue from 'vue'
import router from '../router'
 
const clickDown = Vue.directive('clickDown', {
 
    inserted(el, binding, vnode) {
        let clickTimer = null
    
        // 双击
        el.addEventListener('dblclick', () => {
          if (clickTimer) {
            window.clearTimeout(clickTimer);
            clickTimer = null;
          }
          // console.log(binding.value.clickFunType)
          router.push({ path:'/file/myfile' ,query:{fid:binding.value.clickFunType}})
          // vnode.context[binding.value.dblclickFu](binding.value.dblclickFuType);
        })
      },
});
 
export default { clickDown }