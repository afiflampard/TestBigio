/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package javaapplication6;

import java.util.Scanner;

/**
 *
 * @author Afif Musyayyidin
 */
public class JavaApplication6 {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        System.out.println("Masukkan kata : ");
        String kata = in.nextLine();
        System.out.println("Masukkan angka geser");
        int geser = in.nextInt();
        char []ktempKata = new char[kata.length()];
        for (int i = 0; i < kata.length(); i++) {
            ktempKata[i] =(char) ((char) kata.charAt(i)+geser);
        }
        for (int i = 0; i < ktempKata.length; i++) {
            System.out.print(ktempKata[i]);
        }
    }
    
}
