DIRECTIONS = {
    "NORTH": {"line_index": -1, "char_index": 0},
    "SOUTH": {"line_index": 1, "char_index": 0},
    "WEST": {"line_index": 0, "char_index": -1},
    "EAST": {"line_index": 0, "char_index": 1},
    "NORTHEAST": {"line_index": -1, "char_index": 1},
    "NORTHWEST": {"line_index": -1, "char_index": -1},
    "SOUTHEAST": {"line_index": 1, "char_index": 1},
    "SOUTHWEST": {"line_index": 1, "char_index": -1},
}


def get_search_directions(
    line_index: int, char_index: int, total_number_of_lines: int, line_length
):
    search_directions = None

    # first line
    if line_index == 0:
        # first character
        if char_index == 0:
            search_directions = [
                DIRECTIONS["EAST"],
                DIRECTIONS["SOUTH"],
                DIRECTIONS["SOUTHEAST"],
            ]

        # last character
        elif char_index == line_length - 1:
            search_directions = [
                DIRECTIONS["WEST"],
                DIRECTIONS["SOUTH"],
                DIRECTIONS["SOUTHWEST"],
            ]

        else:
            search_directions = [
                DIRECTIONS["WEST"],
                DIRECTIONS["EAST"],
                DIRECTIONS["SOUTH"],
                DIRECTIONS["SOUTHEAST"],
                DIRECTIONS["SOUTHWEST"],
            ]

    # last line
    elif line_index == total_number_of_lines - 1:
        # first character
        if char_index == 0:
            search_directions = [
                DIRECTIONS["NORTH"],
                DIRECTIONS["EAST"],
                DIRECTIONS["NORTHEAST"],
            ]

        # last character
        elif char_index == line_length - 1:
            search_directions = [
                DIRECTIONS["NORTH"],
                DIRECTIONS["WEST"],
                DIRECTIONS["NORTHWEST"],
            ]

        else:
            search_directions = [
                DIRECTIONS["NORTH"],
                DIRECTIONS["WEST"],
                DIRECTIONS["EAST"],
                DIRECTIONS["NORTHEAST"],
                DIRECTIONS["NORTHWEST"],
            ]

    else:
        # first character
        if char_index == 0:
            search_directions = [
                DIRECTIONS["NORTH"],
                DIRECTIONS["SOUTH"],
                DIRECTIONS["EAST"],
                DIRECTIONS["NORTHEAST"],
                DIRECTIONS["SOUTHEAST"],
            ]

        # last character
        elif char_index == line_length - 1:
            search_directions = [
                DIRECTIONS["NORTH"],
                DIRECTIONS["SOUTH"],
                DIRECTIONS["WEST"],
                DIRECTIONS["NORTHWEST"],
                DIRECTIONS["SOUTHWEST"],
            ]

        else:
            search_directions = [value for value in DIRECTIONS.values()]

    return search_directions


def part_2(lines):
    initial_state = "".join(lines)
    current_state = initial_state
    prev_state = ""

    # iterate until no further removals possible
    while current_state != prev_state:
        prev_state = current_state

        for line_index in range(len(lines)):
            line = lines[line_index]

            for char_index in range(len(line)):
                adjacent_toilet_roll_count = 0
                char = line[char_index]

                # toilet roll found, determine adjacent
                if char == "@":
                    search_directions = get_search_directions(
                        line_index=line_index,
                        char_index=char_index,
                        total_number_of_lines=len(lines),
                        line_length=len(line),
                    )

                    # count adjacent toilet rolls
                    for direction in search_directions:
                        relative_line_index, relative_char_index = (
                            direction["line_index"],
                            direction["char_index"],
                        )

                        adjacent_char = lines[relative_line_index + line_index][
                            relative_char_index + char_index
                        ]

                        if adjacent_char == "@":
                            adjacent_toilet_roll_count += 1

                    if adjacent_toilet_roll_count < 4:
                        # mark as picked up
                        lines[line_index] = (
                            line[:char_index] + "." + line[char_index + 1 :]
                        )

        current_state = "".join(lines)

    initial_toilet_roll_count = initial_state.count("@")
    final_toilet_roll_count = current_state.count("@")

    return initial_toilet_roll_count - final_toilet_roll_count
