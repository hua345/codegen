package {{.ProjectInfo.PackageName}}.service;

import {{.ProjectInfo.PackageName}}.dto.request.{{.DTOName}}InputDTO;
import {{.ProjectInfo.PackageName}}.dto.response.{{.DTOName}}OutputDTO;

import java.util.List;

/**
 * @author CHENJIANHUA001
 * @date 2019/03/18 15:54
 */
public interface {{.ControllerName}}Service {

    /**
     * {{.Description}}
     * @param param 入参
     * @return 出参
     */
    {{.DTOName}}OutputDTO {{.MethodName}}({{.DTOName}}InputDTO param);
}
