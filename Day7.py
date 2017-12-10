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


# Reads the file, skipping the empty lines
with open('A7data.txt') as f_in:
    lines = (line.rstrip() for line in f_in) # All lines including the blank ones
    lines = (line for line in lines if line) # Non-blank lines
    # Turns it into a List
    tralines = list(lines)


    Layers = [] # List of layers (Last, Middle, ..., First)
    NaLay  = [] # Same list of layers, but this time, with only the name
    linnum = len(tralines)
    assinum = 0 # Assigned points so far
    couLay = 0 # COunts the layers

    # Variable that controls the loop
    Still_data = True
    while Still_data == True:
        Layers.append([]) # Adds one extra layer
        NaLay.append([])

        for this_line in tralines:
            seplin = this_line.split('->')
            # Only obtains the name, not the weight
            namlin = seplin[0].split(' (')[0]

            if len(Layers) == 1:
                # Only for the first run, selects the end of the tower
                if len(seplin) == 1:
                    Layers[-1].append(seplin[0])
                    NaLay[-1].append(namlin)
                    assinum += 1 # Adds one to the points assigned count

            else:

                # Skips points already covered
                for elelayer in NaLay[:-1]:
                    if namlin in elelayer:
                        Cov_point = True
                        break
                else:
                    Cov_point = False

                if Cov_point == True:
                    continue

                # In all other cases, it adds the subsequent point to the list
                # Assuming that they lead to a point already in the list

                # Separates the second part
                separ2 = seplin[1]
                # Separate after eliminating spaces
                act_separ2 = separ2.replace(' ', '').split(',')
                        # Checks if in the one before
                for befyer in act_separ2:
                    if (befyer in NaLay[-2]) and (namlin not in NaLay[-1]):
                        Layers[-1].append(seplin[0])
                        NaLay[-1].append(namlin)
                        assinum += 1


        # Ends the loop when all points are covered
        if assinum == linnum:
            Still_data = False

        print('Rows covered: '+str(couLay))
        couLay += 1

# Plots the resultant nodes
for qq in range(0, len(NaLay)):
    print('Row: '+str(qq)+'; starting from the end')
    for hh, zz in zip(NaLay[qq], Layers[qq]):
        print(zz, hh)

print(NaLay)
print(Layers)

# Creates a tree of lists
Tree_names = []
Tree_wei = []

# Reads the list again but this time usisng layers
# Thankfully the points are in the same order
# Inverts both lists first
Inv_Layers = Layers[::-1]
Inv_NaLay = NaLay[::-1]


"""
VISUALIZATION
"""

# Plots the results to have a better visual
