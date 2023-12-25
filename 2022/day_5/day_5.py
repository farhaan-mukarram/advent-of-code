def get_inventory(stacks):
    last_row = stacks.pop()

    inventory = {}
    for line in stacks:

        for idx, char in enumerate(last_row):

            if char.isdigit():
                item = line[idx]

                if item.isalpha():
                    inventory[int(char)] = [item] + inventory.get(int(char), [])
    return inventory


def rearrange_inventory(inventory, rearrangement_steps):
    for instruction in rearrangement_steps:
        temp = [int(s) for s in instruction.split() if s.isdigit()]
        number_of_crates_to_move, source, destination = temp

        # picking crates
        crates_to_move = inventory[source][-number_of_crates_to_move:]

        # removing crates from source stack
        inventory[source] = inventory[source][:-number_of_crates_to_move]

        # adding crates to destination stack
        inventory[destination].extend(crates_to_move[::-1])


def get_message(inventory):
    temp = []

    for k, v in sorted(inventory.items()):
        temp.append(v[-1])

    message = "".join(temp)
    return message


with open("day_5.txt") as file:
    stacks, rearrangement_steps = file.read().split("\n\n")

stacks = stacks.splitlines()
rearrangement_steps = rearrangement_steps.splitlines()

inventory = get_inventory(stacks=stacks)

rearrange_inventory(inventory=inventory, rearrangement_steps=rearrangement_steps)

message = get_message(inventory=inventory)
print(message)
