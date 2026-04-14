/*! 
 Build based on BoycentAdmin 
 Time : 1742993886000 */
import{cP as t}from"./087AC4D233B64EB0index.ec102100.js";import{linkListApi as i}from"./087AC4D233B64EB0link.dc252d2a.js";const a=t("link",{state:()=>({linkData:[]}),actions:{getLink(){return new Promise(((t,a)=>{this.linkData.length>0?t(this.linkData):i({page:1,limit:Number.MAX_SAFE_INTEGER,key:""}).then((i=>{0==i.code?this.linkData=i.data.list:this.linkData=[],t(this.linkData)})).catch((t=>{console.log(t)}))}))}}});export{a as u};
