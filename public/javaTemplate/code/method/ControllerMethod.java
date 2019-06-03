    {{ if .SupportSwagger }}
    @ApiOperation(value="{{.Description}}", notes="{{.Description}}"){{ end }}
    @{{.HttpMethod}}(path = "{{.MethodURL}}")
    public ResponseVO<{{.ResponseDTOName}}> {{.MethodName}}(@RequestBody(required=false) {{.RequestDTOName}} param) {
        log.info("Handing request {{.MethodName}} begin, req: {}", JSONObject.toJSONString(param));

        {{.ResponseDTOName}} {{.VarResponseDTOName}} = service.{{.MethodName}}(param);
        ResponseVO<{{.ResponseDTOName}}> result = ResponseUtil.ok();
        log.info("Handing request {{.MethodName}} end, req: {}", JSONObject.toJSONString(result));
        return result;
    }