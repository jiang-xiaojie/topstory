//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    lastItems: [
      {
        'title': '妻子的浪漫旅行 杜江霍思燕',
        'description': ''
      }
    ]
  },
  onLoad() {
    wx.request({
      url: 'http://127.0.0.1:8080/nodes/1/lastitem',
      success: res => {
        let items = res.data.data.items
        this.setData({
          lastItems: items
        })
      }
    })
  }
})
