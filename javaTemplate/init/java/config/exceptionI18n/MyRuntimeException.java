package {{.PackageName}}.config.exception;

import {{.PackageName}}.common.ResponseVO;
import {{.PackageName}}.common.ResponseStatusEnum;
import lombok.Getter;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Getter
public class MyRuntimeException extends RuntimeException {

    private Integer code;

    public MyRuntimeException(Integer code, String message) {
        super(message);
        this.code = code;
    }

    public MyRuntimeException(ResponseStatusEnum responseStatusEnum) {
        super(responseStatusEnum.getI18nKey());
        this.code = responseStatusEnum.getErrorCode();
    }

    public ResponseVO getResponseResult() {
        return new ResponseVO(this.code, this.getMessage());
    }
}