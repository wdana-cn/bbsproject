import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

const store = new Vuex.Store({
  state: {
    logined: false,
    loginedUser: {},
    token:''
  },
  mutations: {
    toggleLoginStatus(state, loginStatus) {
      state.logined = loginStatus
    },
    setLoginedUser(state, userMessage) {
      state.loginedUser = userMessage
    },
    setToken(state,token){
      state.token = token
      sessionStorage.token = token
    },
    delToken(state){
      state.token = ''
      sessionStorage.removeItem("token")
    }

  }
})

export default store
