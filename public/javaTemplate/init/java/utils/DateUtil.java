package {{.PackageName}}.utils;

import org.slf4j.Logger;
import org.slf4j.LoggerFactory;

import java.time.*;
import java.time.temporal.TemporalAdjusters;
import java.util.Date;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public final class DateUtil {
    private static final Logger log = LoggerFactory.getLogger(DateUtil.class);
    /**
     * 从`Java 8`开始,`java.time`包提供了新的日期和时间API,主要涉及的类型有:
     *
     * - 本地日期和时间:`LocalDateTime`,`LocalDate`,`LocalTime`；
     * - 带时区的日期和时间:`ZonedDateTime`;
     * - 时刻:`Instant`;
     * - 时区:`ZoneId`,`ZoneOffset`;
     * - 时间间隔:`Duration`
     * 以及一套新的用于取代`SimpleDateFormat`的格式化类型`DateTimeFormatter`
     */
    /**
     * Date 转 localDate
     */
    public static LocalDate date2LocalDate(Date date) {
        Instant instant = date.toInstant();
        ZonedDateTime zdt = instant.atZone(ZoneId.systemDefault());
        return zdt.toLocalDate();
    }

    /**
     * Date 转 localDate
     */
    public static LocalDateTime date2LocalDateTime(Date date) {
        Instant instant = date.toInstant();
        ZonedDateTime zdt = instant.atZone(ZoneId.systemDefault());
        return zdt.toLocalDateTime();
    }

    /**
     * localDate转Date
     */
    public static Date localDate2Date(LocalDate localDate) {
        ZonedDateTime zonedDateTime = localDate.atStartOfDay(ZoneId.systemDefault());
        Instant instant1 = zonedDateTime.toInstant();
        Date from = Date.from(instant1);
        return from;
    }

    /**
     * 解析日期
     * 默认:yyyy-MM-dd HH:mm:ss
     * LocalDateTime LocalDate + LocalTime 两部分都得有
     */
    public static LocalDateTime parseDateTime(String dateStr) {
        return LocalDateTime.parse(dateStr, DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS.getSdf());
    }

    /**
     * 解析日期
     * 默认:yyyy-MM-dd
     * LocalDateTime LocalDate + LocalTime 两部分都得有
     */
    public static LocalDate parseDate(String dateStr) {
        return LocalDate.parse(dateStr, DateFormatEnum.DATE_YYYY_MM_DD.getSdf());
    }

    /**
     * 格式化时间戳
     */
    public static LocalDateTime parseTimestamp(Long timestamp) {
        return LocalDateTime.ofInstant(Instant.ofEpochMilli(timestamp), ZoneId.systemDefault());
    }

    /**
     * 格式化日期
     */
    public static String formatDateTime(Date date) {
        return formatDateTime(date, DateFormatEnum.DATE_YYYY_MM_DD_HH_MM_SS);
    }

    /**
     * 格式化日期
     */
    public static String formatDateTime(Date date, DateFormatEnum pattern) {
        return formatDateTime(date2LocalDateTime(date), pattern);
    }

    /**
     * 格式化日期
     */
    public static String formatDateTime(LocalDateTime localDateTime, DateFormatEnum pattern) {
        return localDateTime.format(pattern.getSdf());
    }

    /**
     * 格式化日期
     */
    public static String formatDateTime(Long timestamp, DateFormatEnum pattern) {
        return parseTimestamp(timestamp).format(pattern.getSdf());
    }

    /**
     * 从性能上System.currentTimeMillis()大于Instant.now().toEpochMilli();
     * 100万次
     * getEpochMilliWithSystem cost:5ms
     * getEpochMilliWithInstant cost:60ms
     */
    public static Long getTimeStamp() {
        return System.currentTimeMillis();
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
        LocalDateTime now = LocalDateTime.now();
        return now.format(pattern.getSdf());
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
     * 获取月第一天
     */
    public static Date getStartDayOfMonth(String date) {
        LocalDate now = LocalDate.parse(date);
        return getStartDayOfMonth(now);
    }

    public static Date getStartDayOfMonth(LocalDate date) {
        LocalDate now = date.with(TemporalAdjusters.firstDayOfMonth());
        return localDate2Date(now);
    }

    public static Date getStartDayOfMonth() {
        return getStartDayOfMonth(LocalDate.now());
    }

    /**
     * 获取月最后一天
     */
    public static Date getEndDayOfMonth(String date) {
        LocalDate localDate = LocalDate.parse(date);
        return getEndDayOfMonth(localDate);
    }

    public static Date getEndDayOfMonth(Date date) {
        return getEndDayOfMonth(date2LocalDate(date));
    }

    public static Date getEndDayOfMonth(LocalDate date) {
        LocalDate now = date.with(TemporalAdjusters.lastDayOfMonth());
        Date.from(now.atStartOfDay(ZoneId.systemDefault()).plusDays(1L).minusNanos(1L).toInstant());
        return localDate2Date(now);
    }

    public static Date getEndDayOfMonth() {
        return getEndDayOfMonth(LocalDate.now());
    }

    /**
     * 一天的开始
     */
    public static LocalDateTime getStartOfDay(LocalDate date) {
        LocalDateTime time = LocalDateTime.of(date, LocalTime.MIN);
        return time;
    }

    public static LocalDateTime getStartOfDay() {
        return getStartOfDay(LocalDate.now());
    }

    /**
     * 一天的结束
     */
    public static LocalDateTime getEndOfDay(LocalDate date) {
        LocalDateTime time = LocalDateTime.of(date, LocalTime.MAX);
        return time;
    }

    public static LocalDateTime getEndOfDay() {
        return getEndOfDay(LocalDate.now());
    }

    public static void main(String[] args) {
        System.out.println(getNowDate());
        System.out.println(formatDateTime(new Date()));
    }
}
