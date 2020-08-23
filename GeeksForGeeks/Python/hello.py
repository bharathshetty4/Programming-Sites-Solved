from functools import reduce
import json

print ("Hello World!")

#python debugging
#command to be run " python -m pdb my_script.py "




#arguments
#single star (*)
def test_var_args(f_arg, *argv):
    print("first normal arg:", f_arg)
    for arg in argv:
        print("another arg through *argv:", arg)

test_var_args('yasoob', 'python', 'eggs', 'test')

#double star (**)
def greet_me(**kwargs):
    for key, value in kwargs.items():
        print("{0} = {1}".format(key, value))

greet_me(name="yasoob")



#python datastructures ( list, dicts, tuples, sets, strings,frozensets)


#sets , shows  only unique values. Have intersection, differences etc
some_list = ['a', 'b', 'c', 'b', 'd', 'm', 'n', 'n']
duplicates = set([x for x in some_list if some_list.count(x) > 1])  #set defined here
print("Sets example", duplicates)

#array
a= [1,"c"]
for i in a:
    print("Array ex ", i)




#dictionary
d = {'x': 1}
for key, value in d.items():
    print("dicttionary example ",key,value)




#tuples

t = ('a','b')
print ("tuples ", t)



#list
l = [1,2]
print("list exaample", l)



#list comprehensions
#used to construct lists in a very natural, easy way, like a mathematician is used to do.
x = [x**2 for x in range(0,5)]
print ("list comprehensions ",x,type(x))

#prime using list comprehensions
noprimes = [j for i in range(2, 8) for j in range(i*2, 50, i)]
primes = [x for x in range(2, 50) if x not in noprimes]
print("prime using list comprehensions ",primes)


#lambda in map
items = [1, 2, 3, 4, 5]
squared = list(map(lambda x: x**2, items))
print("lambda in map ",squared)

#lambda in filter
#creates a list of elements for which a function returns true
number_list = range(-5, 5)
less_than_zero = list(filter(lambda x: x < 0, number_list))
print("lambda in filter ",less_than_zero)


#lambda in reduce
#performing some computation on a list and returning the result
product = reduce((lambda x, y: x * y), [1, 2, 3, 4])
print("lambda in reduce ",product)


#ternary operator

is_fat = True
state = "fat" if is_fat else "not fat"
print ("Ternary example ", state)



#repr() and str()
#repr() is mainly used for debugging and development
s = 'Hello, Geeks.'
print ("str() ex ", str(s), str(2.0/11.0))
print ("repr() ex ", repr(s), repr(2.0/11.0))



#iterator, next, generator, yield
def generator_function():
    for i in range(3):
        yield i

gen = generator_function()
print("generator next(0)",next(gen)) # Output: 0
print("generator next(1)",next(gen)) # Output: 1






# try, catch, exception, finally

try:
    file = open('test.txt', 'rb')
except IOError as e:
    print("Except example: An error occurred.")
    #raise e
finally:
    print("Finally example: This would be printed whether or not an exception occurred!")



#Control statements

# has pass, continue, break options

#for loop
for i in range (0,1):
    print("for loop ex: ",i)



#while loop
count = 0
while (count < 1):
   print ('While loop ex: ', count)
   count = count + 1







#json encoding and decoding
data = {'apple': 'cat', 'banana':'dog', 'pear':'fish'}
dump_data = json.dumps(data)
print("Dumped json data is ", dump_data)
load_data = json.loads(dump_data)
print("Loaded json data is ", load_data)












#classes and multiple inheritance

"""
class X: pass
class Y: pass
class Z: pass

class A(X,Y): pass    // multiple inheritance
class B(Y,Z): pass

class M(B,A,Z): pass   // multilevel inheritance

# Output:
# [<class '__main__.M'>, <class '__main__.B'>,
# <class '__main__.A'>, <class '__main__.X'>,
# <class '__main__.Y'>, <class '__main__.Z'>,
# <class 'object'>]

print(M.mro())

"""






#default dict


#multithreading
"""
import thread
def shout(n):
    print("Thread here", ('n',))

thread.start_new_thread(shout,1)
"""

