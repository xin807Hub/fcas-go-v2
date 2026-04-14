/*! 
 Build based on BoycentAdmin 
 Time : 1742993886000 */
System.register(["./087AC4D233B64EB0index-legacy.bd8574ed.js"],(function(e,t){"use strict";var i;return{setters:[function(e){i=e.s}],execute:function(){e("deviceListApi",(function(e){return i({url:"/configuration/dimdeviceinfo/list",method:"get",params:e})})),e("createDeviceApi",(function(e){return i({url:"/configuration/dimdeviceinfo/save",method:"post",data:e})})),e("updateDeviceApi",(function(e){return i({url:"/configuration/dimdeviceinfo/update",method:"post",data:e})})),e("deleteDeviceApi",(function(e){return i({url:"/configuration/dimdeviceinfo/delete",method:"post",data:e})}))}}}));
