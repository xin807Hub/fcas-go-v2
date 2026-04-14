import ipaddr from "ipaddr.js"
// IP 比较
export const compareIps = (ip1, ip2) => {
  const addr1 = ipaddr.parse(ip1)
  const addr2 = ipaddr.parse(ip2)

  const addr1Parts = addr1.toByteArray()
  const addr2Parts = addr2.toByteArray()

  for (let i = 0; i < addr1Parts.length; i++) {
    if (addr1Parts[i] < addr2Parts[i]) {
      return -1
    } else if (addr1Parts[i] > addr2Parts[i]) {
      return 1;
    }
  }
  return 0
}

// 组播地址判断
export const isMulticast = (ip) => {
  const addr = ipaddr.parse(ip)
  if (addr.kind() === "ipv4") {
    return addr.range() === "multicast"
  } else if (addr.kind() === "ipv6") {
    return addr.range() === "multicast"
  } else {
    return false
  }
}
