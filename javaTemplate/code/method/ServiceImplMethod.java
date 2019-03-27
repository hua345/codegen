
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