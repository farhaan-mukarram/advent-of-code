from utils import parse_input
from p1 import part_1
from p2 import part_2

if __name__ == "__main__":
    puzzle_input = parse_input()
    p1_solution = part_1(puzzle_input)
    p2_solution = part_2(puzzle_input)
    print(f"Rolls of paper accessible by forklift (part 1) = {p1_solution}")
    print(f"Rolls of paper accessible by forklift (part 2) = {p2_solution}")
