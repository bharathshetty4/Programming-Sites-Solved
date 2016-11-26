class Solution:
    # Write your code here
    stackstring = []
    queuestring = []
    def pushCharacter(self,ch):
        self.stackstring.extend(ch)
        
    def enqueueCharacter(self,ch):
        self.queuestring.extend(ch)

    def popCharacter(self):
        char = self.stackstring.pop()
        return char
    def dequeueCharacter(self):
        char = self.queuestring.pop(0)
        return char
