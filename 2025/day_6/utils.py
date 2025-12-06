import re


def parse_input():
    with open("input.txt") as f:
        data = f.read().strip().split("\n")

        return data


def get_problems_and_operators(data: list[str]):
    problems, operators = data[:-1], data[-1:]

    problems_list = [re.split(r"\s+", x.strip()) for x in problems]
    operators_list = [re.split(r"\s+", x.strip()) for x in operators][0]

    return problems_list, operators_list
