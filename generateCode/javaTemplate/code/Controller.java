package {{.ProjectInfo.PackageName}}.controller;

import com.alibaba.fastjson.JSONObject;
import {{.ProjectInfo.PackageName}}.common.ResponseModel;
import {{.ProjectInfo.PackageName}}.common.ResultCode;
import {{.ProjectInfo.PackageName}}.config.exception.MyRuntimeException;
import {{.ProjectInfo.PackageName}}.dto.request.{{.DTOName}}InputDTO;
import {{.ProjectInfo.PackageName}}.dto.response.{{.DTOName}}OutputDTO;
import {{.ProjectInfo.PackageName}}.service.{{.ControllerName}}Service;
import io.swagger.annotations.ApiOperation;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.PostMapping;
import org.springframework.web.bind.annotation.GetMapping;
import org.springframework.web.bind.annotation.RequestBody;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RestController;
import org.springframework.web.bind.annotation.RequestParam;


/**
 * @author CHENJIANHUA001
 * @date 2019/03/18 15:28
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
    public ResponseModel<{{.DTOName}}OutputDTO> {{.MethodName}}(@RequestBody {{.DTOName}}InputDTO param) {
        log.info("Handing request {{.MethodName}} begin, req: {}", JSONObject.toJSONString(param));

        {{.DTOName}}OutputDTO {{.MethodName}}OutputDTO = service.{{.MethodName}}(param);
        return ResponseModel.success({{.MethodName}}OutputDTO);
    }
}
