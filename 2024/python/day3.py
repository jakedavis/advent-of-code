import functools
import re
from typing import List

from util import input_file

def parse_instruction(match: str) -> List[int]:
    match = match.lstrip("mul(")
    match = match.rstrip(")")
    x, y = match.split(",")

    return int(x), int(y)

def match_clause(clause: str) -> List[int]:
    matches = re.findall(r"mul\(\d+,\d+\)", clause)
    return [ parse_instruction(match) for match in matches ]

def match_clauses(input: str) -> List[List[int]]:
    splits = input.split("do()")
    elements = []
    for split in splits:
        parts = split.split("don't()")
        elements.append(match_clause(parts[0]))

    return elements

def part1() -> None:
    f = input_file("day3", is_test=False)
    instructions = match_clause(open(f, "r").read())
    multiples = [ x * y for x, y in instructions ]
    result = functools.reduce(lambda x, y: x + y, multiples)
    print(f"[Day 3][Part 1] result={result}")

def part2() -> None:
    f = input_file("day3", is_test=False)
    instruction_sets = match_clauses(open(f, "r").read())
    
    results = []
    for set in instruction_sets:
        multiples = [ x * y for x, y in set ]
        result = functools.reduce(lambda x, y: x + y, multiples)
        results.append(result)

    final = functools.reduce(lambda x, y: x + y, results)
    print(f"[Day 3][Part 2] result={final}")

part1()
part2()