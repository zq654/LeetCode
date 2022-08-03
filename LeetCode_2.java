public class LeetCode_2 {
   public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
      int a1=CountNumber(l1);
      int a2=CountNumber(l2);
      int count=a1+a2;
      int act=count%10;
      ListNode head=new ListNode(act);
      count=count/10;
      ListNode pointer=head;
      while(count!=0){
         int value=count%10;
         count=count/10;
         ListNode Next=new ListNode(value);
         pointer.next=Next;
         pointer=pointer.next;
      }
      pointer.next=null;
      return head;
   }

   static int CountNumber(ListNode l) {
      int num=0;
      int res=0;
      while(l!=null){
         res+=(int)(Math.pow(10,num)*l.val);
         l=l.next;
      }
      return res;
   }

   public static void main(String[] args) {
      ListNode l1=new ListNode(5);

      ListNode l2=new ListNode(2);
      System.out.println(CountNumber(l1));
   }
}


   class ListNode {
      int val;
      ListNode next;
      ListNode() {}
      ListNode(int val) { this.val = val; }
      ListNode(int val, ListNode next) { this.val = val; this.next = next; }
  }
