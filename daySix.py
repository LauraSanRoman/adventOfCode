def main():
    # distro = [5,1,10,0,1,7,13,14,3,12,8,10,7,12,0,6]
    distro = [0,2,7,0]
    listOfDistros = [[0,2,7,0]]
    nonRepeat = True
    print(listOfDistros)
    x = 0
    while nonRepeat:

        maxVal = max(distro)
        maxIndex = distro.index(maxVal)
        addingToIndex = maxIndex+1
        distro[maxIndex] = 0
        for i in range(7):
            if(addingToIndex == len(distro)):
                addingToIndex = 0
            distro[addingToIndex] = distro[addingToIndex]+1
            addingToIndex = addingToIndex+1
        listOfDistros.append([distro])
        print(listOfDistros)
        x = x+1
        for item in listOfDistros:
            for num in distro:
                # if(item == distro):
                #     print(item)
                #     print(distro)
                #     nonRepeat = False
                #     print("it worked")

                if(x == 3):
                    nonRepeat = False


main()
