# Day 3 of the Advent of Code 2019

# get the puzzle input
with open("day3_input", "r") as file:
    # get the lines of the input, which is in CSV format
    puzzle_input = file.read().split("\n")

    # split each line into an array of the individual instructions
    wire1 = puzzle_input[0].split(",")
    wire2 = puzzle_input[1].split(",")


def parse_wire(wire_instructions):
    """
    Takes an array of instructions and outputs the series of points the wire traverses as an array of integer pairs
    :param wire_instructions: The array containing the wires instructions
    :return: The points the wire traverses
    """

    # We build the wires path recursively using the following helper function
    def build_wire_path(wire_instructions, partial_path):
        # this constructs the wires path end-recursively storing the already constructed path as partiaL_path

        # recursion start
        if len(wire_instructions) == 0:
            return partial_path

        # get and parse the next instruction
        instruction = wire_instructions.pop(0)
        direction = instruction[0]
        distance = int(instruction[1:])

        # get the coordinates of the last point
        x, y = partial_path[-1]

        # go for distance in direction from last point
        if direction == "U":
            extension = [(x, y + i) for i in range(1, distance + 1)]
            partial_path.extend(extension)
        elif direction == "D":
            extension = [(x, y - i) for i in range(1, distance + 1)]
            partial_path.extend(extension)
        if direction == "R":
            extension = [(x + i, y) for i in range(1, distance + 1)]
            partial_path.extend(extension)
        elif direction == "L":
            extension = [(x - i, y) for i in range(1, distance + 1)]
            partial_path.extend(extension)

        # now the wire instructions are one element shorter and the partial_path is longer correspondingly

        return build_wire_path(wire_instructions, partial_path)

    # build the path starting from the center
    return build_wire_path(wire_instructions, [(0, 0)])


path1 = parse_wire(wire1)
path2 = parse_wire(wire2)

# find intersection
intersections = []

# we sort the paths for a quicker search
sorted_path1 = sorted(path1)
sorted_path2 = sorted(path2)

# Interate through the paths simultaneously, always incrementing the index of the path with the smallest current element
i, j = 0, 0

while i < len(sorted_path1) and j < len(sorted_path2):
    # add intersection, if sorted_paths agree on the point, then increment i
    if sorted_path1[i] == sorted_path2[j]:
        intersections.append(sorted_path1[i])
        i += 1
    # otherwise increment the index pointing to the smaller element
    elif sorted_path1[i] < sorted_path2[j]:
        i += 1
    elif sorted_path1[i] > sorted_path2[j]:
        j += 1

# the center does not count. remove it
intersections.remove((0, 0))

# get manhatten distances to center at (0,0)
distances = [abs(x) + abs(y) for x, y in intersections]

print("Part 1: ", min(distances))

# now we calculate the delay for each of the intersections
delays = []

for intersection in intersections:
    # the delay is just the index of the intersections first appearance in the cable
    delay1 = path1.index(intersection)
    delay2 = path2.index(intersection)

    delays.append(delay1 + delay2)

print("Part 2: ", min(delays))
