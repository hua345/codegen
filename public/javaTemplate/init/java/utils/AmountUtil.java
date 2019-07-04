package {{.PackageName}}.utils;

import org.apache.commons.lang3.StringUtils;

import java.math.BigDecimal;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public class AmountUtil {
    /**
     * 保留小数点后两位
     */
    private final static Integer DEFAULT_SCALE = 2;

    public static String formatBigDecimal(String amountStr) {
        if (StringUtils.isBlank(amountStr)) {
            return null;
        }
        BigDecimal amount = new BigDecimal(amountStr);
        return formatBigDecimal(amount);
    }

    public static String formatBigDecimal(BigDecimal amount) {
        if (null == amount) {
            return null;
        }
        // ROUND_HALF_UP = 四舍五入
        amount = amount.setScale(DEFAULT_SCALE, BigDecimal.ROUND_HALF_UP);
        return amount.toString();
    }

//    public static void main(String[] args) {
//        System.out.println(formatBigDecimal("2000"));
//        System.out.println(formatBigDecimal("12.34560"));
//    }
}
