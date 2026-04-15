import * as echarts from 'echarts'

/**
 * 生成ECharts图表的自动渐变色样式
 * @param {number} count - 需要生成的颜色样式数量
 * @param {Object} options - 配置选项
 * @param {string} options.startColor - 起始颜色 (十六进制颜色，如 #3B82F6)
 * @param {string} options.endColor - 结束颜色 (十六进制颜色，如 #10B981)
 * @param {number} options.gradientTopAlpha - 渐变顶部透明度 (0-1)
 * @param {number} options.gradientBottomAlpha - 渐变底部透明度 (0-1)
 * @param {number} options.shadowAlpha - 阴影透明度 (0-1)
 * @param {number} options.shadowBlur - 阴影模糊程度
 * @param {number} options.lineAlpha - 线条透明度 (0-1)
 * @param {boolean} options.useECharts - 是否使用ECharts生成渐变 (默认true)
 * @returns {Array<{lineColor: string, areaStyle: Object}>} 样式对象数组
 */
export const generateColorStyles = (count = 1, {
    startColor = '#3B82F6', // 起始颜色
    endColor = '#10B981',   // 结束颜色
    gradientTopAlpha = 0.3, // 渐变顶部透明度 (0-1)
    gradientBottomAlpha = 0, // 渐变底部透明度 (0-1)
    shadowAlpha = 0.6, // 阴影透明度 (0-1)
    shadowBlur = 16, // 阴影模糊程度
    lineAlpha = 1, // 线条透明度 (0-1)
    useECharts = true // 是否使用ECharts生成渐变 (默认true)
} = {}) => {
    // 将十六进制颜色转换为RGB数组
    const hexToRGB = (hex) => {
        // 移除#号并处理简写形式 (#FFF -> #FFFFFF)
        const cleanHex = hex.replace('#', '');
        const fullHex = cleanHex.length === 3
            ? cleanHex.split('').map(c => c + c).join('')
            : cleanHex;

        return [
            parseInt(fullHex.substring(0, 2), 16),
            parseInt(fullHex.substring(2, 4), 16),
            parseInt(fullHex.substring(4, 6), 16)
        ];
    };

    // RGB颜色插值函数
    const interpolateRGB = (start, end, t) =>
        start.map((channel, i) => Math.round(channel + (end[i] - channel) * t));

    // 将RGB颜色转换为rgba字符串
    const toRGBA = (rgb, alpha) => `rgba(${rgb.join(', ')}, ${alpha})`;

    // 创建渐变对象函数
    const createGradient = (rgb, topAlpha, bottomAlpha) => {
        if (useECharts) {
            return new echarts.graphic.LinearGradient(0, 0, 0, 1, [
                {offset: 0, color: toRGBA(rgb, topAlpha)},
                {offset: 0.9, color: toRGBA(rgb, bottomAlpha)}
            ], false);
        } else {
            // 返回CSS渐变字符串，适用于非ECharts环境
            return `linear-gradient(to bottom, ${toRGBA(rgb, topAlpha)} 0%, ${toRGBA(rgb, bottomAlpha)} 90%)`;
        }
    };

    // 验证输入参数
    if (count < 1) count = 1;

    // 将十六进制颜色转换为RGB数组
    const startRGB = hexToRGB(startColor);
    const endRGB = hexToRGB(endColor);

    // 生成样式数组
    return Array.from({length: count}, (_, i) => {
        const ratio = count === 1 ? 0 : i / (count - 1);
        const rgb = interpolateRGB(startRGB, endRGB, ratio);

        return {
            lineColor: toRGBA(rgb, lineAlpha),
            areaStyle: {
                color: createGradient(rgb, gradientTopAlpha, gradientBottomAlpha),
                shadowColor: toRGBA(rgb, shadowAlpha),
                shadowBlur,
            }
        };
    });
};
