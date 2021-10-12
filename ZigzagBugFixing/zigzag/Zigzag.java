/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package zigzag;

import java.util.Arrays;
import java.util.Scanner;

/**
 *
 * @author Afif Musyayyidin
 */
public class Zigzag {

    /**
     * @param args the command line arguments
     */
    public static void main(String[] args) {
        Scanner kb = new Scanner(System.in);
        int test_cases = kb.nextInt();
        for (int cs = 1; cs <= test_cases; cs++) {
            int n = kb.nextInt();
            int a[] = new int[n];
            for (int i = 0; i < n; i++) {
                a[i] = kb.nextInt();
            }
            findZigZagSequence(a, n);
        }
    }

    public static void findZigZagSequence(int[] a, int n) {
        Arrays.sort(a);
        int mid = ((n) / 2);
        int temp = a[mid];
        a[mid] = a[n - 1];
        a[n - 1] = temp;

        int st = mid + 1;
        
        int ed = n - 1;
        while (st <= ed) {
            if (ed != n) {
                temp = a[st];
                a[st] = a[ed];
                a[ed] = temp;
                ed = ed + 1;
            }
            st = st + 1;
            
        }
        for (int i = 0; i < n; i++) {
            if (i > 0) {
                System.out.print(" ");
            }
            System.out.print(a[i]);
        }
        System.out.println();
    }

}
