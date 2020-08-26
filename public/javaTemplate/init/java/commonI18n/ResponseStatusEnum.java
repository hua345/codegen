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
    SUCCESS(200, "response.success"),
    REQUEST_ERROR(401, "response.REQUEST_ERROR"),
    REQUEST_METHOD_ERROR(402, "response.REQUEST_METHOD_ERROR"),
    SERVER_ERROR(500, "response.SERVER_ERROR"),
    REQUEST_PATH_ERROR(501, "response.REQUEST_PATH_ERROR"),
    REQUEST_MEDIA_TYPE_ERROR(502, "response.REQUEST_MEDIA_TYPE_ERROR"),
    PARAMETER_CHECK_ERROR(10001, "response.PARAMETER_CHECK_ERROR"),
    LOGIN_EXPIRE(20001, "response.LOGIN_EXPIRE"),
    AUTH_VALID_ERROR(20002, "response.AUTH_VALID_ERROR"),
    AUTH_UNKNOWN_ACCOUNT(20003, "response.AUTH_UNKNOWN_ACCOUNT"),
    AUTH_INCORRECT_PASSWORD(20004, "response.AUTH_INCORRECT_PASSWORD"),
    AUTH_TOKEN_MISSED(2005, "response.AUTH_TOKEN_MISSED"),
    DATE_TIME_PARSE_EXCEPTION(2006, "response.DATE_TIME_PARSE_EXCEPTION");

    private Integer errorCode;

    private String i18nKey;

    ResponseStatusEnum(Integer errorCode, String i18nKey) {
        this.errorCode = errorCode;
        this.i18nKey = i18nKey;
    }
}
