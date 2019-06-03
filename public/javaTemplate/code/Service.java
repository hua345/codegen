package {{.ProjectInfo.PackageName}}.service;

{{.ImportRequestDTOPath}}
{{.ImportResponseDTOPath}}

import java.util.List;

/**
 * @author {{.Author}}
 * @date {{.NowDate}}
 */
public interface {{.ControllerName}}Service {

    /**
     * {{.Description}}
     * @param param 入参
     * @return 出参
     */
    {{.ResponseDTOName}} {{.MethodName}}({{.RequestDTOName}} param);
}
