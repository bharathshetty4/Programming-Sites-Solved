class Student(Person):
    def __init__(self, firstName, lastName, idNumber, scores):
        self.firstName = firstName
        self.lastName = lastName
        self.idNumber = idNumber
        self.scores=scores
    def calculate(self):
        sum = 0
        count = 0
        for i in scores:
            sum =sum +i
            count = count +1
        avg=sum/count
        if(avg<40):
            return 'T'
        elif(avg<55):
            return 'D'
        elif(avg<70):
            return 'P'
        elif(avg<80):
            return 'A'
        elif(avg<90):
            return 'E'
        return 'O'
