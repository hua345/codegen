package {{.PackageName}}.service.impl;

import com.alibaba.fastjson.JSONObject;
import hello.dto.request.AddUserInputDTO;
import hello.dto.response.AddUserOutputDTO;
import hello.mapper.UserMapper;
import hello.model.User;
import hello.service.UserService;
import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

/**
 * @author CHENJIANHUA001
 * @date 2019/03/18 15:54
 */
@Slf4j
@Service
public class UserServiceImpl implements UserService {
    @Resource
    private UserMapper userMapper;

    /**
     * 添加用户
     *
     * @param param 入参
     * @return 出参
     */
    @Override
    public AddUserOutputDTO addUser(AddUserInputDTO param) {
        log.info("call service addUser begin, req: {}", JSONObject.toJSONString(param));
        AddUserOutputDTO result = new AddUserOutputDTO();

        // TODO 补充业务逻辑代码
        User n = new User();
        n.setName(param.getName());
        n.setAge(param.getAge());
        userMapper.insertSelective(n);

        log.info("call service addUser end, resp: {}", JSONObject.toJSONString(result));
        return result;
    }

    /**
     * 查询所有用户
     *
     * @return
     */
    @Override
    public List<User> getAllUsers() {
        return userMapper.selectAllUser();
    }
}
