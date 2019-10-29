//index.js
//获取应用实例
const app = getApp()

Page({
  data: {
    lastItems: []
  },
  onLoad() {
    wx.request({
      url: 'http://127.0.0.1:8080/nodes/1/lastitem',
      success: res => {
        let items = res.data.data.items
        items.map((item, index) => {
          item.extra = String(Math.round(Number(item.extra)/1000)/10) + '万'
        })
        this.setData({
          lastItems: items
        })
      }
    })
  }
})
