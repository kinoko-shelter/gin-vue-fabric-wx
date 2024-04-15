// pages/maps.js

const amapFile = require('../../utils/amap-wx.js');

Page({

  /**
   * 页面的初始数据
   */
  data: {
    tips:{},
    mapMsg: "",
    weatherMsg: "",
    iconPath:"../../image/location.png",
    xiAnCenter: "",
    chargingStation: {}, 
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: async function (options) { // 将 onLoad 函数改为异步函数
    await this.fetchChargingStations(); // 等待获取充电站数据
    this.initializeMap(); // 充电站数据加载完毕后初始化地图和天气数据
  },

   initializeMap: function() {
    var that = this;
    var myAmapFun = new amapFile.AMapWX({ key: '8486d21c23c54a65a83a9bea7ab33d74' });
    console.log("1",this.data.xiAnCenter);
    myAmapFun.getRegeo({
      location: that.data.xiAnCenter,
      success: function (data) {
        getApp().globalData.mapInfo = data[0];
        that.setData({
          mapMsg: data[0]
        });
      },
      fail: function (info) {
        console.log(info);
      }
    });
    myAmapFun.getWeather({
      city: "西安",
      success: function (data) {
        that.setData({
          weatherMsg: data.liveData
        });
      },
      fail: function (info) {
        console.log(info);
      }
    });
  },

  fetchChargingStations: function () {
    var that = this;
    return new Promise((resolve, reject) => {  // 注意确保返回一个新的 Promise 对象
      wx.request({
        url: 'http://129.204.60.53:8080/api/cmapinfo/getCmapinfoList?page=1&pageSize=10',
        method: 'GET',
        header: {
          'X-Token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiN2YyYThlMDEtY2UwYi00NTAwLThkOWUtZWE5NDkzMGMyYWE4IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IueuoeeQhuWRmCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTcxMzM3NjYxOCwibmJmIjoxNzEyNzcxODE4fQ.gNZYF3coToTlz9mei8jYmeZHDz0VpemPIu88kFdIpdY',
          'X-User-Id': '1'
        },
        success: function (res) {
          if (res.data && res.data.data && res.data.data.list) {
            const stations = res.data.data.list;
            const randomStation = stations[Math.floor(Math.random() * stations.length)];
            that.setData({
              xiAnCenter: randomStation.location_x + "," + randomStation.location_y,
              chargingStation: {
                latitude: randomStation.location_y,
                longitude: randomStation.location_x
              }
            }, resolve); // 使用回调参数 `resolve` 确保在 setData 完成后解决 Promise
          } else {
            wx.showToast({
              title: '数据解析错误',
              icon: 'none'
            });
            reject('No data found');
          }
        },
        fail: function (error) {
          wx.showToast({
            title: '获取充电桩数据失败',
            icon: 'none'
          });
          reject(error);
        }
      });
    });
  },
  

  bindInput: function (e) {
    var that = this;
    var keywords = e.detail.value;
    var key = config.Config.key;
    var myAmapFun = new amapFile.AMapWX({ key: '8486d21c23c54a65a83a9bea7ab33d74' });
    myAmapFun.getInputtips({
      keywords: keywords,
      location: this.data.xiAnCenter,
      success: function (data) {
        console.log(data);
        if (data && data.tips) {
          that.setData({
            tips: data.tips
          });
        }

      }
    })
  },
  bindSearch: function (e) {

    wx.navigateTo({
      url: '../search/search'
    })

  },
  toAdmin:function(){
    wx.navigateTo({
      url: '../admin/admin'
    })

  },
  toLocation:function(){
    this.mapCtx.moveToLocation()
  },
  toLoadline:function(){
    wx.navigateTo({
      url: '../navigation/navigation?xiAnCenter='+this.data.xiAnCenter,
    })

  },


  onReady: function () {
    
    this.mapCtx = wx.createMapContext('myMap')
  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
  
  },

  /**
   * 生命周期函数--监听页面隐藏
   */
  onHide: function () {
  
  },

  /**
   * 生命周期函数--监听页面卸载
   */
  onUnload: function () {
  
  },

  /**
   * 页面相关事件处理函数--监听用户下拉动作
   */
  onPullDownRefresh: function () {
  
  },

  /**
   * 页面上拉触底事件的处理函数
   */
  onReachBottom: function () {
  
  },

  /**
   * 用户点击右上角分享
   */
  onShareAppMessage: function () {
    return {
      title: '导航lite，地图导航、路线导航、公交地铁换乘方案',
      path: '/pages/index/index',
      success: function (res) {
        // 转发成功
        wx.showToast({
          title: '分享成功',
          icon: 'success',
          duration: 1000
        })
      },
      fail: function (res) {
        // 转发失败
        wx.showToast({
          title: '分享取消',
          icon: 'success',
          duration: 1000
        })
      }
    }
  }
})