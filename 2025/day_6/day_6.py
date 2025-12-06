from utils import get_problems_and_operators, parse_input
from p1 import part_1

if __name__ == "__main__":
    puzzle_input = parse_input()
    problems, operators = get_problems_and_operators(puzzle_input)
    p1_solution = part_1(problems, operators)
    print(f"Grand total (part 1) = {p1_solution}")
