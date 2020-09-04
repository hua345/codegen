package {{.PackageName}}.utils.holder;

import org.springframework.web.context.request.RequestContextHolder;
import org.springframework.web.context.request.ServletRequestAttributes;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;
import javax.servlet.http.HttpSession;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public class ServletContextHolder {

    private ServletContextHolder() { /* no instance */ }

    /**
     * 通过静态方法获得当前的request对象
     * @return 当前线程对应的request对象
     */
    public static HttpServletRequest request() {
        return ((ServletRequestAttributes) RequestContextHolder.getRequestAttributes()).getRequest();
    }

    /**
     * 通过静态方法获得当前的response对象
     * @return 当前线程对应的response对象
     */
    public static HttpServletResponse response() {
        return ((ServletRequestAttributes) RequestContextHolder.getRequestAttributes()).getResponse();
    }

    /**
     * 通过静态方法获得当前的session对象
     * @return 当前线程对应的session对象
     */
    public static HttpSession session() {
        return request().getSession();
    }
}
