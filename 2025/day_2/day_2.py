from utils import parse_input
from p1 import part_1
from p2 import part_2

if __name__ == "__main__":
    puzzle_input = parse_input()
    p1_solution = part_1(puzzle_input)
    p2_solution = part_2(puzzle_input)
    print(f"The sum of invalid ids (part 1) is {p1_solution}")
    print(f"The sum of invalid ids (part 2) is {p2_solution}")
