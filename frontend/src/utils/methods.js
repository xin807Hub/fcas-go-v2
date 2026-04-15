export const throttle = function(func, wait, options) {
  var context, args, result
  var timeout = null
  var previous = 0
  if (!options) options = {}
  var later = function() {
    previous = options.leading === false ? 0 : new Date().getTime()
    timeout = null
    result = func.apply(context, args)
    if (!timeout) context = args = null
  }
  return function() {
    var now = new Date().getTime()
    if (!previous && options.leading === false) previous = now
    var remaining = wait - (now - previous)
    context = this
    args = arguments
    if (remaining <= 0 || remaining > wait) {
      if (timeout) {
        clearTimeout(timeout)
        timeout = null
      }
      previous = now
      result = func.apply(context, args)
      if (!timeout) context = args = null
    } else if (!timeout && options.trailing !== false) {
      timeout = setTimeout(later, remaining)
    }
    return result
  }
}
export const debounce = function(func, wait, immediate) {
  var timeout
  return function() {
    var context = this
    var args = arguments
    var later = function() {
      timeout = null
      if (!immediate) func.apply(context, args)
    }
    var callNow = immediate && !timeout
    clearTimeout(timeout)
    timeout = setTimeout(later, wait)
    if (callNow) func.apply(context, args)
  }
}

export const getDataUnit = function(num) {
  if (num / 1024 / 1024 / 1024 / 1024 >= 1) {
    return {
      data: (num / 1024 / 1024 / 1024 / 1024).toFixed(0),
      unit: "1099511627776",
    }
  } else if (num / 1024 / 1024 / 1024 >= 1) {
    return { data: (num / 1024 / 1024 / 1024).toFixed(0), unit: "1073741824" }
  } else if (num / 1024 / 1024 >= 1) {
    return { data: (num / 1024 / 1024).toFixed(0), unit: "1048576" }
  } else if (num / 1024 >= 1) {
    return { data: (num / 1024).toFixed(0), unit: "1024" }
  } else {
    return { data: (num - 0).toFixed(0), unit: "1" }
  }
}

export const getDataByteFormat = function(num, unit = "B", rate = 1000) {
  // if (unit === "bps") {
  // num = num * 8
  // }
  if (num / rate / rate / rate / rate >= 1) {
    return (num / rate / rate / rate / rate).toFixed(0) + "T" + unit
  } else if (num / rate / rate / rate >= 1) {
    return (num / rate / rate / rate).toFixed(0) + "G" + unit
  } else if (num / rate / rate >= 1) {
    return (num / rate / rate).toFixed(0) + "M" + unit
  } else if (num / rate >= 1) {
    return (num / rate).toFixed(0) + "K" + unit
  } else {
    return (num - 0).toFixed(0) + unit
  }
}

export const getMinuteStrings = function(startDate, endDate) {
  const minuteStrings = []
  // 将时间转换为毫秒
  const startMillis = startDate.getTime() + 8 * 3600 * 1000
  const endMillis = endDate.getTime() + 8 * 3600 * 1000

  // 计算分钟数
  let currentMillis = startMillis
  while (currentMillis <= endMillis) {
    const currentDate = new Date(currentMillis)

    // 格式化日期
    const dateString = currentDate
      .toISOString()
      .replace("T", " ")
      .substring(0, 19)

    // 将格式化后的日期字符串添加到数组
    minuteStrings.push(dateString)

    // 增加一分钟
    currentMillis += 60000 // 60秒 * 1000毫秒
  }

  return minuteStrings
}
