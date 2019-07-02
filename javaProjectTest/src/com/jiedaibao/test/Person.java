package com.jiedaibao.test;

import java.util.concurrent.SynchronousQueue;

public class Person {
    public int aaa;
    public Person(){
        System.out.println("父类的构造函数");
    }
    public int get(){
        return aaa;
    }
}
