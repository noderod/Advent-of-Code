"""
Carlos Redondo
Started December 10th, 2017
"""

"""
BASICS

Structure a tower where each element only returns its weight and the elements
immediately above it.
"""


import matplotlib.pyplot as plt
import itertools as itt



# Reads the file, skipping the empty lines
with open('A7data.txt') as f_in:
    lines = (line.rstrip() for line in f_in) # All lines including the blank ones
    lines = (line for line in lines if line) # Non-blank lines
    # Turns it into a List
    tralines = list(lines)

    # Lists of content
    WeiLay = [] # Contains weights
    NaLay = [] # Contains names

    linnum = len(tralines)
    assinum = 0 # Assigned points so far
    rowcom = 0


    Still_Data = True
    while Still_Data == True:

        print('Row: '+str(rowcom))

        for each_line in tralines:
            seplin = each_line.split('->')
            # Only obtains the name, not the weight
            namlin = seplin[0].split(' (')[0]
            skip_thing = False

            # In the case of the end of the tower
            if len(seplin) == 1:

                # Skips if the point has already been assigned
                if rowcom == 0:
                    # Only valid for first row
                    if namlin in NaLay:
                        continue
                else:
                    for this_BBB in NaLay:
                        if namlin in this_BBB:
                            skip_thing = True

                    if skip_thing == True:
                        continue

                    WeiLay.append([seplin[0]])
                    NaLay.append([namlin])
                    continue


            # Skips points already assigned
            for loltz in range(0, len(NaLay)):
                leBranch = NaLay[loltz]
                if namlin in leBranch:
                    continue


                # For the rest, appends to the corresponding branch
                # Separates the second part
                separ2 = seplin[1]

                for each_elem in leBranch:

                    if each_elem in separ2:
                        WeiLay[loltz].append(seplin[0])
                        NaLay[loltz].append(namlin)


        # Computes the length of all the lists combined
        All_po = []

        for leitem in itt.chain.from_iterable(NaLay):
            if leitem not in All_po:
                All_po.append(leitem)

        assinum = len(All_po)

        # Stops when all lists finish in the same point
        finpo = []

        for yeti in range(1, len(NaLay)):
            if NaLay[yeti-1][-1] == NaLay[yeti][-1]:
                finpo.append(True)
            else:
                finpo.append(False)




        # Ends the program when all towers have been assigned
        if (assinum == linnum) and (all(finpo) == True):
            Still_Data = False

        if rowcom > 18:
            Still_Data = False



        rowcom += 1


"""
LIST CHANGES
"""

# Changes the way the system is designed to have a new path

# Inverted list of things
Inv_Nam = []
Inv_Wei = []

for braWEI, braNAM in zip(WeiLay, NaLay):
    Inv_Wei.append(braWEI[::-1])
    Inv_Nam.append(braNAM[::-1])

print('Starting point: '+str(Inv_Nam[0][0]))


"""
PART 2
"""

# Turns the list into a list of list, much more usable
TOWER = []

for wolver in Inv_Wei:

    print(wolver)
