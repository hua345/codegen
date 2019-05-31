package {{.PackageName}}.common;

import lombok.Getter;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Getter
public enum ResponseStatusEnum {
    /**
     * 请求成功
     */
    SUCCESS(200, "Success"),
    REQUEST_ERROR(401, "请求失败"),
    REQUEST_METHOD_ERROR(402, "GET/POST请求方法错误"),
    SERVER_ERROR(500, "服务器异常"),
    REQUEST_PATH_ERROR(501, "请求路径不存在"),
    PARAMETER_CHECK_ERROR(10001, "参数错误"),
    LOGIN_EXPIRE(20001, "登陆已过期,请重新登陆"),
    AUTH_VALID_ERROR(20002, "用户权限不足");

    private Integer errorCode;

    private String errorMsg;

    ResponseStatusEnum(Integer errorCode, String errorMsg) {
        this.errorCode = errorCode;
        this.errorMsg = errorMsg;
    }
}