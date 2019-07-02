package com.jiedaibao.test;

import org.omg.CORBA.Object;

import java.util.HashMap;

public class Test2 {
    public static void main(String[] args)
    {
//        System.out.println("123");
//        StringBuffer str1 = new StringBuffer("h");
//        str1.setCharAt(5,'A');
////        System.out.println(str1);
//        int arr1[];
        int[] arr2 = {4,5,2,6,8,18};
//        System.out.println(getMaxNum(arr2));
//        System.out.println(getMaxNum(arr2).get("maxValue"));
//        printArr(arr2);
        int[] res = maoPaoSort(arr2);
        for (int v:res
             ) {
            System.out.println(v);
        }
        Student student1 = new Student("zhangsan",43,true);
    }
    public static HashMap<String, Integer> getMaxNum(int arr2[]){
        HashMap<String,Integer> res = new HashMap<>();
        int maxValue = arr2[0];
        for(int i=1;i<arr2.length;i++){
            if (arr2[i]>maxValue){
                res.put("maxIndex",i);
                res.put("maxValue",arr2[i]);
            }
        }
        return res;
    }

    public static void printArr(int arr[]){
//        HashMap<Integer,Integer> res = new HashMap<>();
        int res2[] = new int[5];
        for (int i=0;i<arr.length;i++){
            if(i%2 == 0){
//                res.put(i,arr[i]);
                res2[i] = arr[i];
            }
        }
        for (int v: res2
             ) {
            System.out.println(v);
        }
    }

    public static int[] maoPaoSort(int arr[]){
        for (int i=0;i<arr.length-1;i++){
            for (int j=0;j<arr.length-1-i;j++){
                if(arr[j]>arr[j+1]){
                    int temp = arr[j];
                    arr[j] = arr[j+1];
                    arr[j+1] = temp;
                }
            }
        }
        return arr;
    }

}
