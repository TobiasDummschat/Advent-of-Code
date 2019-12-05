# Day 1 of the Advent of Code 2019

# read the puzzle input from a file and split into single lines
with open("day1_input", "r") as file:
    comp_mass_strings = file.read().split("\n")

# the input might end with an empty input. Remove it
if comp_mass_strings[-1] == "":
    comp_mass_strings = comp_mass_strings[:-1]

# parse the strings to get integers
comp_masses = [int(string) for string in comp_mass_strings]


# for each component, we need the fuel requirement of a single component
def get_comp_fuel_requirement(mass):
    return mass // 3 - 2


# now sum all the fuel requirements
total_fuel_required = 0
for mass in comp_masses:
    total_fuel_required += get_comp_fuel_requirement(mass)

print(total_fuel_required)
