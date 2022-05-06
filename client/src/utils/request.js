import axios from 'axios'
export function request(config) {
  const instance = axios.create({
    baseURL: 'http://172.24.129.29:8087',
    // baseURL: '/mock',
    withCredentials: false
  })
  instance.interceptors.request.use(config=>{
    if(sessionStorage.getItem('token')){
      config.headers[ 'Authorization' ] = 'Bearer ' + sessionStorage.getItem('token')
    }
    return config
  },err=>{
    return Promise.reject(error)
  });
  return instance(config)
}
