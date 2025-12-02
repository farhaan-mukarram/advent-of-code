def is_id_valid(id: str) -> bool:
    id_length = len(id)
    is_valid = True

    for i in range(id_length // 2 + 1):
        window = id[:i]
        window_length = len(window)

        if window_length == 0:
            continue

        count = id.count(window)

        if count < 2:
            continue

        if (count * window_length) == id_length:
            is_valid = False
            break

    return is_valid


def part_2(id_list):
    invalid_ids_sum = 0
    for id_range in id_list:
        start_id, end_id = id_range

        for id in range(int(start_id), int(end_id) + 1):
            is_id_invalid = not is_id_valid(str(id))

            if is_id_invalid:
                invalid_ids_sum += id

    return invalid_ids_sum
