package {{.PackageName}}.utils;

import org.apache.commons.lang3.StringUtils;
import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.text.SimpleDateFormat;
import java.util.Calendar;
import java.util.Date;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public final class DateUtil {
    private static final Logger log = LoggerFactory.getLogger(DateUtil.class);

    /**
     * 格式化日期
     * 默认:yyyy-MM-dd HH:mm:ss
     *
     * @param dateStr 时间格式字符串
     * @return 格式化后的日期
     */
    public static Date parseDate(String dateStr) {
        return parseDate(dateStr, DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS);
    }

    /**
     * 格式化日期
     *
     * @param dateStr 时间格式字符串
     * @param pattern DateFormatEnum时间格式
     * @return 格式化后的日期
     */
    public static Date parseDate(String dateStr, DateFormatEnum pattern) {
        if (StringUtils.isBlank(dateStr)) {
            return null;
        }
        try {
            return pattern.getSdf().parse(dateStr);
        } catch (Exception e) {
            log.error("Parse Date error dateStr={}, pattern={}", dateStr, pattern.getDateFormat());
        }
        return null;
    }

    /**
     * 格式化日期
     *
     * @param date Date对象
     * @return 格式化后的日期
     */
    public static String formatDate(Date date) {
        return formatDate(date, DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS);
    }

    /**
     * 格式化日期
     *
     * @param date Date对象
     * @return 格式化后的日期
     */
    public static String formatDate(Date date, DateFormatEnum pattern) {
        return pattern.getSdf().format(date);
    }

    /**
     * 获取当前时间字符串
     */
    public static String getNowDate() {
        return getNowDate(DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS);
    }

    /**
     * 获取当前时间字符串
     *
     * @param pattern 时间格式
     * @return 时间字符串
     */
    public static String getNowDate(DateFormatEnum pattern) {
        return pattern.getSdf().format(new Date());
    }

    /**
     * 获取当天日期（yyyy-MM-dd）
     *
     * @return 当天日期
     */
    public static String getTodayDate() {
        return getNowDate(DateFormatEnum.DATE_YYYY_MM_DD);
    }

    /**
     * 判断一个时间是否在另一个时间之前
     *  before("2019-07-02","2019-07-01",DateFormatEnum.DATE_YYYY_MM_DD) -> 0
     *  before("2019-07-01","2019-07-02",DateFormatEnum.DATE_YYYY_MM_DD) -> 1
     *  error -> null
     * @param dateStr1 第一个时间
     * @param dateStr2 第二个时间
     * @return 判断结果
     */
    public static Integer before(String dateStr1, String dateStr2, DateFormatEnum pattern) {
        Date dateTime1 = parseDate(dateStr1, pattern);
        Date dateTime2 = parseDate(dateStr2, pattern);
        if (null == dateTime1 || null == dateTime2) {
            log.error("compare error dateStr1={}, dateStr2={}, pattern={}", dateStr1, dateStr2, pattern.getDateFormat());
            return null;
        }
        if (dateTime1.before(dateTime2)) {
            return 1;
        } else {
            return 0;
        }
    }

    public static Integer before(String dateStr1, String dateStr2) {
        return before(dateStr1, dateStr2, DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS);
    }

    public static void main(String[] args) {
        System.out.println(getNowDate(DateFormatEnum.DATE_YYYYMMDD));
        System.out.println(getNowDate());
    }
}
