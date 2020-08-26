package {{.PackageName}}.utils;

import java.text.SimpleDateFormat;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public enum DateFormatEnum {
    /**
     * 时间格式
     */
    DATE_YYYY_MM_DD_HH_MM_SS("yyyy-MM-dd HH:mm:ss"),
    DATE_YYYY_MM_DD("yyyy-MM-dd"),
    DATE_YYYYMMDD("yyyyMMdd"),
    DATE_YYYY_MM_DD_CN("yyyy年MM月dd日"),
    DATE_YYYY_MM_DD_HH_MM_SS_CN("yyyy年MM月dd日 HH:mm:ss"),
    TIME_HH_MM("HH:mm"),
    TIME_HH_MM_SS("HH:mm:ss"),
    DATE_YYYY_MM_DD_HH_MM("yyyy-MM-dd HH:mm"),
    DATE_YYYYMM("yyyyMM"),
    DATE_YYYY_MM("yyyy-MM"),
    DATE_YYYY_MM_CN("yyyy年MM月"),
    DATE_YYYY_MM_DD_HH_MM_CN("yyyy年MM月dd日 HH:mm");

    private String dateFormat;
    private DateTimeFormatter sdf;

    DateFormatEnum(String dateFormat) {

        this.dateFormat = dateFormat;
        this.sdf = DateTimeFormatter.ofPattern(dateFormat, Locale.CHINA);
    }

    public String getDateFormat() {
        return dateFormat;
    }

    public DateTimeFormatter getSdf() {
        return sdf;
    }
}
