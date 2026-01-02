with open("input.txt", "r") as f:
    content = f.read().strip()

part1, part2 = content.split("\n\n")

freshRangeArray = part1.splitlines()
ingredientRange = list(map(int, part2.splitlines()))

freshCounter = 0
rangeCounter = 0

for ingredient in ingredientRange:
    for freshRange in freshRangeArray:
        minRange, maxRange = map(int, freshRange.split('-'))
        if minRange <= ingredient <= maxRange:
            freshCounter += 1
            break
        
print("Fresh available ingredients:", freshCounter)