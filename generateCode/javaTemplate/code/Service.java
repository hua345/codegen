package {{.ProjectInfo.PackageName}}.service;

{{.ImportRequestDTOPath}}
{{.ImportResponseDTOPath}}

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
    {{.ResponseDTOName}} {{.MethodName}}({{.RequestDTOName}} param);
}
