def rotate(direction: str, value: int, prev_dial_position: int = 50) -> tuple[int, int]:
    dial_position = prev_dial_position
    zeroes_count = 0

    clicks = value

    while clicks != 0:
        if direction == "L":
            dial_position -= 1

            if dial_position < 0:
                dial_position += 100

        elif direction == "R":
            dial_position += 1

            if dial_position > 99:
                dial_position -= 100

        if dial_position == 0:
            zeroes_count += 1

        clicks -= 1

    return dial_position, zeroes_count


def part_2(rotation_list) -> int:
    zeroes_count = 0
    dial_position = 50

    for rotation in rotation_list:
        direction, rotation_value = rotation[:1], int(rotation[1:])
        dial_position, num_zeroes = rotate(
            direction=direction,
            prev_dial_position=dial_position,
            value=rotation_value,
        )

        zeroes_count += num_zeroes

    return zeroes_count
