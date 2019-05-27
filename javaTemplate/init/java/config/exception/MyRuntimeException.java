package {{.PackageName}}.config.exception;

import {{.PackageName}}.common.ResponseModel;
import {{.PackageName}}.common.ResultCode;
import lombok.Getter;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Getter
public class MyRuntimeException extends RuntimeException {

    private String code;
    public MyRuntimeException(String code, String message){
        super(message);
        this.code = code;
    }
    public MyRuntimeException(ResultCode resultCode){
        super(resultCode.getErrorMsg());
        this.code = resultCode.getErrorCode();
    }
    public ResponseModel getResponseResult(){
        return ResponseModel.result(this.code, this.getMessage());
    }
}
