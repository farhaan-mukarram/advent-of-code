from utils import get_ranges_and_ingredients, parse_input
from p1 import part_1

if __name__ == "__main__":
    puzzle_input = parse_input()
    ranges, ingredients = get_ranges_and_ingredients(puzzle_input)
    p1_solution = part_1(ranges, ingredients)
    print(f"Num of fresh ingredients (part 1) = {p1_solution}")
