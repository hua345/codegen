package {{.PackageName}}.config.exception;

import com.alibaba.fastjson.JSONObject;

import {{.PackageName}}.common.ResponseModel;
import {{.PackageName}}.common.ResultCode;
import lombok.extern.slf4j.Slf4j;
import org.springframework.web.HttpRequestMethodNotSupportedException;
import org.springframework.web.bind.MissingServletRequestParameterException;
import org.springframework.web.bind.annotation.ControllerAdvice;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.ResponseBody;

import javax.servlet.http.HttpServletRequest;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Slf4j
@ControllerAdvice
@ResponseBody
public class GlobalExceptionHandler {
	@ExceptionHandler(value = Exception.class)
	public ResponseModel defaultErrorHandler(HttpServletRequest req, Exception e) {
		String errorPosition = "";
		//如果错误堆栈信息存在
		if (e.getStackTrace().length > 0) {
			StackTraceElement element = e.getStackTrace()[0];
			String fileName = element.getFileName() == null ? "未找到错误文件" : element.getFileName();
			int lineNumber = element.getLineNumber();
			errorPosition = fileName + ":" + lineNumber;
		}
		ResponseModel responseModel = ResponseModel.result(ResultCode.SERVER_ERROR);
		JSONObject errorObject = new JSONObject();
		errorObject.put("errorLocation", e.toString() + "    错误位置:" + errorPosition);
		responseModel.setData(errorObject);
		log.error("Exception", e);
		return responseModel;
	}

	/**
	 * GET/POST请求方法错误的拦截器
	 * 因为开发时可能比较常见,而且发生在进入controller之前,上面的拦截器拦截不到这个错误
	 * 所以定义了这个拦截器
	 */
	@ExceptionHandler(HttpRequestMethodNotSupportedException.class)
	public ResponseModel httpRequestMethodHandler() {
		log.error("Catch HttpRequestMethodNotSupportedException");
		return ResponseModel.result(ResultCode.REQUEST_METHOD_ERROR);
	}

	@ExceptionHandler(MissingServletRequestParameterException.class)
    public ResponseModel requestParameterExceptionHandler(MissingServletRequestParameterException missingServletRequestParameterException) {
        log.error("Catch MissingServletRequestParameterException {}.", missingServletRequestParameterException.getMessage());
        return ResponseModel.result(ResultCode.PARAMETER_CHECK_ERROR.getErrorCode(), missingServletRequestParameterException.getMessage());
    }
	/**
	 * 本系统自定义错误的拦截器
	 * 拦截到此错误之后,就返回这个类里面的json给前端
	 * 常见使用场景是参数校验失败,抛出此错,返回错误信息给前端
	 */
	@ExceptionHandler(MyRuntimeException.class)
	public ResponseModel myRuntimeExceptionHandler(MyRuntimeException myRuntimeException) {
		log.error("Catch MyRuntimeException ; {}", myRuntimeException.getResponseResult());
		return myRuntimeException.getResponseResult();
	}
}
