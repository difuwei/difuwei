package com.jiedaibao.demo2.service;


import com.jiedaibao.demo2.entity.User;
import com.jiedaibao.demo2.mapper.UserMapper;
import org.springframework.stereotype.Service;

import javax.annotation.Resource;
import java.util.List;

@Service
public class UserService {

    @Resource
    private UserMapper userMapper;

    public List<User> findByName(String name) {
        return userMapper.findUserByName(name);
    }

    public User insertUser(User user) {
        userMapper.insertUser(user);
        return user;
    }
    public List<User> ListUser(){
        return	userMapper.ListUser();
    }


    public int Update(User user){
        return userMapper.Update(user);
    }

    public int delete(int id){
        return userMapper.delete(id);
    }
}
