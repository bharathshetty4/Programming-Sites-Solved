#https://www.geeksforgeeks.org/minimum-removal-make-palindrome-permutation/

#we can use the hash aswell here, as given in the above g4g link

def charRemove(rawString):
    total_char_count = 0
    char_dict = {}

    for single_char in rawString:
        if single_char in char_dict.keys():
            if char_dict[single_char] == 1 :
                char_dict[single_char] = 0
                total_char_count = total_char_count - 1
            else:
                char_dict[single_char] = 1
                total_char_count = total_char_count + 1
        else:
            char_dict[single_char] = 1
            total_char_count = total_char_count + 1

    # you can have one odd numbered character which you can fit it in the middle
    return total_char_count - 1

raw_string = "geeksforgeeks"
char_count = charRemove(raw_string)

print (char_count)