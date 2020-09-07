package {{.PackageName}}.common;

import java.io.Serializable;

import {{.PackageName}}.utils.JsonUtil;
import lombok.Getter;
import lombok.Setter;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Getter
@Setter
public class ResponseVO<T> implements Serializable {
    private static final long serialVersionUID = 1L;
    /**
     * 返回状态码
     */
    private Integer code;

    /**
     * 返回消息
     */
    private String msg;

    /**
     * 返回内容
     */
    private T data;

    public ResponseVO(Integer code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    public ResponseVO(Integer code, String msg, T data) {
        this.code = code;
        this.msg = msg;
        this.data = data;
    }

    @Override
    public String toString() {
        return JsonUtil.toJSONString(this);
    }
}