

f = open("input.txt", "r")

input = f.read()
input = input.split("\n")


calls = input[0].split(",")
input = input[2:]

print(input[0])

for idx in range(0, len(input)):
    inputrow = input[idx].replace("  ", " ")
    inputrow = [x for x in inputrow.split(" ")] 
    print(inputrow)
    break