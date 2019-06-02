package {{.PackageName}}.utils;

import java.util.UUID;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public class UUIDUtil {
    /**
     * 获取UUID
     * @return
     */
    public static String getUUID32() {
        return UUID.randomUUID().toString().replace("-", "").toLowerCase();
    }
}
