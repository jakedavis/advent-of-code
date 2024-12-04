from typing import List
from util import input_file

def is_safe(report: List[int]) -> bool:
    increasing, decreasing = False, False

    print(report, end=" ")
    for i in range(len(report)-1):
        i1 = report[i]
        i2 = report[i+1]

        diff = i2 - i1
        if diff == 0:
            print("no increase or decrease")
            return False
        if abs(diff) > 3 or abs(diff) < 1:
            print("diff too large")
            return False
        
        if diff in [1, 2, 3]:
            increasing = True
        if diff in [-1, -2, -3]:
            decreasing = True

    if increasing and decreasing:
        print("increasing and decreasing")
        return False
    
    print("OK")
    return True

def is_safe_dampened(report: List[int]) -> bool:
    if is_safe(report):
        return True

    for r in report:
        subreport = report.copy()
        subreport.remove(r)
        if is_safe(subreport):
            return True

def parse_input(path: str) -> List[List[int]]:
    return [[int(i) for i in line] for line in [ l.split(" ") for l in open(path, "r").read().splitlines() ]]

def part1():
    reports = parse_input(input_file("day2", is_test=False))
    
    num_safe_reports = 0
    for report in reports:
        if is_safe(report):
            num_safe_reports += 1

    print(f"[Day 2][Part 1] Number of safe reports: {num_safe_reports}")

def part2():
    reports = parse_input(input_file("day2", is_test=False))
    
    num_safe_reports = 0
    for report in reports:
        if is_safe(report):
            num_safe_reports += 1
        else:
            for i in range(len(report)):
                subreport = report.copy()
                del subreport[i]
                if is_safe(subreport):
                    num_safe_reports += 1
                    break

    print(f"[Day 2][Part 2] Number of safe reports: {num_safe_reports}")

#part1()
part2()