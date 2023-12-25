def get_start_marker(data_stream, number_of_previous_distinct_characters=4):
    for i in range(number_of_previous_distinct_characters - 1, len(data_stream)):
        window = data_stream[i - (number_of_previous_distinct_characters - 1) : i + 1]

        if len(set(window)) == number_of_previous_distinct_characters:
            return i + 1


with open("day_6.txt") as file:
    puzzle_input = file.read()

start_of_packet_marker = get_start_marker(
    puzzle_input, number_of_previous_distinct_characters=4
)

start_of_message_marker = get_start_marker(
    puzzle_input, number_of_previous_distinct_characters=14
)

print(
    "First start of packet marker was detected after character :",
    start_of_packet_marker,
)

print(
    "First start of message marker was detected after character :",
    start_of_message_marker,
)
