package com.jiedaibao.test;

import java.time.Instant;
import java.util.*;

/**
 * bankAccountName : test string
 */
public class Test3 {

    public static int add(int a,int b){
        return a+b;
    }
    public int add(int... args){
        int sum=0;
        for (int v:args
             ) {
            sum+=v;
        }
        return sum;
    }
    public static void main(String args[]){
        System.out.println(add(1,2));
        Test3 obj1 = new Test3();
        Class labclass = obj1.getClass();

        System.out.println(labclass);
//        Date.from(Instant.now());

        ArrayList list1 = new ArrayList();
        ArrayList list2 = new ArrayList();
        list1.add(3);
        list1.add(55);
//        System.out.println(list1.indexOf(55));
        Iterator it1 = list1.iterator();
//        while (it1.hasNext()){
//            System.out.println(it1.next());
//        }

        LinkedList<String> linklist1 = new LinkedList<>();
        linklist1.add("a");
        linklist1.add("c");
        linklist1.add("b");
        linklist1.add("aaaa");
        Iterator it2 = linklist1.iterator();
//        while (it2.hasNext()){
//            System.out.println(it2.next());
//        }
        HashMap<Integer,String> hashmap = new HashMap<>();
        hashmap.put(1,"aaaa");
        hashmap.put(2,"bb");
        hashmap.put(3,"eee");
        hashmap.put(4,"rr");
        System.out.println(hashmap.keySet());
//        for (int k:hashmap.keySet()
//             ) {
//            System.out.println(hashmap.get(k));
//        }
    }
}
