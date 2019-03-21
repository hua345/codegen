package {{.PackageName}}.service;

import hello.dto.request.AddUserInputDTO;
import hello.dto.response.AddUserOutputDTO;
import hello.model.User;

import java.util.List;

/**
 * @author CHENJIANHUA001
 * @date 2019/03/18 15:54
 */
public interface UserService {

    /**
     * 添加用户
     * @param param 入参
     * @return 出参
     */
    AddUserOutputDTO addUser(AddUserInputDTO param);

    /**
     * 查询所有用户
     * @return
     */
    List<User> getAllUsers();
}
