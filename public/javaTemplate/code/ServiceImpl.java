package {{.ProjectInfo.PackageName}}.service.impl;

import com.alibaba.fastjson.JSONObject;
{{.ImportRequestDTOPath}}
{{.ImportResponseDTOPath}}
import {{.ProjectInfo.PackageName}}.service.{{.ControllerName}}Service;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Slf4j
@Service
public class {{.ControllerName}}ServiceImpl implements {{.ControllerName}}Service {

    /**
     * {{.Description}}
     *
     * @param param 入参
     * @return 出参
     */
    @Override
    public {{.ResponseDTOName}} {{.MethodName}}({{.RequestDTOName}} param) {
        log.info("call service {{.MethodName}} begin, req: {}", JSONObject.toJSONString(param));
        {{.ResponseDTOName}} result = new {{.ResponseDTOName}}();

        log.info("call service {{.MethodName}} end, resp: {}", JSONObject.toJSONString(result));
        return result;
    }
}
