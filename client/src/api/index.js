import {request} from '@/utils/request'

//获取分类为3的所有的帖子
export function getPosts(categoryId) {
  return request({
    url: `/api/user/posts/${categoryId}`,
  })
}

//增加帖子
export function addPosts(data, value) {
  return request({
    url: `/api/user/posts/${value}`,
    method: 'post',
    data
  })
}

//获得回复
export function getComments(postId) {
  return request({
    url: `/api/user/comments/${postId}`,
  })
}


//增加回复
export function addComment(postId, data) {
  return request({
    url: `/api/user/comments/${postId}`,
    method: 'post',
    data
  })
}

//获取个人信息
// export function getProfile(userId) {
//   return request({
//     url: '/api/user/profile',
//     method: 'post',
//     data: {
//       "userId": userId
//     }
//   })
// }
export function getProfile(userId) {
    return request({
      url: `/api/user/profile/${userId}`,
      method: 'get'
    })
  }

//获取数据评论
// export function getPostComment(id) {
//   return request({
//     url: '/api/user/post',
//     method: 'post',
//     data: {
//       id: id
//     }
//   })
// }
export function getPostComment(id) {
  return request({
    url: `/api/user/getcomment/${id}`,
    method: 'get',
  })
}
//注册
export function register(data) {
  return request({
    url: '/api/user/register',
    method: 'post',
    data
  })
}

//登录
export function login(data) {
  return request({
    url: '/api/user/login',
    method: 'post',
    data
  })
}

//获得帖子详情
export function postDetail(postId) {
  return request({
    url: '/api/user/postDetail',
    method: 'get',
    params: {
      postId
    }
  })
}

//页面刷新时，通过cookie判断用户是否登录，设置vuex里面的loginedUser
// export function userIsLogined() {
//   return request({
//     url: '/isLogined',
//     method: 'get'
//   })
// }
export function userIsLogined() {
  return request({
    url: '/api/auth/info',
    method: 'get'
  })
}

//登出
export function signOut() {
  return request({
    url: '/api/user/signout'
  })
}
