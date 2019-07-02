package com.jiedaibao.demo1.mapper;

import com.jiedaibao.demo1.entity.User;
import org.apache.ibatis.annotations.Mapper;
import org.apache.ibatis.annotations.Param;

import java.util.List;

@Mapper
public interface UserMapper {

    List<User> findUserByName(@Param("name") String name);

    public List<User> ListUser();

    public int insertUser(User user);

    public int delete(@Param("id") int id);

    public int Update(User user);
}
