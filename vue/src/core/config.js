/**
 * 网站配置文件
 */
import chalk from "chalk";
// import dayjs, { type Dayjs } from "dayjs"
// import duration from "dayjs/plugin/duration"
// dayjs.extend(duration)

const config = {
  appName: '流量控制分析系统',
  appLogo: 'logo.svg',
  showViteLogo: true,
  logs: [],
}

export const viteLogo = (env) => {
  if (config.showViteLogo) {
    const chalk = require("chalk")
    console.log(
        chalk.cyan(
            `> 您好！欢迎使用 流量控制分析系统！ `
        )
    )
  }
}

export default config
