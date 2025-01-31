// pages/navigation/navigation.js
const appData = getApp().globalData;
const amapFile = require('../../utils/amap-wx.js');
Page({

  /**
   * 页面的初始数据
   */
  data: {
    startPointBean: true,
    endPointBean: true,
    hasanimate: '',
    toolData: ["充电桩"],
    _num: 0,
    changeData: '',
    hasHistory: true,
    historyRoute: '',
    homePart: '',
    companyPart: '',
    xiAnCenter: "",
    chargingStations: [],
    latitude:"",
    longitude:"",
  },

  /**
   * 生命周期函数--监听页面加载
   */
  onLoad: function (options) {
    console.log("options",options);
    this.processLocation(options.xiAnCenter)
    this.setData({
      // changeData: options,
      xiAnCenter:options.xiAnCenter
    });
    this.fetchChargingStations();
    // this.chooseEndPoint();
  },
  // 假设您的页面逻辑包含以下函数
  processLocation: function(locationString) {
  var parts = locationString.split(',');  // 使用逗号分隔字符串
  if(parts.length === 2) {
    var latitude = parts[1].trim();  // 清除可能的空格
    var longitude = parts[0].trim();
    this.getLocationName(latitude, longitude).then(name => {
      this.setData({
        'changeData.startName': name,
        latitude:latitude,
        longitude:longitude,// 更新页面数据
      });
      wx.showToast({
        title: '地点名称: ' + name,
        icon: 'none'
      });
    }).catch(error => {
      console.error('获取地点名称失败:', error);
      wx.showToast({
        title: '获取地点名称失败',
        icon: 'none'
      });
    });
  } else {
    wx.showToast({
      title: '经纬度格式错误',
      icon: 'none'
    });
  }
},

  getLocationName: function(latitude, longitude) {
    return new Promise((resolve, reject) => {
      var myAmapFun = new amapFile.AMapWX({key: '8486d21c23c54a65a83a9bea7ab33d74'});
      myAmapFun.getRegeo({
        location: `${longitude},${latitude}`,
        success: (data) => {
          if(data[0] && data[0].name) {
            var decodedName = decodeURIComponent(data[0].name);
            console.log(data[0].name); 
            resolve(decodedName);     
          } else {
            reject("No name found");  
          }
        },
        fail: (info) => {
          console.log(info);           
          reject("Failed to get location name"); 
        }
      });
    });
  },  
  fetchChargingStations: function () {
    var that = this;
    var parts = that.data.xiAnCenter.split(',');  // 使用逗号分隔字符串
      var latitude = parts[1].trim();  // 清除可能的空格
      var longitude = parts[0].trim();
    const url = `http://8.130.15.61:8080/api/cmapinfo/getCmapinfoList?page=1&pageSize=10&location_x=${latitude}&location_y=${longitude}`;
    console.log(url);
    wx.request({
      url: url,
      method: 'GET',
      header: {
        'X-Token': 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJVVUlEIjoiN2YyYThlMDEtY2UwYi00NTAwLThkOWUtZWE5NDkzMGMyYWE4IiwiSUQiOjEsIlVzZXJuYW1lIjoiYWRtaW4iLCJOaWNrTmFtZSI6IueuoeeQhuWRmCIsIkF1dGhvcml0eUlkIjo4ODgsIkJ1ZmZlclRpbWUiOjg2NDAwLCJpc3MiOiJxbVBsdXMiLCJhdWQiOlsiR1ZBIl0sImV4cCI6MTcxMzM3NjYxOCwibmJmIjoxNzEyNzcxODE4fQ.gNZYF3coToTlz9mei8jYmeZHDz0VpemPIu88kFdIpdY',
        'X-User-Id': '1'
      },
      success: function (res) {
        console.log("res",res);
        const data = res.data.data || {};
        const paths = data.path || [];
        const lastPosition = paths[paths.length - 1] || { lat: 0, lng: 0 }; // Default to {lat: 0, lng: 0} if no path
        console.log("res",res);
        that.setData({
          chargingStations: [{
            name: that.data.changeData.startName,
            latitude: lastPosition.lat,
            longitude: lastPosition.lng
          }]
        });
      },
      fail: function () {
        wx.showToast({
          title: '获取充电桩数据失败',
          icon: 'none'
        });
      }
    });
  },
  

  chooseStartPoint: function () {
    var that = this;
    wx.chooseLocation({
      success: function (res) {
        if (res.name) {
          that.setData({
            changeData: {
              ...that.data.changeData,
              startName: "当前位置",
              startPoint: `${res.longitude},${res.latitude}`,
            },
            startPointBean: true
          });
        } else {
          wx.showToast({
            title: '请选择正确的地址',
            icon: 'none',
            duration: 2000
          });
        }
      }
    });
  },

  chooseStartPoint: function () {
    var that = this;
    wx.chooseLocation({
      success: function (res) {
        if (res.name) {
          that.setData({
            changeData: {
              ...that.data.changeData,
              startName: res.name,
              startPoint: `${res.longitude},${res.latitude}`,
            },
            startPointBean: true
          });
        } else {
          wx.showToast({
            title: '请选择正确的地址',
            icon: 'none',
            duration: 2000
          });
        }
      }
    });
  },
  chooseEndPoint: function () {
    if (this.data.chargingStations.length > 0) {
        const randomIndex = Math.floor(Math.random() * this.data.chargingStations.length);
        const station = this.data.chargingStations[randomIndex];

        this.setData({
            changeData: {
                ...this.data.changeData,
                endName: station.name,
                endPoint: `${station.longitude},${station.latitude}`
            },
            endPointBean: true
        });
    } else {
        wx.showToast({
            title: '充电桩数据未加载，请稍后重试',
            icon: 'none',
            duration: 2000
        });
    }
},

goTo: function () {
  var that = this;
  if (!this.data.startPointBean) {
      wx.showToast({
          title: '请输入起点',
          icon: 'none',
          duration: 1500
      });
      return; 
  }
  if (this.data.endPointBean) {
      // Automatically select a charging station if endpoint is not set
      if (this.data.chargingStations.length > 0) {
          const randomIndex = Math.floor(Math.random() * this.data.chargingStations.length);
          const station = this.data.chargingStations[randomIndex];
          console.log("station",station);
          this.setData({
              changeData: {
                  ...this.data.changeData,
                  endName: "充电桩A",
                  endPoint: `${station.longitude},${station.latitude}`
              },
              endPointBean: true
          }, () => {
              that.navigateToRoute();
          });
      } else {
          wx.showToast({
              title: '充电桩数据未加载，请稍后重试',
              icon: 'none',
              duration: 1500
          });
          return; // Stop further execution if no charging stations are loaded
      }
  } else {
      // If endpoint is already set, navigate immediately
      this.navigateToRoute();
  }
},

navigateToRoute: function () {
  // Centralized navigation logic
  wx.navigateTo({
      url: "../route/route?startName=" + encodeURIComponent(this.data.changeData.startName) + "&startPoint=" + this.data.changeData.startPoint + "&endName=" + encodeURIComponent(this.data.changeData.endName) + "&endPoint=" + this.data.changeData.endPoint + "&ways=" + this.data._num,
  });
  this.storageHistory();
},

  storageHistory: function () {
    const newRoute = {
      startName: this.data.changeData.startName,
      startPoint: this.data.changeData.startPoint,
      endName: this.data.changeData.endName,
      endPoint: this.data.changeData.endPoint,
      ways: this.data._num
    };

    const exists = appData.historyRoute.some(route =>
      route.startPoint === newRoute.startPoint &&
      route.endPoint === newRoute.endPoint &&
      route.ways === newRoute.ways);

    if (!exists) {
      appData.historyRoute.unshift(newRoute);
      this.setData({
        historyRoute: appData.historyRoute,
        hasHistory: true
      });
    }
  },
  changePoint:function(){
    this.setData({
      hasanimate:'animate'
    })
    var that = this
    setTimeout(function(){
      that.setData({
        hasanimate: ''
      })
    },300)
    if (this.data.endPointBean && this.data.startPointBean) {
      this.setData({
        changeData: {
          startName: this.data.changeData.endName,
          startPoint: this.data.changeData.endPoint,
          endName: this.data.changeData.startName,
          endPoint: this.data.changeData.startPoint,
        },

      })

    } else if (this.data.startPointBean){
      this.setData({
        changeData: {
          startName: '输入起点',
          startPoint: '',
          endName: this.data.changeData.startName,
          endPoint: this.data.changeData.startPoint,
        },
        endPointBean: true,
        startPointBean: false
      })

    } else if (this.data.endPointBean){
      this.setData({
        changeData: {
          startName: this.data.changeData.endName,
          startPoint: this.data.changeData.endPoint,
          endName: '输入终点',
          endPoint: "",
        },
        startPointBean: true,
        endPointBean: false
      })
    }
  },
  changetool:function(e){

    this.setData({
      _num: e.target.dataset.num
    })

  },
  chooseEndPoint:function(){
    var that = this
    wx.chooseLocation({
      type: 'gcj02', // 返回可以用于wx.openLocation的经纬度
      success(res) {

        if (res.name != '') {
          
          if (that.data.startPointBean){
            wx.navigateTo({
              url: "../route/route?startName=" + that.data.changeData.startName + "&startPoint=" + that.data.changeData.startPoint + "&endName=" + res.name + "&endPoint=" + res.longitude + ',' + res.latitude + "&ways=" + that.data._num
            })

            that.setData({
              changeData: {
                startName: that.data.changeData.startName,
                startPoint: that.data.changeData.startPoint,
                endName: res.name,
                endPoint: res.longitude + ',' + res.latitude,
              },
              endPointBean: true

            })

            that.storageHistory();


          }else{
            that.setData({
              changeData: {
                endName: res.name,
                endPoint: res.longitude + ',' + res.latitude,

              },
            })


          }

          

        } else {
          wx.showToast({
            title: '请选择正确的地址',
            icon: 'none',
            duration: 1500
          })

        }

      }
    })

  },
  // goTo:function(){

  //   if (this.data.endPointBean && this.data.startPointBean){
  //     wx.navigateTo({
  //       url: "../route/route?startName=" + this.data.changeData.startName + "&startPoint=" + this.data.changeData.startPoint + "&endName=" + this.data.changeData.endName + "&endPoint=" + this.data.changeData.endPoint + "&ways=" + this.data._num,
  //     })

  //     this.storageHistory();

  //   }else{

  //     if (this.data.startPointBean == false){
  //       wx.showToast({
  //         title: '请输入起点',
  //         icon: 'none',
  //         duration: 1500
  //       })

  //     }
  //     if (this.data.endPointBean == false) {
  //       wx.showToast({
  //         title: '请输入终点',
  //         icon: 'none',
  //         duration: 1500
  //       })

  //     }

  //   }


  // },
  emptyHistory: function () {
    var that = this
    wx.showModal({
      title: '提示',
      content: '是否清空历史纪录',
      success(res) {
        if (res.confirm) {
          that.setData({
            historyRoute: '',
            hasHistory: false
          })
          appData.historyRoute = [];
        } else if (res.cancel) {

        }
      }
    })


  },

  setupHome:function(){
    var that = this
    if (this.data.homePart.homeText == "设置一个地址"){
      wx.chooseLocation({
        type: 'gcj02', // 返回可以用于wx.openLocation的经纬度
        success(res) {

          if (res.name != '') {

            appData.homePart.homeText = res.name
            appData.homePart.homePoint = res.longitude + ',' + res.latitude


          } else {
            wx.showToast({
              title: '请选择正确的地址',
              icon: 'none',
              duration: 1500
            })

          }

        }
      })

    }else{

      wx.showActionSheet({
        itemList: ['从这出发', '到这里去', '删除'],
        success(res) {
          switch (res.tapIndex) {
            case 0:

              if ( that.data.changeData.endPoint == undefined ){
                that.setData({
                  changeData:{
                    startName: that.data.homePart.homeText,
                    startPoint: that.data.homePart.homePoint
                  }

                })

              }else{
                wx.navigateTo({
                  url: "../route/route?startName=" + that.data.homePart.homeText + "&startPoint=" + that.data.homePart.homePoint + "&endName=" + that.data.changeData.endName + "&endPoint=" + that.data.changeData.endPoint + "&ways=" + that.data._num,
                })

                that.setData({
                  changeData: {
                    startName: that.data.homePart.homeText,
                    startPoint: that.data.homePart.homePoint,
                    endName: that.data.changeData.endName,
                    endPoint: that.data.changeData.endPoint
                  },
                  startPointBean: true
                })

                that.storageHistory();

              }

              break;
            case 1:

              if ( that.data.changeData.startPoint == undefined ) {
                that.setData({
                  changeData: {
                    endName: that.data.homePart.homeText,
                    endPoint: that.data.homePart.homePoint,

                  }

                })

              } else {
                
                wx.navigateTo({
                  url: "../route/route?startName=" + that.data.changeData.startName + "&startPoint=" + that.data.changeData.startPoint + "&endName=" + that.data.homePart.homeText + "&endPoint=" + that.data.homePart.homePoint + "&ways=" + that.data._num,
                })

                that.setData({

                  changeData: {
                    startName: that.data.changeData.startName,
                    startPoint: that.data.changeData.startPoint,
                    endName: that.data.homePart.homeText,
                    endPoint: that.data.homePart.homePoint
                  },
                  endPointBean: true

                })

                that.storageHistory();

              }

              break;
            case 2:

              that.setData({
                homePart:{
                  homeText: "设置一个地址",
                  homePoint: ""
                }

              })
              appData.homePart.homeText = "设置一个地址";
              appData.homePart.homePoint = "";

          }
        },

      })

    }
    

  },

  setupCompany: function () {
    var that = this
    if (this.data.companyPart.companyText == "设置一个地址") {
      wx.chooseLocation({
        type: 'gcj02', // 返回可以用于wx.openLocation的经纬度
        success(res) {

          if (res.name != '') {

            appData.companyPart.companyText = res.name
            appData.companyPart.companyPoint = res.longitude + ',' + res.latitude


          } else {
            wx.showToast({
              title: '请选择正确的地址',
              icon: 'none',
              duration: 1500
            })

          }

        }
      })

    } else {

      wx.showActionSheet({
        itemList: ['从这出发', '到这里去', '删除'],
        success(res) {
          switch (res.tapIndex) {
            case 0:

              if (that.data.changeData.endPoint == undefined) {
                that.setData({
                  changeData: {
                    startName: that.data.companyPart.companyText,
                    startPoint: that.data.companyPart.companyPoint
                  }

                })

              } else {
                wx.navigateTo({
                  url: "../route/route?startName=" + that.data.homePart.homeText + "&startPoint=" + that.data.companyPart.companyPoint + "&endName=" + that.data.changeData.endName + "&endPoint=" + that.data.changeData.endPoint + "&ways=" + that.data._num,
                })

                that.setData({
                  changeData: {
                    startName: that.data.companyPart.companyText,
                    startPoint: that.data.companyPart.companyPoint,
                    endName: that.data.changeData.endName,
                    endPoint: that.data.changeData.endPoint
                  },
                  startPointBean: true
                })

                that.storageHistory();

              }

              break;
            case 1:

              if (that.data.changeData.startPoint == undefined) {
                that.setData({
                  changeData: {
                    endName: that.data.companyPart.companyText,
                    endPoint: that.data.companyPart.companyPoint,

                  }

                })

              } else {

                wx.navigateTo({
                  url: "../route/route?startName=" + that.data.changeData.startName + "&startPoint=" + that.data.changeData.startPoint + "&endName=" + that.data.companyPart.companyText + "&endPoint=" + that.data.companyPart.companyPoint + "&ways=" + that.data._num,
                })

                that.setData({

                  changeData: {
                    startName: that.data.changeData.startName,
                    startPoint: that.data.changeData.startPoint,
                    endName: that.data.companyPart.companyText,
                    endPoint: that.data.companyPart.companyPoint
                  },
                  endPointBean: true

                })

                that.storageHistory();

              }

              break;
            case 2:

              that.setData({
                companyPart: {
                  companyText: "设置一个地址",
                  companyPoint: ""
                }

              })
              appData.companyPart.companyText = "设置一个地址";
              appData.companyPart.companyPoint = "";

          }
        },

      })

    }


  },


  storageHistory:function(){
    //历史纪录
    if (appData.historyRoute != '') {
      var bean = true;

      for (var i = 0; i < appData.historyRoute.length; i++) {

        if (appData.historyRoute[i].startPoint == this.data.changeData.startPoint && appData.historyRoute[i].endPoint == this.data.changeData.endPoint && appData.historyRoute[i].ways == this.data._num){
          bean = false

        }

      }

      if(bean){
        appData.historyRoute.unshift({
          startName: this.data.changeData.startName,
          startPoint: this.data.changeData.startPoint,
          endName: this.data.changeData.endName,
          endPoint: this.data.changeData.endPoint,
          ways: this.data._num
        })

      }

    } else {
      appData.historyRoute.unshift({
        startName: this.data.changeData.startName,
        startPoint: this.data.changeData.startPoint,
        endName: this.data.changeData.endName,
        endPoint: this.data.changeData.endPoint,
        ways: this.data._num

      })

    }


  },
  onReady: function () {

  },

  /**
   * 生命周期函数--监听页面显示
   */
  onShow: function () {
    if (appData.historyRoute != '') {
      this.setData({
        hasHistory: true,

      })

    } else {
      this.setData({
        hasHistory: false,

      })

    }
    this.setData({
      
      historyRoute: appData.historyRoute,
      homePart: appData.homePart,
      companyPart: appData.companyPart
    })
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