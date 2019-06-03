package {{.PackageName}}.utils;

import {{.PackageName}}.common.ResponseStatusEnum;
import {{.PackageName}}.common.ResponseVO;
import org.springframework.http.HttpStatus;

/**
 * @author CHENJIANHUA
 * @date 2019/5/31 12:34
 */
public class ResponseUtil {

    private ResponseUtil() { /* no instance */ }

    /**
     * 快速返回成功
     *
     * @param <T>
     * @return
     */
    public static <T> ResponseVO ok() {
        return new ResponseVO<T>(ResponseStatusEnum.SUCCESS.getErrorCode(),
                I18nMessageUtil.getMessage(ResponseStatusEnum.SUCCESS.getI18nKey(), null));
    }

    public static <T> ResponseVO ok(T result) {
        return new ResponseVO<>(ResponseStatusEnum.SUCCESS.getErrorCode(),
                I18nMessageUtil.getMessage(ResponseStatusEnum.SUCCESS.getI18nKey(), null), result);
    }

    public static <T> ResponseVO ok(String message, T result) {
        return new ResponseVO<>(ResponseStatusEnum.SUCCESS.getErrorCode(),
                I18nMessageUtil.getMessage(message, null),
                result);
    }

    /**
     * 快速返回失败状态
     *
     * @param <T>
     * @return
     */
    public static <T> ResponseVO fail() {
        return new ResponseVO<T>(ResponseStatusEnum.REQUEST_ERROR.getErrorCode(),
                I18nMessageUtil.getMessage(ResponseStatusEnum.REQUEST_ERROR.getI18nKey(), null));
    }

    public static <T> ResponseVO fail(ResponseStatusEnum responseStatusEnum) {
        return new ResponseVO<>(responseStatusEnum.getErrorCode(),
                I18nMessageUtil.getMessage(responseStatusEnum.getI18nKey(), null));
    }

    /**
     * org.springframework.http.HttpStatus
     * 快速返回Http状态
     */
    public static <T> ResponseVO httpStatus(HttpStatus httpStatus, String message) {
        return new ResponseVO<T>(httpStatus.value(),
                I18nMessageUtil.getMessage(message, null));
    }

    public static <T> ResponseVO httpStatus(HttpStatus httpStatus, String message, T result) {
        return new ResponseVO<>(httpStatus.value(),
                I18nMessageUtil.getMessage(message, null),
                result);
    }
}

