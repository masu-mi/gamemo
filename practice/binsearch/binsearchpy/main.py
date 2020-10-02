from typing import List

def binsearch(ar: List[int], l: int, r: int, cond) -> int:
    while abs(r-l) > 1:
        mid = l + ((r-l)>>1)
        if cond(ar, mid):
            r = mid
        else:
            l = mid
    return r

if __name__ == '__main__':
    length = int(input())
    target_list = [ int(input()) for i in range(length) ]
    boundary = int(input())
    print(target_list)
    print(binsearch(target_list, -1, len(target_list), lambda ar, i: ar[i] > boundary))
