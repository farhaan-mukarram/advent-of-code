with open("day_3.txt") as file:
    puzzle_input = file.read().strip().splitlines()


def find_common_item(compartment_one: str, compartment_two: str):
    for char in compartment_one:
        if char in compartment_two:
            return char

    return ""


def determine_item_priority(item: str):
    if item.isupper():
        return ord(item) - 38
    elif item.islower():
        return ord(item) - 96

    return 0


sum_of_priorities = 0

for line in puzzle_input:
    compartment_length = len(line)
    first_compartment, second_compartment = (
        line[: compartment_length // 2],
        line[compartment_length // 2 :],
    )

    common_item = find_common_item(
        compartment_one=first_compartment, compartment_two=second_compartment
    )

    sum_of_priorities += determine_item_priority(common_item)

print(sum_of_priorities)
