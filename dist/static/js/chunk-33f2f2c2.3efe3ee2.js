(window["webpackJsonp"]=window["webpackJsonp"]||[]).push([["chunk-33f2f2c2"],{"4eab":function(t,e,a){},"769d":function(t,e,a){"use strict";a("4eab")},"817d":function(t,e,a){var s,i,o;(function(r,c){i=[e,a("313e")],s=c,o="function"===typeof s?s.apply(e,i):s,void 0===o||(t.exports=o)})(0,(function(t,e){var a=function(t){"undefined"!==typeof console&&console&&console.error&&console.error(t)};if(e){var s=["#2ec7c9","#b6a2de","#5ab1ef","#ffb980","#d87a80","#8d98b3","#e5cf0d","#97b552","#95706d","#dc69aa","#07a2a4","#9a7fd1","#588dd5","#f5994e","#c05050","#59678c","#c9ab00","#7eb00a","#6f5553","#c14089"],i={color:s,title:{textStyle:{fontWeight:"normal",color:"#008acd"}},visualMap:{itemWidth:15,color:["#5ab1ef","#e0ffff"]},toolbox:{iconStyle:{normal:{borderColor:s[0]}}},tooltip:{backgroundColor:"rgba(50,50,50,0.5)",axisPointer:{type:"line",lineStyle:{color:"#008acd"},crossStyle:{color:"#008acd"},shadowStyle:{color:"rgba(200,200,200,0.2)"}}},dataZoom:{dataBackgroundColor:"#efefff",fillerColor:"rgba(182,162,222,0.2)",handleColor:"#008acd"},grid:{borderColor:"#eee"},categoryAxis:{axisLine:{lineStyle:{color:"#008acd"}},splitLine:{lineStyle:{color:["#eee"]}}},valueAxis:{axisLine:{lineStyle:{color:"#008acd"}},splitArea:{show:!0,areaStyle:{color:["rgba(250,250,250,0.1)","rgba(200,200,200,0.1)"]}},splitLine:{lineStyle:{color:["#eee"]}}},timeline:{lineStyle:{color:"#008acd"},controlStyle:{normal:{color:"#008acd"},emphasis:{color:"#008acd"}},symbol:"emptyCircle",symbolSize:3},line:{smooth:!0,symbol:"emptyCircle",symbolSize:3},candlestick:{itemStyle:{normal:{color:"#d87a80",color0:"#2ec7c9",lineStyle:{color:"#d87a80",color0:"#2ec7c9"}}}},scatter:{symbol:"circle",symbolSize:4},map:{label:{normal:{textStyle:{color:"#d87a80"}}},itemStyle:{normal:{borderColor:"#eee",areaColor:"#ddd"},emphasis:{areaColor:"#fe994e"}}},graph:{color:s},gauge:{axisLine:{lineStyle:{color:[[.2,"#2ec7c9"],[.8,"#5ab1ef"],[1,"#d87a80"]],width:10}},axisTick:{splitNumber:10,length:15,lineStyle:{color:"auto"}},splitLine:{length:22,lineStyle:{color:"auto"}},pointer:{width:5}}};e.registerTheme("macarons",i)}else a("ECharts is not Loaded")}))},8480:function(t,e,a){},8554:function(t,e,a){},"8a3d":function(t,e,a){"use strict";a("8554")},9406:function(t,e,a){"use strict";a.r(e);var s=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",{staticClass:"dashboard-container"},[a("div",{staticClass:"dashboard-text"},[a("div",{staticClass:"dashboard-header"},[a("span",[t._v("您好，欢迎 "+t._s(t.account))])]),t._v(" "),1==t.pch5_status?a("div",{staticClass:"pch5-wrap"},[a("div",{staticClass:"pch5-wrap-header"},[t._v("网页广告概览\n")]),t._v(" "),a("data-card",{attrs:{cardlist:t.cardlist}}),t._v(" "),a("div",[a("chart",{attrs:{chartlist:t.chartlist,num:"pch5"}})],1)],1):t._e(),t._v(" "),2==t.status?a("div",{staticClass:"pch5-wrap"},[a("div",{staticClass:"nopass"},[t._v("\n          账号状态：未激活\n        ")]),t._v(" "),a("el-alert",{staticStyle:{"margin-top":"10px","line-height":"20px"},attrs:{closable:!1,title:"提示",type:"info","show-icon":"",description:"您的渠道资格已提交审核，请耐心等待，审核通过后才能进行其他操作。如有问题请联系客服：kefu@xingchenjia.com"}})],1):t._e(),t._v(" "),3==t.status?a("div",{staticClass:"pch5-wrap"},[a("div",{staticClass:"nopass"},[t._v("\n          账号状态：未激活\n        ")]),t._v(" "),a("el-alert",{staticStyle:{"margin-top":"10px","line-height":"20px"},attrs:{closable:!1,title:"提示",type:"info","show-icon":"",description:"您的渠道资格已提交审核，请耐心等待，审核通过后才能进行其他操作。如有问题请联系客服：kefu@xingchenjia.com"}})],1):t._e(),t._v(" "),a("div",{staticClass:"pch5-wrap"},[a("div",{staticClass:"pch5-wrap-header"},[t._v("余额 "+t._s(t.count)+"\n          ")]),t._v(" "),a("data-card",{attrs:{cardlist:t.cardlist}}),t._v(" "),a("div")],1),t._v(" "),3==t.status?a("div",{staticClass:"pch5-wrap"},[a("div",{staticClass:"pch5-wrap-header"}),t._v(" "),a("data-card",{attrs:{cardlist:t.cardlistcp}})],1):t._e()])])},i=[],o=(a("6b54"),a("db72")),r=a("1bfe"),c=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("el-row",{staticClass:"info-box-wrap",attrs:{gutter:10}},[a("el-col",{staticClass:"info-box-col",attrs:{span:t.cl,shadow:"hover"}},[a("div",{staticClass:"info-box"},[a("div",{staticClass:"info-box-text"},[t._v("今日消费 (USDT)")]),t._v(" "),a("div",{staticClass:"info-box-number"},[a("span",{staticClass:"prespan"},[t._v(t._s(t._f("pre")(t._f("formatNumberRgx")(t.list.yesterday_income))))]),a("span",{staticClass:"endspan"},[t._v(t._s(t._f("end")(t.list.yesterday_income)))])])])]),t._v(" "),a("el-col",{staticClass:"info-box-col",attrs:{span:t.cl,shadow:"hover"}},[a("div",{staticClass:"info-box"},[a("div",{staticClass:"info-box-text"},[t._v("昨日消费 (USDT)")]),t._v(" "),a("div",{staticClass:"info-box-number"},[a("span",{staticClass:"prespan"},[t._v(t._s(t._f("pre")(t._f("formatNumberRgx")(t.list.sevendays_income))))]),a("span",{staticClass:"endspan"},[t._v(t._s(t._f("end")(t.list.sevendays_income)))])])])]),t._v(" "),a("el-col",{staticClass:"info-box-col",attrs:{span:t.cl,shadow:"hover"}},[a("div",{staticClass:"info-box"},[a("div",{staticClass:"info-box-text"},[t._v("本周消费 (USDT)")]),t._v(" "),a("div",{staticClass:"info-box-number"},[a("span",{staticClass:"prespan"},[t._v(t._s(t._f("pre")(t._f("formatNumberRgx")(t.list.month_income))))]),a("span",{staticClass:"endspan"},[t._v(t._s(t._f("end")(t.list.month_income)))])])])]),t._v(" "),a("el-col",{staticClass:"info-box-col",attrs:{span:t.cl,shadow:"hover"}},[a("div",{staticClass:"info-box"},[a("div",{staticClass:"info-box-text"},[t._v("本月消费 (USDT)")]),t._v(" "),a("div",{staticClass:"info-box-number"},[a("span",{staticClass:"prespan"},[t._v(t._s(t._f("pre")(t._f("formatNumberRgx")(t.list.lastmonth_income))))]),a("span",{staticClass:"endspan"},[t._v(t._s(t._f("end")(t.list.lastmonth_income)))])])])])],1)},n=[],l={name:"datacard",props:["cardlist"],filters:{pre:function(t){var e=t.toString(),a=e.indexOf(".");return a>0?e.substring(0,a):e},end:function(t){var e=t.toString(),a=e.indexOf(".");return a>-1?e.substr(a):""},two:function(t){return parseFloat(t).toFixed(2)}},data:function(){return{cl:3,list:{lastmonth_income:"",yesterday_income:"",sevendays_income:"",month_income:"",latelydays_income:""}}},created:function(){this.setlist()},mounted:function(){},watch:{cardlist:{handler:function(t){t&&(this.list=Object(o["a"])(Object(o["a"])({},this.list),t))}}},methods:{setlist:function(){this.cl=parseInt(4.8)}}},d=l,h=(a("8a3d"),a("2877")),p=Object(h["a"])(d,c,n,!1,null,"5552cbf9",null),f=p.exports,m=function(){var t=this,e=t.$createElement,a=t._self._c||e;return a("div",["pch5"==t.num?a("div",{style:{width:"100%",height:"430px"},attrs:{id:"myChart"}}):t._e(),t._v(" "),"meida"==t.num?a("div",{style:{width:"100%",height:"430px"},attrs:{id:"myChartMedia"}}):t._e()])},u=[],_=(a("7f7f"),a("817d"),{name:"chart",props:["chartlist","chartlistcp","num"],components:{},data:function(){return{dialogVisible:!1,datarange:[],isActive:0,chartdata:[],xAxisArr:[],incomeArr:[],web_showArr:[],ecpmArr:[],hits_disArr:[],click_rateArr:[]}},created:function(){var t=new Date,e=t.getTime()-2592e6;e=new Date(e),this.datarange=[e,t]},mounted:function(){this.setdata()},watch:{chartlist:{handler:function(t){t.length>0&&(this.chartdata=this.chartlist,this.setdata())}}},methods:{setdata:function(){this.chartdata=this.chartlist,this.xAxisArr=[],this.incomeArr=[],this.web_showArr=[],this.ecpmArr=[],this.hits_disArr=[],this.click_rateArr=[];for(var t=0;t<this.chartdata.length;t++)this.xAxisArr.push(this.chartdata[t].day),this.incomeArr.push(this.chartdata[t].income),this.web_showArr.push(this.chartdata[t].web_show),this.ecpmArr.push(this.chartdata[t].ecpm),this.hits_disArr.push(this.chartdata[t].hits_dis),this.click_rateArr.push((100*this.chartdata[t].click_rate).toFixed(2));this.drawLine()},drawLine:function(t,e){var a,s,i=document.getElementById("myChart"),o=document.getElementById("myChartMedia");i&&(a=this.$echarts.init(i,"macarons")),o&&(s=this.$echarts.init(o,"macarons"));var r={title:{text:"最近7天统计",left:50,textStyle:{fontSize:16,color:"#666666"}},tooltip:{trigger:"axis",axisPointer:{type:"cross",label:{backgroundColor:"#6a7985"}},formatter:function(t){return 3==t[0].seriesIndex?"时间:"+t[0].name+"</br>"+t[0].seriesName+":"+t[0].value+"%":t[0].name+"</br>"+t[0].seriesName+":"+t[0].value}},legend:{data:["收入","展现量","点击量","点击率","eCPM"],selectedMode:"single",y:"bottom"},toolbox:{feature:{saveAsImage:{}}},grid:{left:"3%",right:"4%",containLabel:!0},xAxis:[{type:"category",boundaryGap:!1,data:this.xAxisArr}],yAxis:[{type:"value",axisLabel:{formatter:function(t){return t}}}],series:[{name:"收入",type:"line",areaStyle:{normal:{type:"default"}},data:this.incomeArr,label:{normal:{show:!0,position:"top"}}},{name:"展现量",type:"line",areaStyle:{normal:{}},data:this.web_showArr,label:{normal:{show:!0,position:"top"}},colorStops:[{offset:0,color:"red"},{offset:1,color:"blue"}],globalCoord:!1},{name:"点击量",type:"line",areaStyle:{normal:{}},data:this.hits_disArr,label:{normal:{show:!0,position:"top"}},colorStops:[{offset:0,color:"red"},{offset:1,color:"blue"}],globalCoord:!1},{name:"点击率",type:"line",areaStyle:{normal:{}},data:this.click_rateArr,label:{normal:{show:!0,position:"top"}},itemStyle:{normal:{label:{show:!0,position:"inside",formatter:"{c}%"}}},colorStops:[{offset:0,color:"red"},{offset:1,color:"blue"}],globalCoord:!1},{name:"eCPM",type:"line",label:{normal:{show:!0,position:"top"}},areaStyle:{normal:{}},data:this.ecpmArr}]};i&&a.setOption(r,!0),o&&s.setOption(r,!0)}}}),b=_,v=(a("d4b8"),Object(h["a"])(b,m,u,!1,null,"826a0d3e",null)),y=v.exports,x=(a("cf45"),a("2f62")),g={components:{dataCard:f,chart:y},data:function(){return{contact:"",cardlist:{},cardlistcp:{},chartlist:[{day:"2",income:2,web_show:2,ecpm:2,hits_dis:9,click_rate:.2},{day:"2",income:2,web_show:2,ecpm:2,hits_dis:9,click_rate:.2},{day:"2",income:2,web_show:2,ecpm:2,hits_dis:9,click_rate:.2},{day:"2",income:2,web_show:2,ecpm:2,hits_dis:9,click_rate:.2},{day:"2",income:2,web_show:2,ecpm:2,hits_dis:9,click_rate:.2},{day:"2",income:2,web_show:2,ecpm:2,hits_dis:9,click_rate:.2},{day:"2",income:2,web_show:2,ecpm:2,hits_dis:9,click_rate:.2}],chartlistcp:[],count:this.Cookies.get("m"),isshow:!1,isedit:!1,isedit1:!1}},created:function(){},computed:Object(o["a"])({},Object(x["b"])(["taurus_status","cp_status","pch5_status","account","status","financestatus"])),methods:{getDay:function(t){var e=new Date,a=e.getTime()+864e5*t;e.setTime(a);var s=e.getFullYear(),i=e.getMonth(),o=e.getDate();return i=this.doHandleMonth(i+1),o=this.doHandleMonth(o),s+"-"+i+"-"+o},doHandleMonth:function(t){var e=t;return 1==t.toString().length&&(e="0"+t),e},getList:function(){var t=this;if(1==this.cp_status){this.getDay(-7),this.getDay(-1);Object(r["b"])().then((function(e){if(200==e.code){var a=e.data;t.chartlistcp=a}})).catch((function(t){console.log(t)})),Object(r["a"])().then((function(e){200==e.code&&(t.cardlistcp=e.data)})).catch((function(t){console.log(t)}))}1!=this.pch5_status&&1!=this.taurus_status||(Object(r["c"])().then((function(e){200==e.code&&(t.chartlist=e.data)})).catch((function(t){console.log(t)})),Object(r["d"])().then((function(e){200==e.code&&(t.cardlist=e.data)})).catch((function(t){console.log(t)})))}}},w=g,C=(a("769d"),Object(h["a"])(w,s,i,!1,null,"9cfed464",null));e["default"]=C.exports},d4b8:function(t,e,a){"use strict";a("8480")}}]);