/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package javaapplication3;

import java.util.Scanner;

/**
 *
 * @author Afif Musyayyidin
 */
public class JavaApplication3 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        System.out.println("Masukkan berapa huruf");
        int masukan = in.nextInt();
        char data[] = new char[masukan];
        for (int i = 0; i < masukan; i++) {
            data[i] = in.next().charAt(0);
        }
        for (int i = data.length-1; i >= 0; i--) {
            System.out.print(data[i]);
        }
    }

}
