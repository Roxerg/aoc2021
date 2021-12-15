import math
f = open("input.txt", "r")

input = f.read()

input = [ [int(i) for i in list(x)] for x in input.split("\n") ]


class coord:
    def __init__(self,y,x,v=0):
        self.x = x
        self.y = y
        self.v = v
    def __eq__(self, n):
        return hasattr(n, 'x') and hasattr(n, 'y') and self.x == n.x and self.y == n.y
    def __hash__(self):
        return hash("{}-{}".format(self.y, self.x))


def stupidStinkyAStar(grid):
    start = coord(0,0)
    end = coord(len(grid)-1, len(grid[0])-1)
    end.v = grid[end.y][end.x]

    allNodes = set()
    frontline = set()
    path = dict()

    gScore = dict()
    fScore = dict()

    for y in range(0, len(grid)):
        for x in range(0,len(grid[y])):
            p = coord(y,x,grid[y][x])
            gScore[p] = math.inf
            fScore[p] = math.inf
            # allNodes.add[]






stupidStinkyAStar(input)