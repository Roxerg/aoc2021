f = open("input.txt", "r")


input = f.read()
input = input.split("\n")
input = [i.split(" | ") for i in input]
input = [[i[0].split(" "), i[1].split(" ")] for i in input]
print(input[0])

