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
public class ResponseModel<T> implements Serializable {
    private static final long serialVersionUID = 1L;
    /**
     * 返回状态码
     */
    private String code;

    /**
     * 返回消息
     */
    private String msg;

    /**
     * 返回内容
     */
    private T data;

    public ResponseModel(String code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    public ResponseModel(String code, String msg, T data) {
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
    public static <T> ResponseModel success() {
        return new ResponseModel<T>(ResultCodeEnum.SUCCESS.getErrorCode(), ResultCodeEnum.SUCCESS.getErrorMsg());
    }

    public static <T> ResponseModel success(T result) {
        return new ResponseModel<>(ResultCodeEnum.SUCCESS.getErrorCode(), ResultCodeEnum.SUCCESS.getErrorMsg(), result);
    }

    public static <T> ResponseModel success(String message, T result) {
        return new ResponseModel<>(ResultCodeEnum.SUCCESS.getErrorCode(), message, result);
    }

    /**
     * 快速返回失败状态
     *
     * @param <T>
     * @return
     */
    public static <T> ResponseModel fail() {
        return new ResponseModel<T>(ResultCodeEnum.REQUEST_ERROR.getErrorCode(), ResultCodeEnum.REQUEST_ERROR.getErrorMsg());
    }

    public static <T> ResponseModel fail(T result) {
        return new ResponseModel<>(ResultCodeEnum.REQUEST_ERROR.getErrorCode(), ResultCodeEnum.REQUEST_ERROR.getErrorMsg(), result);
    }

    public <T> ResponseModel fail(String message, T result) {
        return new ResponseModel<>(ResultCodeEnum.REQUEST_ERROR.getErrorCode(), message, result);
    }

    /**
     * 快速返回自定义状态码
     * @param resultCodeEnum
     * @return
     */
    public static ResponseModel result(ResultCodeEnum resultCodeEnum) {
        return result(resultCodeEnum.getErrorCode(), resultCodeEnum.getErrorMsg());
    }
    public static <T> ResponseModel result(String statusCode, String message) {
        return new ResponseModel<T>(statusCode, message);
    }

    public static <T> ResponseModel result(String statusCode, String message, T result) {
        return new ResponseModel<>(statusCode, message, result);
    }
    /**
     * org.springframework.http.HttpStatus
     * 快速返回Http状态
     */
    public static <T>ResponseModel httpStatus(HttpStatus httpStatus, String message){
        return result(httpStatus.toString(),message);
    }

    public static <T>ResponseModel httpStatus(HttpStatus httpStatus, String message, T result){
        return result(httpStatus.toString(),message,result);
    }
}
