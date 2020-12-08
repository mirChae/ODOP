def solution(n):
    i = 0
    j = 0
    l = []
    while n:
        l.append(n & 0x01)
        n >>= 1

    l.reverse()

    for c in l:
        if c == 0:
            j += 1
        else:
            if j > i:
                i = j
            j = 0

    return i

if __name__ == "__main__":
    while True:
        k = input()
        print(solution(int(k)))

