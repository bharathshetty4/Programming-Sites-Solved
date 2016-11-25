#!/usr/bin/python


input_count = input().strip().split()
frnd_count = int(input_count[1])
special_frdns = input().strip().split()
special_frdns = list(map(int, special_frdns))
spl_posts = {}
normal_posts = {}
while (frnd_count):
    entry = input().strip().split()
    if (int(entry[0]) in special_frdns):
        spl_posts[int(entry[1])] = entry[2]
    else:
        normal_posts[int(entry[1])] = entry[2]
    frnd_count = frnd_count - 1

for i in sorted(spl_posts, reverse=True):
    print(spl_posts[i])

for i in sorted(normal_posts, reverse=True):
    print(normal_posts[i])
