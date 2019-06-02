package {{.PackageName}}.utils;

import {{.PackageName}}.utils.holder.ServletContextHolder;
import {{.PackageName}}.utils.holder.SpringContextHolder;
import org.springframework.context.MessageSource;
import org.springframework.context.NoSuchMessageException;
import org.springframework.stereotype.Component;
import org.springframework.web.servlet.support.RequestContextUtils;

import javax.annotation.Nonnull;
import javax.annotation.Nullable;
import java.util.Locale;

/**
 * @author CHENJIANHUA
 * @date 2019/5/30 10:05
 */
@Component
public class I18nMessageUtil {

    private I18nMessageUtil() { /* no instance */ }

    /**
     * 根据key和参数获取对应的内容信息
     *
     * @param key  在国际化资源文件中对应的key
     * @param args 参数
     * @return 对应的内容信息
     */
    public static String getMessage(@Nonnull String key, @Nullable Object[] args) {
        MessageSource messageSource = SpringContextHolder.getBean(MessageSource.class);
        Locale locale = RequestContextUtils.getLocale(ServletContextHolder.request());
        String message = key;
        try {
            message = messageSource.getMessage(key, args, locale);
        } catch (NoSuchMessageException e) {
            log.error("NoSuchMessageException : {}", key);
        }
        return message;
    }
}