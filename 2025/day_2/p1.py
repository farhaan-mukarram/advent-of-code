def is_id_valid(id: str) -> bool:

    id_length = len(id)

    is_odd = id_length % 2

    # odd ids are valid
    if is_odd:
        return True

    left_half = id[: id_length // 2]
    right_half = id[id_length // 2 :]

    return left_half != right_half


def part_1(id_list):
    invalid_ids_sum = 0
    for id_range in id_list:
        start_id, end_id = id_range

        for id in range(int(start_id), int(end_id) + 1):
            is_id_invalid = not is_id_valid(str(id))

            if is_id_invalid:
                invalid_ids_sum += id

    return invalid_ids_sum
