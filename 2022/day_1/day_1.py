# Running count of the total calories for an elf
total_calories = 0

# Keeping track of the largest count of total calories
most_calories_so_far = 0

with open("day_1.txt") as file:
    puzzle_input = file.readlines()


for line in puzzle_input:
    # line break means a new elf, reset the running count and update largest total so far
    if line == "\n":
        most_calories_so_far = max(most_calories_so_far, total_calories)
        total_calories = 0

    # update the running count and largest total so far
    else:
        total_calories += int(line)
        most_calories_so_far = max(most_calories_so_far, total_calories)


print("Calories carried by elf carrying the most calories:", most_calories_so_far)
