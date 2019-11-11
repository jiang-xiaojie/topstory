// 云函数入口文件
const cloud = require('wx-server-sdk')

cloud.init({
  // API 调用都保持和云函数当前所在环境一致
  env: cloud.DYNAMIC_CURRENT_ENV
})

const request = require('request')

// 云函数入口函数
exports.main = async (event, context) => {
  return new Promise((RES, REJ) => {
    request(event.options, (err, res, body) => {
      if (err) return REJ(err)
      RES(res)
    })
  })
}