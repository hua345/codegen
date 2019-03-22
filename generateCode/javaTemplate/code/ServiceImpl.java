package {{.ProjectInfo.PackageName}}.service.impl;

import com.alibaba.fastjson.JSONObject;
import {{.ProjectInfo.PackageName}}.dto.request.{{.DTOName}}InputDTO;
import {{.ProjectInfo.PackageName}}.dto.response.{{.DTOName}}OutputDTO;
import {{.ProjectInfo.PackageName}}.service.{{.ControllerName}}Service;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

/**
 * @author CHENJIANHUA001
 * @date 2019/03/18 15:54
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
    public {{.DTOName}}OutputDTO {{.MethodName}}({{.DTOName}}InputDTO param) {
        log.info("call service {{.MethodName}} begin, req: {}", JSONObject.toJSONString(param));
        {{.DTOName}}OutputDTO result = new {{.DTOName}}OutputDTO();

        log.info("call service {{.MethodName}} end, resp: {}", JSONObject.toJSONString(result));
        return result;
    }
}
