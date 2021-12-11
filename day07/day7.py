f = open("input.txt", "r")


input = f.read()
input = input.split(",")
input = list(map(int, input))

# part 1 to confirm median works

def median(l):
    half = len(l) // 2
    l.sort()
    if not len(l) % 2:
        return (l[half - 1] + l[half]) / 2.0
    return l[half]

med = median(input)

c = 0
for i in range(0, len(input)):
    c += abs(input[i] - med)

print(c)