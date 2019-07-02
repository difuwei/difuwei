package com.jiedaibao.demo1.entity;

import lombok.Data;

@Data
public class User {
    private int id;
    private String name;
    private String password;
    private String number;

//    public int getId(){
//        return id;
//    }
//
//    public void setId(int id){
//        this.id = id;
//    }
//
//    public String getName(){
//        return name;
//    }
//    public void setName(String name){
//        this.name = name;
//    }
//    public String getPassword(){
//        return password;
//    }
//    public void setPassword(String password){
//        this.password = password;
//    }
//    public String getNumber(){
//        return number;
//    }
//    public void setNumber(String number){
//        this.number = number;
//    }

    public String tostring(){
        return "userId:"+id+"name:"+name+"password:"+password;
    }
}
