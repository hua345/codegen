package {{.PackageName}}.utils;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public class SnowFlakeUtil {
    private static SnowFlake snowFlakeInstance = null;

    private SnowFlakeUtil() {
    }

    public static long getNextId() {
        if (null == snowFlakeInstance) {
            synchronized (SnowFlakeUtil.class) {
                if (null == snowFlakeInstance) {
                    snowFlakeInstance = new SnowFlake(2, 3);
                }
            }
        }
        return snowFlakeInstance.nextId();
    }

//    public static void main(String[] args) {
//        for (int i = 0; i < (1 << 12); i++) {
//            System.out.println(SnowFlakeUtil.getNextId());
//        }
//    }
}
