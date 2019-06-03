package {{.PackageName}}.config.exception;

import com.alibaba.fastjson.JSONObject;

import {{.PackageName}}.common.ResponseVO;
import {{.PackageName}}.common.ResponseStatusEnum;
import {{.PackageName}}.utils.ResponseUtil;
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
	public ResponseVO defaultErrorHandler(HttpServletRequest req, Exception e) {
		log.error("Exception", e);
		return  ResponseUtil.fail(ResponseStatusEnum.SERVER_ERROR);
	}

	/**
	 * GET/POST请求方法错误的拦截器
	 * 因为开发时可能比较常见,而且发生在进入controller之前,上面的拦截器拦截不到这个错误
	 * 所以定义了这个拦截器
	 */
	@ExceptionHandler(HttpRequestMethodNotSupportedException.class)
	public ResponseVO httpRequestMethodHandler() {
		log.error("Catch HttpRequestMethodNotSupportedException");
		return ResponseUtil.fail(ResponseStatusEnum.REQUEST_METHOD_ERROR);
	}

	@ExceptionHandler(MissingServletRequestParameterException.class)
	public ResponseVO requestParameterExceptionHandler(MissingServletRequestParameterException missingServletRequestParameterException) {
		log.error("Catch MissingServletRequestParameterException {}.", missingServletRequestParameterException.getMessage());
		return new ResponseVO(ResponseStatusEnum.PARAMETER_CHECK_ERROR.getErrorCode(), missingServletRequestParameterException.getMessage());
	}
	/**
	 * 本系统自定义错误的拦截器
	 * 拦截到此错误之后,就返回这个类里面的json给前端
	 * 常见使用场景是参数校验失败,抛出此错,返回错误信息给前端
	 */
	@ExceptionHandler(MyRuntimeException.class)
	public ResponseVO myRuntimeExceptionHandler(MyRuntimeException myRuntimeException) {
		log.error("Catch {} MyRuntimeException ; {}", myRuntimeException.getStackTrace()[0].toString(), myRuntimeException.getResponseResult());
		return myRuntimeException.getResponseResult();
	}
}