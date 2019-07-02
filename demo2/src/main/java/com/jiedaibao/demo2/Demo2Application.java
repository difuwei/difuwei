package com.jiedaibao.demo2;

//import org.mybatis.spring.annotation.MapperScan;
import org.springframework.boot.autoconfigure.SpringBootApplication;

import java.text.DecimalFormat;
import java.util.HashMap;

@SpringBootApplication
//@MapperScan("com.jiedaibao.demo2.mapper")
public class Demo2Application {

    public static void main(String[] args) {
//        new Demo2Application().formatMoneyYuan("10.2");
        HashMap<String,String> testMap = new HashMap<>();
        HashMap<String,String> testMap2 = new HashMap<>();
        testMap2.put("aaa","1111");
        testMap.put("aaa","1111");
        testMap2.put("bbb","22");
        testMap.put("bbb","22");

        System.out.println(testMap.hashCode());
    }

    public void formatMoneyYuan(String amount){
        String formatMoney = "";
        Double dAmount = Double.parseDouble(amount);
        DecimalFormat df = new DecimalFormat("0.00");
        formatMoney = df.format(dAmount);
        System.out.println(formatMoney);
    }

    public enum Color
    {
        RED,BLUE,GREEN,BLACK;
    }


}
