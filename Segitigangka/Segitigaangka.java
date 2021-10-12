/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package segitigaangka;

import java.util.Scanner;

/**
 *
 * @author Afif Musyayyidin
 */
public class Segitigaangka {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);
        System.out.println("Masuukan tinggi segitiga : ");
        int tinggi = in.nextInt();
        int x = tinggi * 2;
        for (int i = 1; i <= tinggi; i++) {
            for (int j = 1; j <= i; j++) {
                System.out.print(j);
            }
            for (int j = x - 1; j >= i; j--) {
                System.out.print(" ");
            }

            for (int d = i; d > 0; d--) {
                System.out.print(d);
            }
            x -= 1;
            System.out.println();

        }
        int g = 1;
        for (int i = tinggi; i >= 1; i--) {
            for (int j = 1; j <= i; j++) {
                System.out.print(j);
            }
             for (int s = i; s < tinggi+1; s++) {
                System.out.print(" ");
            }
            for (int d = (tinggi+1) - g; d > 0; d--) {
                System.out.print(d);
            }
            g += 1;
            System.out.println("");

        }



    }

}
