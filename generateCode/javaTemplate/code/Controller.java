package {{.PackageName}}.controller;

import com.alibaba.fastjson.JSONObject;
import {{.PackageName}}.common.ResponseModel;
import {{.PackageName}}.common.ResultCode;
import {{.PackageName}}.config.exception.MyRuntimeException;
import hello.dto.request.AddUserInputDTO;
import hello.dto.response.AddUserOutputDTO;
import hello.service.UserService;
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
@RequestMapping(path = "/user")
public class UserController {

    /**
     * {{.Description}}
     */
    @Autowired
    private UserService userService;

    @ApiOperation(value="添加用户", notes="添加用户")
    @{{.HttpMethod}}(path = "/add")
    public ResponseModel<AddUserOutputDTO> addNewUser(@RequestBody AddUserInputDTO param) {
        log.info("Handing request addNewUser begin, req: {}", JSONObject.toJSONString(param));

        userService.addUser(param);
        AddUserOutputDTO addUserOutputDTO = new AddUserOutputDTO();
        addUserOutputDTO.setStatus("Saved");
        return ResponseModel.success(addUserOutputDTO);
    }
}
