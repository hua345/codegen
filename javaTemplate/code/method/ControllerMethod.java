
    @ApiOperation(value="{{.Description}}", notes="{{.Description}}")
    @{{.HttpMethod}}(path = "{{.MethodURL}}")
    public ResponseModel<{{.ResponseDTOName}}> {{.MethodName}}(@RequestBody {{.RequestDTOName}} param) {
        log.info("Handing request {{.MethodName}} begin, req: {}", JSONObject.toJSONString(param));

        {{.ResponseDTOName}} {{.VarResponseDTOName}} = service.{{.MethodName}}(param);
        ResponseModel<{{.ResponseDTOName}}> result = ResponseModel.success();
        log.info("Handing request {{.MethodName}} end, req: {}", JSONObject.toJSONString(result));
        return result;
    }