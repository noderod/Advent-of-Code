"""
Carlos Redondo
December 11th, 2017
"""

"""
BASICS

Count characters which are equal
"""


# PART 1
# The text file has been divided into more line sot make it more readable

with open('A1data.txt') as f_in:
    lines = (line.rstrip() for line in f_in) # All lines including the blank ones
    lines = (line for line in lines if line) # Non-blank lines

    # Final string of text
    Fins = ''
    for each_line in lines:
        #each_line.replace('\n', '')
        Fins += each_line



print(len(Fins))

counter = 0 # Final result

# Checks if an element is the same as the one in front
for elem in range(0, len(Fins)):
    if Fins[elem] == Fins[(elem + 1) % len(Fins)]:
        counter += float(Fins[elem])

print('Repeated elements: ', counter)


"""
PART 2

Count characters which are equal halfway accross the number
"""

jumper_counter = 0

for elem in range(0, len(Fins)):
    if Fins[elem] == Fins[(elem + int(0.5*len(Fins))) % len(Fins)]:
        jumper_counter += float(Fins[elem])

print('Repeated elements: ', jumper_counter)
