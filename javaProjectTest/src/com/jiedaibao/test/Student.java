package com.jiedaibao.test;

import java.sql.Struct;

public class Student extends Person{
    public static String namePrefix = "hello";
    public String Name;
    private int Age;
    private boolean Sex;

    private int aaa=66666;
    public boolean isSex() {
        return Sex;
    }

    public void setSex(boolean sex) {
        this.Sex = sex;
    }

    public Student(String Name, int Age, boolean sex) {
        this.Name = Name;
        this.Age = Age;
        this.setSex(sex);
    }

    public Student() {
        System.out.println("子类的构造方法");

    }

    public int getAge() {
        return Age;
    }

    public void setName(String name) {
        this.Name = name;
    }

    public String getName() {
        return Name;
    }
    public static void print(){
        System.out.println(122222);
    }

    public static void main(String args[]) {
        Student student1 = new Student();
        student1.Name = "zhangsan";
        student1.Age = 18;
        student1.Sex = false;


//        String sex = student1.isSex() ? "mail" : "femail";
//        System.out.println("name:" + student1.Name + ",age:" + student1.Age + ",sex:" + sex);

//        Student student2 = new Student("lisi", 20, true);
//        student2.setName("lisi2");
//        System.out.println("name:" + student2.Name + ",age:" + student2.Age + ",sex:" + student2.Sex);
//        System.out.println(namePrefix + "_" + student2.getName());
    }

}
