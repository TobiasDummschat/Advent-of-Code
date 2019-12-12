# Day 4 of the Advent of Code 2019

# get the puzzle input
with open("day4_input", "r") as file:
    puzzle_input = file.read().split("-")
    left = int(puzzle_input[0])
    right = int(puzzle_input[1])

count_part1 = 0
count_part2 = 0
for i in range(left, right + 1):
    i_str = str(i)

    # test 6 digits
    if len(i_str) != 6:
        continue

    # test ascending digits
    has_asc_digits = True
    # run through all five adjacent pairs
    for j in range(5):
        # if descending digits found, set False and break
        if int(i_str[j + 1]) < int(i_str[j]):
            has_asc_digits = False
            break

    if not has_asc_digits:
        continue

    # test identical pairs
    has_identical_pair = False
    has_isolated_identical_pair = False
    # run through all five adjacent pairs
    for j in range(5):
        # if identical pair is found, set True and test if it is isolated
        if i_str[j] == i_str[j + 1]:
            has_identical_pair = True
            # Isolation array border cases before general case
            if (j == 0 and i_str[j] != i_str[j + 2]) or \
                    (j == 4 and i_str[j] != i_str[j - 1]) or \
                    (i_str[j] != i_str[j - 1] and i_str[j] != i_str[j + 2]):
                has_isolated_identical_pair = True
                break

    if not has_identical_pair:
        continue

    # if we reach this point, we have only ascending digits and an identical pair
    count_part1 += 1
    if has_isolated_identical_pair:
        count_part2 += 1

print("Part 1: ", count_part1)
print("Part 2: ", count_part2)
