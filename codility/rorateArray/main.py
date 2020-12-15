def solution(a, k):
    for i in range(k):
        a = a[-1:] + a[:-1]

    return a

if __name__ == "__main__":
    a = [3,8,9,7,6]
    k = 3
    print(solution(a, k))

