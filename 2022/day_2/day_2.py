ROUND_SCORE = {"WIN": 6, "DRAW": 3, "LOSS": 0}
SHAPE_SCORE = {"ROCK": 1, "PAPER": 2, "SCISSORS": 3}

OPPONENT_MOVES_MAP = {"A": "ROCK", "B": "PAPER", "C": "SCISSORS"}
PLAYER_MOVES_MAP = {"X": "ROCK", "Y": "PAPER", "Z": "SCISSORS"}


def determine_round_result(player_move: str, opponent_move: str):
    result = ""
    if opponent_move == player_move:
        result = "DRAW"

    elif opponent_move == "ROCK":
        if player_move == "PAPER":
            result = "WIN"
        else:
            result = "LOSS"

    elif opponent_move == "PAPER":
        if player_move == "SCISSORS":
            result = "WIN"
        else:
            result = "LOSS"

    elif opponent_move == "SCISSORS":
        if player_move == "ROCK":
            result = "WIN"
        else:
            result = "LOSS"

    return result


def determine_round_score(opponent_input: str, player_input: str):
    round_score = 0

    opponent_move = OPPONENT_MOVES_MAP[opponent_input]
    player_move = PLAYER_MOVES_MAP[player_input]

    round_result = determine_round_result(
        opponent_move=opponent_move, player_move=player_move
    )

    round_score = ROUND_SCORE[round_result] + SHAPE_SCORE[player_move]

    return round_score


with open("day_2.txt") as file:
    puzzle_input = file.read().splitlines()

total_score = 0
for line in puzzle_input:
    opponent_move, player_move = line.split()

    total_score += determine_round_score(opponent_move, player_move)

print(total_score)
