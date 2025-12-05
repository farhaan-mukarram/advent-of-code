def parse_input():
    with open("input.txt") as f:
        data = f.read().strip().split("\n")

        return data


def get_ranges_and_ingredients(data: list[str]):
    ingredients_list = [x for x in data if len(x.strip()) > 0 and "-" not in x]
    range_list = [x for x in data if len(x.strip()) > 0 and "-" in x]

    return range_list, ingredients_list
