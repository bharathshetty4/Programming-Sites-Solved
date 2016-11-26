    def insert(self,head,data):
        temp = Node(data)
        temp.next=None
        current = head
        if(head==None):
            current=temp
        else:
            while(head.next!=None):
                head=head.next
            head.next=temp
        return current
      
