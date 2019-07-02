package com.jiedaibao.test;

import java.lang.reflect.Array;
import java.util.Calendar;
import java.util.HashMap;
import java.util.Scanner;

public class Test1 {
    public static void test(int a,int b)
    {
//        int c = a+b;
//        return c;
//            System.out.println(a+b);

//        if (a < 5){
//            System.out.println("yes");
//        }else if (a > 5){
//            System.out.println("yes");
//
//        }else if (a == 5){
//            System.out.println("no");
//
//        }
//        System.out.println(a+b);
        Scanner sc = new Scanner(System.in);
        System.out.println(sc.nextInt());

    }
    public static void main(String[] args)
    {
//        Calendar calendar = Calendar.getInstance();
//        System.out.println(calendar.DAY_OF_WEEK);
//        int c = add(1,4);
//        System.out.println(c);
        String str1 = "hello我";
        String str2 = new String("hello");
        System.out.println(str2.getBytes());
        int[] arr1={3,6,7};
        System.out.println(str1.split(","));
        String[] arr2={"a","b","c","d","e"};
        char[] arr3={'a','b','c'};
//        char[] arr4={d};
        boolean equals = arr1.equals(str2);
        System.out.println();
        for (int $i=0;$i<arr1.length;$i++){
            System.out.println($i+"--------"+arr1[$i]);
        }

//        HashMap
        System.out.println(str1);

    }


    public static int add(int a,int b){
        return a+b;
    }

}

