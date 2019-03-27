package {{.PackageName}}.common;

import com.alibaba.fastjson.JSONObject;

import java.io.Serializable;

import lombok.Getter;
import lombok.Setter;
import org.springframework.http.HttpStatus;
/**
 * @author CHENJIANHUA001
 * @date 2019/03/19 20:05
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
        return new ResponseModel<T>(ResultCode.SUCCESS.getErrorCode(), ResultCode.SUCCESS.getErrorMsg());
    }

    public static <T> ResponseModel success(T result) {
        return new ResponseModel<>(ResultCode.SUCCESS.getErrorCode(), ResultCode.SUCCESS.getErrorMsg(), result);
    }

    public static <T> ResponseModel success(String message, T result) {
        return new ResponseModel<>(ResultCode.SUCCESS.getErrorCode(), message, result);
    }

    /**
     * 快速返回失败状态
     *
     * @param <T>
     * @return
     */
    public static <T> ResponseModel fail() {
        return new ResponseModel<T>(ResultCode.REQUEST_ERROR.getErrorCode(), ResultCode.REQUEST_ERROR.getErrorMsg());
    }

    public static <T> ResponseModel fail(T result) {
        return new ResponseModel<>(ResultCode.REQUEST_ERROR.getErrorCode(), ResultCode.REQUEST_ERROR.getErrorMsg(), result);
    }

    public <T> ResponseModel fail(String message, T result) {
        return new ResponseModel<>(ResultCode.REQUEST_ERROR.getErrorCode(), message, result);
    }

    /**
     * 快速返回自定义状态码
     * @param resultCode
     * @return
     */
    public static ResponseModel result(ResultCode resultCode) {
        return result(resultCode.getErrorCode(), resultCode.getErrorMsg());
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
