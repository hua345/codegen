package {{.PackageName}}.common;

import com.alibaba.fastjson.JSONObject;

import java.io.Serializable;

import lombok.Getter;
import lombok.Setter;
import org.springframework.http.HttpStatus;
/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Getter
@Setter
public class ResponseVO<T> implements Serializable {
    private static final long serialVersionUID = 1L;
    /**
     * 返回状态码
     */
    private Integer code;

    /**
     * 返回消息
     */
    private String msg;

    /**
     * 返回内容
     */
    private T data;

    public ResponseVO(Integer code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    public ResponseVO(Integer code, String msg, T data) {
        this.code = code;
        this.msg = msg;
        this.data = data;
    }

    @Override
    public String toString() {
        return JSONObject.toJSONString(this);
    }

    /**
     * 快速返回成功
     *
     * @param <T>
     * @return
     */
    public static <T> ResponseVO ok() {
        return new ResponseVO<T>(ResponseStatusEnum.SUCCESS.getErrorCode(), ResponseStatusEnum.SUCCESS.getErrorMsg());
    }

    public static <T> ResponseVO ok(T result) {
        return new ResponseVO<>(ResponseStatusEnum.SUCCESS.getErrorCode(), ResponseStatusEnum.SUCCESS.getErrorMsg(), result);
    }

    public static <T> ResponseVO ok(String message, T result) {
        return new ResponseVO<>(ResponseStatusEnum.SUCCESS.getErrorCode(), message, result);
    }

    /**
     * 快速返回失败状态
     *
     * @param <T>
     * @return
     */
    public static <T> ResponseVO fail() {
        return new ResponseVO<T>(ResponseStatusEnum.REQUEST_ERROR.getErrorCode(), ResponseStatusEnum.REQUEST_ERROR.getErrorMsg());
    }

    public static <T> ResponseVO fail(ResponseStatusEnum responseStatusEnum) {
        return new ResponseVO<>(responseStatusEnum.getErrorCode(), responseStatusEnum.getErrorMsg());
    }

    /**
     * org.springframework.http.HttpStatus
     * 快速返回Http状态
     */
    public static <T> ResponseVO httpStatus(HttpStatus httpStatus, String message) {
        return new ResponseVO<T>(httpStatus.value(), message);
    }

    public static <T> ResponseVO httpStatus(HttpStatus httpStatus, String message, T result) {
        return new ResponseVO<>(httpStatus.value(), message, result);
    }
}