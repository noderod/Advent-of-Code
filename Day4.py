"""
Carlos Redondo
December 11th, 2017
"""

"""
BASICS

Obtain how many passphares are valid.
"""



"""
PART 1
"""

# A passphrase is considered valid if it has the same length than a set with
# no repeated elements

Nonrep1 = 0 # Final result

# Reads the file, skipping the empty lines
with open('A4data.txt') as f_in:
    lines = (line.rstrip() for line in f_in) # All lines including the blank ones
    lines = (line for line in lines if line) # Non-blank lines

    # Divides the line by spaces
    for prelin in lines:

        divlin = prelin.split()
        if len(divlin) == len(set(divlin)):
            Nonrep1 += 1

print('Part 1, Non-repeated pass phrases: ', Nonrep1)



"""
PART 2
"""

# A pass phrase must not contain an anagram of other letters

Nonrep2 = 0


# Counts the number of letters of each in a string
# young (string): String which letters are to be counted
# Return format: [[Char 1, Count], [Char 2, Count], ...]

def alllet(young):

    # List of all letters in the string
    abece = []
    # Final result
    fullcount = []
    for char in young:
        if char not in abece:
            abece.append(char)
            fullcount.append([char, young.count(char)])

    # Returns the list sorted by numerical order
    return sorted(fullcount, key = lambda x: x[0])


with open('A4data.txt') as f_in:
    lines = (line.rstrip() for line in f_in)
    lines = (line for line in lines if line)

    for prelin in lines:

        divlin = prelin.split()
        if len(divlin) != len(set(divlin)):
            # First condition failed, no need to continue
            continue

        # Computes all the repeated anagrams
        gramated = []
        for iren in divlin:
            gramated.append(alllet(iren))

        # Filters the anagrams so that only the non-repeated ones remain
        anfil = []
        for vict in gramated:
            if vict not in anfil:
                anfil.append(vict)

        # If all conditions are met, the phrase is valid
        if len(divlin) == len(anfil):
            Nonrep2 += 1


print('Non-repeated pass phrases: ', Nonrep2)
