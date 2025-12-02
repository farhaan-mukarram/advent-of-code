def rotate(direction: str, value: int, prev_dial_position: int = 50) -> int:
    dial_position = prev_dial_position

    clicks = value

    while clicks != 0:
        if direction == "L":
            if dial_position == 0:
                dial_position = 99
            else:
                dial_position -= 1

        elif direction == "R":
            if dial_position == 99:
                dial_position = 0

            else:
                dial_position += 1

        clicks -= 1

    return dial_position


def part_1(rotation_list) -> int:
    zeroes_count = 0
    dial_position = 50

    for rotation in rotation_list:
        direction, rotation_value = rotation[:1], int(rotation[1:])
        dial_position = rotate(
            direction=direction,
            prev_dial_position=dial_position,
            value=rotation_value,
        )

        if dial_position == 0:
            zeroes_count += 1

    return zeroes_count
