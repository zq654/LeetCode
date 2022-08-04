public class LeetCode_2 {
   public ListNode addTwoNumbers(ListNode l1, ListNode l2) {
     l1=Reversal(l1);//反转链表
     l2=Reversal(l2);
      ListNode l=new ListNode();
      BulidList(l1,l2,l,0);
      return Reversal(l.next);
   }

   private void BulidList(ListNode l1, ListNode l2,ListNode l, int i) {
      if(l1==null&&l2==null){
         if(i!=0)l.next=new ListNode(1,null);
         return;
      }else{
         if(l1==null){
            l.next=new ListNode((i+l2.val)%10,null);
            BulidList(l1,l2.next,l.next,(i+ l2.val)/10);
         }
         if(l2==null){
            l.next=new ListNode((i+l1.val)%10,null);
            BulidList(l1.next,l2,l.next,(l1.val+ i)/10);
         }
         if(l1!=null&&l2!=null){
            l.next=new ListNode((l1.val+l2.val+i)%10,null);
            BulidList(l1.next,l2.next,l.next,(l1.val+l2.val+i)/10);
         }
      }
   }

   private static ListNode Reversal(ListNode l) {
      ListNode head=new ListNode(0,null);
      ListNode pointer=l.next;
      while(l.next!=null){
         l.next=head.next;
         head.next=l;
         l=pointer;
         pointer=pointer.next;
      }
      l.next=head.next;
      head.next=l;
      return head.next;
   }

   public static void main(String[] args) {
      ListNode a=new ListNode(9,null);
      ListNode b=new ListNode(4,a);
      ListNode c=new ListNode(2,b);
      ListNode d=Reversal(c);
      System.out.println(d.val);
   }
}


   class ListNode {
      int val;
      ListNode next;
      ListNode() {}
      ListNode(int val) { this.val = val; }
      ListNode(int val, ListNode next) { this.val = val; this.next = next; }
  }
