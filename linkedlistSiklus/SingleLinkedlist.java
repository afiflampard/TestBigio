/*
 * To change this license header, choose License Headers in Project Properties.
 * To change this template file, choose Tools | Templates
 * and open the template in the editor.
 */
package linkedlist;

class IntList {

    public int value;
    public IntList next;

    public IntList(int value) {
        this.value = value;
    }
}

public class SingleLinkedlist {

    public IntList head = null, tail = null;
    int counter = 0;

    public void sisipDepan(int value) {
        IntList newData = new IntList(value);
        if (head == null) {
            System.out.println("masuk sini");
            head = newData;
            tail = newData;
            counter++;
        } else {
            newData.next = head;
            head = newData;
            counter++;
        }
    }

    public void sisipDiakhir(int value) {
        IntList newData = new IntList(value);
        if (tail == null) {
            head = tail = newData;
            counter++;
        } else {
            tail.next = newData;
            tail = newData;
            counter++;
        }
    }

    public void sisipAfter(int value, int key) {
        IntList newNode = new IntList(value);
        IntList awal = head;
        while (awal != null) {
            if (awal.value == key) {
                newNode.next = awal.next;
                awal.next = newNode;
                counter++;
                break;
            }
            awal = awal.next;
        }
    }

    public void sisipBefore(int value, int key) {
        IntList awal = head;
        if (awal.value == key) {
            sisipDepan(value);
        } else {
            IntList newNode = new IntList(value);
            while (awal != null) {
                if (awal.next.value == key) {
                    newNode.next = awal.next;
                    awal.next = newNode;
                    counter++;
                    break;
                }
            }
        }

    }
     public boolean isEmpty(){
        return ((head=tail)==null);
    }
      public void cetak(){
        IntList awal = head;
        while(awal != null){
            System.out.print(awal.value+"->");
            //System.out.println("Jumlahnya :"+counter);
            awal = awal.next;
        }
        System.out.println("");
        System.out.println("Jumlahnya : "+counter);
    }
      public int CheckSiklus(){
          if(counter > 1){
              return 1;
          }
          return 0;
      }
}
