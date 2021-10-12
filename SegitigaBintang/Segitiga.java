/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package segitiga;

import java.util.Scanner;

/**
 *
 * @author Afif Musyayyidin
 */
public class Segitiga {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner in = new Scanner(System.in);

        System.out.println("Masukkan tinggi segitiga");
        int masukan = in.nextInt();
        for (int i = masukan; i >= 1; i--) {

            for (int j = 1; j <= masukan - i; j++) {
                System.out.print(" ");//karakter spasi
            }
            for (int j = 1; j < 2 * i; j++) {
                System.out.print("*");
            }

            System.out.println("");

        }

        for (int i = 1; i <= masukan; i++) {
            for (int j = masukan-1; j >= i; j--) {
                System.out.print(' ');
            }
            for (int k = 1; k <= i; k++) {
                System.out.print('*');
            }
            for (int l = 1; l <= i - 1; l++) {
                System.out.print('*');
            }
            System.out.println();
        }
    }
}
