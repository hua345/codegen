
    @ApiOperation(value="{{.Description}}", notes="{{.Description}}")
    @{{.HttpMethod}}(path = "{{.MethodURL}}")
    public ResponseVO<{{.ResponseDTOName}}> {{.MethodName}}(@RequestBody(required=false) {{.RequestDTOName}} param) {
        log.info("Handing request {{.MethodName}} begin, req: {}", JSONObject.toJSONString(param));

        {{.ResponseDTOName}} {{.VarResponseDTOName}} = service.{{.MethodName}}(param);
        ResponseVO<{{.ResponseDTOName}}> result = ResponseVO.ok();
        log.info("Handing request {{.MethodName}} end, req: {}", JSONObject.toJSONString(result));
        return result;
    }