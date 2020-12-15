def solution(A):
    d = {}
    for i in A:
        d.setdefault(i, 0)
        d[i] += 1


    for k, v in d.items():
        if v % 2 == 1:
            return k

if __name__ == '__main__':
    A = [9, 3, 9, 3, 9, 7, 9]
    print(solution(A))
