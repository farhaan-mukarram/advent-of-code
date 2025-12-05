def part_1(ranges: list[str], ingredients: list[str]):
    fresh_ingredients_count = 0

    for ingredient in ingredients:
        is_ingredient_fresh = False

        for r in ranges:
            start, end = int(r.split("-")[0]), int(r.split("-")[1])

            if int(ingredient) >= start and int(ingredient) <= end:
                is_ingredient_fresh = True

        if is_ingredient_fresh:
            fresh_ingredients_count += 1

    return fresh_ingredients_count
