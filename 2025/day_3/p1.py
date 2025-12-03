def find_largest_joltage(nums: list[int]):
    i = 0
    list_length = len(nums)

    largest_combinations = []

    while i < list_length:
        num_1 = nums[i]

        j = i + 1

        while j < list_length:
            num_2 = nums[j]

            # concatenate both numbers and add to combinations
            largest_combinations.append(int(str(num_1) + str(num_2)))

            j += 1

        i += 1

    return max(largest_combinations)


def part_1(battery_bank_list):
    total_joltage = 0

    for bank in battery_bank_list:
        int_bank_list = [int(x) for x in bank]

        largest_joltage_possible = find_largest_joltage(int_bank_list)

        total_joltage += largest_joltage_possible

    return total_joltage
