package {{.ProjectInfo.PackageName}}.controller;

import com.alibaba.fastjson.JSONObject;
import {{.ProjectInfo.PackageName}}.common.ResponseVO;
import {{.ProjectInfo.PackageName}}.common.ResponseStatusEnum;
import {{.ProjectInfo.PackageName}}.config.exception.MyRuntimeException;
{{.ImportRequestDTOPath}}
{{.ImportResponseDTOPath}}
import {{.ProjectInfo.PackageName}}.service.{{.ControllerName}}Service;
import {{.ProjectInfo.PackageName}}.utils.ResponseUtil;
import io.swagger.annotations.ApiOperation;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
@Slf4j
@RestController
@RequestMapping(path = "{{.ControllerURL}}")
public class {{.ControllerName}}Controller {

    /**
     * {{.Description}}
     */
    @Autowired
    private {{.ControllerName}}Service service;

    @ApiOperation(value="{{.Description}}", notes="{{.Description}}")
    @{{.HttpMethod}}(path = "{{.MethodURL}}")
    public ResponseVO<{{.ResponseDTOName}}> {{.MethodName}}(@RequestBody(required=false) {{.RequestDTOName}} param) {
        log.info("Handing request {{.MethodName}} begin, req: {}", JSONObject.toJSONString(param));

        {{.ResponseDTOName}} {{.VarResponseDTOName}} = service.{{.MethodName}}(param);
        ResponseVO<{{.ResponseDTOName}}> result = ResponseUtil.ok();
        log.info("Handing request {{.MethodName}} end, req: {}", JSONObject.toJSONString(result));
        return result;
    }
}
