/*! 
 Build based on BoycentAdmin 
 Time : 1742993886000 */
System.register(["./087AC4D233B64EB0index-legacy.bd8574ed.js","./087AC4D233B64EB0link-legacy.e53c0666.js"],(function(n,t){"use strict";var i,e;return{setters:[function(n){i=n.cP},function(n){e=n.linkListApi}],execute:function(){n("u",i("link",{state:function(){return{linkData:[]}},actions:{getLink:function(){var n=this;return new Promise((function(t,i){n.linkData.length>0?t(n.linkData):e({page:1,limit:Number.MAX_SAFE_INTEGER,key:""}).then((function(i){0==i.code?n.linkData=i.data.list:n.linkData=[],t(n.linkData)})).catch((function(n){console.log(n)}))}))}}}))}}}));
