package com.jiedaibao.demo2.mapper;

import com.jiedaibao.demo2.entity.User;
import java.util.List;

public interface UserMapper {

    List<User> findUserByName( String name);

    public List<User> ListUser();

    public int insertUser(User user);

    public int delete(int id);

    public int Update(User user);
}
