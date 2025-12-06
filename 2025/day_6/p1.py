def part_1(problems_list: list[list[str]], operators_list: list[str]):
    results = []

    no_of_problems = len(problems_list[0])
    no_of_operands = len(problems_list)

    for i in range(no_of_problems):
        operation = operators_list[i]

        res = 1 if operation == "*" else 0

        for j in range(no_of_operands):
            operand = int(problems_list[j][i])

            if operation == "*":
                res *= operand
            elif operation == "+":
                res += operand

        results.append(res)

    return sum(results)
