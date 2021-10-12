/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package javaapplication5;

import java.util.Scanner;

/**
 *
 * @author Afif Musyayyidin
 */
public class JavaApplication5 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        System.out.println("Masukkan kata : ");
        String data = in.nextLine();
        String[] tempData = data.split("");
        for (int i = tempData.length-1; i >= 0; i--){
            System.out.print(tempData[i]);
        }
    }
    
}
