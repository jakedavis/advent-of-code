import os
from typing import List, Tuple

def input_file() -> str:
    parent_dir = os.path.dirname(os.path.realpath(__file__))
    return f"{parent_dir}/../inputs/day1"

def parse_input(path: str) -> Tuple[List[int], List[int]]:
    f = open(path, "r")
    input = f.read()

    left, right = [], []

    for line in input.splitlines():
        l, r = line.split("   ")
        left.append(int(l))
        right.append(int(r))

    return left, right

def part1():
    left, right = parse_input(input_file())

    left.sort()
    right.sort()
    zipped = zip(left, right)

    distance = 0
    for zipper in zipped:
        distance += abs(zipper[0] - zipper[1])

    print(f"distance: {distance}")

def part2():
    left, right = parse_input(input_file())
    
    similarity = 0
    for l in left:
        occurrences = 0

        for r in right:
            if l == r:
                occurrences += 1
        
        similarity += occurrences * l

    print(f"similarity score: {similarity}")

part1()
part2()
