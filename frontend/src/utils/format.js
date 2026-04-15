import { formatTimeToStr } from '@/utils/date'
import { getDict } from '@/utils/dictionary'
import {ref} from "vue";

export const formatBoolean = (bool) => {
  if (bool !== null) {
    return bool ? '是' : '否'
  } else {
    return ''
  }
}
export const formatDate = (time) => {
  if (time !== null && time !== '') {
    var date = new Date(time)
    return formatTimeToStr(date, 'yyyy-MM-dd hh:mm:ss')
  } else {
    return ''
  }
}

export const filterDict = (value, options) => {
  const rowLabel = options && options.filter(item => item.value === value)
  return rowLabel && rowLabel[0] && rowLabel[0].label
}

export const filterDataSource = (dataSource, value) => {
  if (Array.isArray(value)) {
    return value.map(item => {
      const rowLabel = dataSource && dataSource.find(i => i.value === item)
      return rowLabel?.label
    })
  }
  const rowLabel = dataSource && dataSource.find(item => item.value === value)
  return rowLabel?.label
}

export const getDictFunc = async(type) => {
  const dicts = await getDict(type)
  return dicts
}

const path = import.meta.env.VITE_BASE_PATH + ':' + import.meta.env.VITE_SERVER_PORT + '/'
export const ReturnArrImg = (arr) => {
  const imgArr = []
  if (arr instanceof Array) { // 如果是数组类型
    for (const arrKey in arr) {
      if (arr[arrKey].slice(0, 4) !== 'http') {
        imgArr.push(path + arr[arrKey])
      } else {
        imgArr.push(arr[arrKey])
      }
    }
  } else { // 如果不是数组类型
    if (arr.slice(0, 4) !== 'http') {
      imgArr.push(path + arr)
    } else {
      imgArr.push(arr)
    }
  }
  return imgArr
}

export const onDownloadFile = (url) => {
  window.open(path + url)
}
const colorToHex = u=>{
  let e = u.replace("#", "").match(/../g);
  for (let t = 0; t < 3; t++)
    e[t] = parseInt(e[t], 16);
  return e
}

const  hexToColor = (u,e,t)=>{
  let a = [u.toString(16), e.toString(16), t.toString(16)];
  for (let n = 0; n < 3; n++)
    a[n].length === 1 && (a[n] = `0${a[n]}`);
  return `#${a.join("")}`
}
const  generateAllColors = (u,e)=> {
  let t = colorToHex(u);
  const target = [10, 10, 30];
  for (let a = 0; a < 3; a++)
  t[a] = Math.floor(t[a] * (1 - e) + target[a] * e);
  return hexToColor(t[0], t[1], t[2])
}

const generateAllLightColors = (u, e) => {
  let t = colorToHex(u);
  const target = [240, 248, 255]; // RGB for blue white color
  for (let a = 0; a < 3; a++)
    t[a] = Math.floor(t[a] * (1 - e) + target[a] * e);
  return hexToColor(t[0], t[1], t[2]);
}


function addOpacityToColor(u, opacity) {
  let t = colorToHex(u);
  return `rgba(${t[0]}, ${t[1]}, ${ t[2]}, ${opacity})`;
}


export const setBodyPrimaryColor = (  primaryColor, darkMode ) =>{

  let fmtColorFunc = generateAllColors
  if (darkMode === 'light') {
    fmtColorFunc = generateAllLightColors
  }

  document.documentElement.style.setProperty('--el-color-primary', primaryColor)
  document.documentElement.style.setProperty('--el-color-primary-bg', addOpacityToColor(primaryColor, 0.4))
  for (let times = 1; times <= 2; times++) {
    document.documentElement.style.setProperty(`--el-color-primary-dark-${times}`,  fmtColorFunc(primaryColor, times / 10))
  }
  for (let times = 1; times <= 10; times++) {
    document.documentElement.style.setProperty(`--el-color-primary-light-${times}`,  fmtColorFunc(primaryColor, times / 10))
  }
  document.documentElement.style.setProperty(`--el-menu-hover-bg-color`,  addOpacityToColor(primaryColor, 0.2))
}


const baseUrl = ref(import.meta.env.VITE_BASE_API)

export const getBaseUrl = () => {
    return  baseUrl.value === "/" ? "" : baseUrl.value + "/"
}

export const formatByte = (value, suffix) => {
  let unit = "B"
  if (suffix) {
    unit = suffix
  }
  if (value > 1000000000) {
    return (value / 1000000000).toFixed(2) + "(G" + unit + ")"
  } else if (value > 1000000) {
    return (value / 1000000).toFixed(2) + "(M" + unit + ")"
  } else if (value > 1000) {
    return (value / 1000).toFixed(2) + "(K" + unit + ")"
  } else {
    return value ? value.toFixed(2) + "(" + unit + ")" : 0 + "(" + unit + ")"
  }
}

/**
 * 通用格式化大小/速率的函数
 * @param {number|string} value - 要格式化的数值（如字节数、速率等）
 * @param {object} options - 配置项
 * @param {string[]} [options.units=['B','KB','MB','GB','TB']] - 单位数组，从最小到最大
 * @param {number} [options.base=1000] - 单位换算基数（1000 或 1024）
 * @param {number} [options.decimals=2] - 保留的小数位数
 * @param {boolean} [options.showUnit=true] - 是否显示单位
 * @param {string} [options.suffix=''] - 单位后缀，例如 'bps'、'ps'
 * @returns {string} 格式化后的字符串，例如 "12.34(MB)"
 */
export function formatSize(value = 0, options = {}) {
  const {
    units = ['B', 'KB', 'MB', 'GB', 'TB'], // 默认单位
    base = 1000,                           // 换算基数，1000=十进制，1024=二进制
    decimals = 2,                          // 默认保留两位小数
    showUnit = true,                       // 是否显示单位
    suffix = '',                           // 单位后缀，如 "ps", "bps"
  } = options;

  // 确保 value 是数字，若无效则默认设为 0
  const num = Number(value) || 0;

  // 特殊处理：0 直接返回 "0.00(B)"
  if (num === 0) {
    return `0.00${showUnit ? ` (${units[0]}${suffix})` : ''}`;
  }

  // 根据数量级计算应该使用哪个单位（通过 log 取对数）
  const index = Math.min(
      units.length - 1,                             // 防止越界
      Math.floor(Math.log(num) / Math.log(base))   // 得到单位的索引
  );

  // 将原始值缩放到目标单位
  const scaled = num / Math.pow(base, index);

  // 保留指定位数的小数
  const fixed = scaled.toFixed(decimals);

  // 返回带单位或纯数字
  return showUnit
      ? `${fixed} (${units[index]}${suffix})`
      : fixed;
}

export const formatIsOversea = (number) => {
  return (number === 0) ? "国内" : ((number === 1) ? "国外" : "未知")
}
