package {{.PackageName}}.common;

import lombok.Getter;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Getter
public enum ResultCodeEnum {
    /**
     * 请求成功
     */
    SUCCESS(200, "response.success"),
    REQUEST_ERROR(401, "response.REQUEST_ERROR"),
    REQUEST_METHOD_ERROR(402, "response.REQUEST_METHOD_ERROR"),
    SERVER_ERROR(500, "response.SERVER_ERROR"),
    REQUEST_PATH_ERROR(501, "response.REQUEST_PATH_ERROR"),
    PARAMETER_CHECK_ERROR(10001, "response.PARAMETER_CHECK_ERROR"),
    LOGIN_EXPIRE(20001, "response.LOGIN_EXPIRE"),
    AUTH_VALID_ERROR(20002, "response.AUTH_VALID_ERROR");

    private Integer errorCode;

    private String errorMsg;

    ResultCodeEnum(Integer errorCode, String errorMsg) {
        this.errorCode = errorCode;
        this.errorMsg = errorMsg;
    }
}
