package {{.PackageName}}.config.exception;

import {{.PackageName}}.common.ResponseModel;
import {{.PackageName}}.common.ResultCodeEnum;
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
    public MyRuntimeException(ResultCodeEnum resultCodeEnum){
        super(resultCodeEnum.getErrorMsg());
        this.code = resultCodeEnum.getErrorCode();
    }
    public ResponseModel getResponseResult(){
        return ResponseModel.result(this.code, this.getMessage());
    }
}