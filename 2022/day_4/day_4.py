with open("day_4.txt") as file:
    puzzle_input = file.read().strip().splitlines()


def get_section_start_and_end(section: str):
    return section.split("-")


def does_section_fully_contains_other(section_one, section_two):
    section_one_start, section_one_end = get_section_start_and_end(section=section_one)
    section_two_start, section_two_end = get_section_start_and_end(section=section_two)

    set_one = set([x for x in range(int(section_one_start), int(section_one_end) + 1)])
    set_two = set([x for x in range(int(section_two_start), int(section_two_end) + 1)])

    return set_one.issubset(set_two) or set_two.issubset(set_one)


overlaps = 0

for line in puzzle_input:
    res = line.split(",")
    section_one, section_two = res[0], res[1]

    is_fully_contained = does_section_fully_contains_other(
        section_one=section_one, section_two=section_two
    )

    if is_fully_contained:
        overlaps += 1

print(overlaps)
