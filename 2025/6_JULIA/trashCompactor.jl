function solve(input::String)
    lines = split(chomp(input), '\n')
    height = length(lines)
    width = maximum(length.(lines))

    lines = [line * repeat(" ", width - length(line)) for line in lines]

    total = 0
    column = 1

    while column <= width
        println(join(column <= length(line) ? line[column] : ' ' for line in lines))
        column += 1
    end

end

input = read(joinpath(@__DIR__, "exampleInput.txt"), String)
solve(input)
# println(solve(input))
