def parse_input():
    with open("input.txt") as f:
        data = f.read().strip().split(",")

        id_list = []

        for item in data:
            first_id, second_id = item.split("-")
            id_list.append((first_id, second_id))

        return id_list
