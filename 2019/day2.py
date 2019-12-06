# Day 2 part 1 of the Advent of Code 2019

# get the puzzle input
with open("day2_input", "r") as file:
    puzzle_input = file.read()

# split the program and convert the entries to integers
program = [int(s) for s in puzzle_input.split(",")]


# method that runs the program
def run(program, replace_pos1, replace_pos2):
    # copy the program as to not change it while running
    memory = program.copy()

    # replace the given positions
    memory[1] = replace_pos1
    memory[2] = replace_pos2

    # start at instr_pointer 0
    instr_pointer = 0

    # loop until instr_pointer is at a value of 99
    while memory[instr_pointer] != 99:

        # both cases get the next two positions as input and the one after that as output
        input1_pointer = memory[instr_pointer + 1]
        input2_pointer = memory[instr_pointer + 2]
        output_pointer = memory[instr_pointer + 3]

        # addition case
        if memory[instr_pointer] == 1:
            memory[output_pointer] = memory[input1_pointer] + memory[input2_pointer]

        # multiplication case
        elif memory[instr_pointer] == 2:
            memory[output_pointer] = memory[input1_pointer] * memory[input2_pointer]

        # advance to next opcode
        instr_pointer += 4

    return memory[0]


# -------- Part 1 --------
print("Part 1: ", run(program, 12, 2))

# -------- Part 2 --------
# Try each possible combination of noun and verb and output 100 * noun + verb,
# if they match they result in the desired program output
# There might be more than one, so we try each combination.

for noun in range(0, 100):
    for verb in range(0, 100):
        if run(program, noun, verb) == 19690720:
            print("Part 2: ", 100 * noun + verb)
