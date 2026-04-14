/*! 
 Build based on BoycentAdmin 
 Time : 1742993886000 */
System.register(["./087AC4D233B64EB0index-legacy.bd8574ed.js"],(function(t,e){"use strict";var n;return{setters:[function(t){n=t.a$}],execute:function(){t("i",(function(t,e){if(!n||!t||!e)return!1;var i,c=t.getBoundingClientRect();return i=e instanceof Element?e.getBoundingClientRect():{top:0,right:window.innerWidth,bottom:window.innerHeight,left:0},c.top<i.bottom&&c.bottom>i.top&&c.right>i.left&&c.left<i.right})),t("g",(function(t){var e,n;return"touchend"===t.type?(n=t.changedTouches[0].clientY,e=t.changedTouches[0].clientX):t.type.startsWith("touch")?(n=t.touches[0].clientY,e=t.touches[0].clientX):(n=t.clientY,e=t.clientX),{clientX:e,clientY:n}}))}}}));
