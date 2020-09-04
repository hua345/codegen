package {{.PackageName}}.utils.holder;

import org.springframework.context.ApplicationContext;
import org.springframework.context.ApplicationContextAware;
import org.springframework.stereotype.Component;

import javax.annotation.Nonnull;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Component
public class SpringContextHolder implements ApplicationContextAware {

    private static ApplicationContext ctx;

    @Override
    public void setApplicationContext(@Nonnull ApplicationContext applicationContext) {
        ctx = applicationContext;
    }

    /**
     * 根据 Class 类找到在 Spring 中注册的 Bean
     * @param clazz 类
     * @param <T>   泛型
     * @return spring 中的单例
     */
    public static <T> T getBean(@Nonnull Class<T> clazz) {
        return ctx.getBean(clazz);
    }

    /**
     * 根据注册的 beanName 找到在 Spring 中注册的 Bean
     * @param beanName beanName
     * @return spring 中的单例
     */
    public static Object getBean(String beanName) {
        return ctx.getBean(beanName);
    }

    /**
     * 根据注册的 beanName 和 Class 类型找到在 Spring 中注册的 Bean
     * @param beanName beanName
     * @param clazz    类
     * @return spring 中的单例
     */
    public static <T> T getBean(String beanName, Class<T> clazz) {
        return ctx.getBean(beanName, clazz);
    }
}
