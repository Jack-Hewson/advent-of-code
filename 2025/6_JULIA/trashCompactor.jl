function solve(input::String)
    # Takes the input and converts into a long array
    lines = split(chomp(input), '\n')

    # Create a 3d array to represent each row + column
    parsed_lines = []
    for line in lines
        # Split line on spaces, remove empty strings
        tokens = filter(!isempty, split(line, ' '))
        push!(parsed_lines, tokens)
    end

    # 'enumerate' takes an array and returns the element and its index
    for (i, tokens) in enumerate(parsed_lines)
        println("Row $(i): ", tokens)
    end

    rows = length(parsed_lines)
    columns = maximum(length.(parsed_lines)) - 1

    problems = []
    for column_index in 1:columns
        numbers = Int[]
        operator = ""
        for row_index in 1:(rows-1)
            # add each number in a column to an array (except for operator)
            push!(numbers, parse(Int, parsed_lines[row_index][column_index]))
        end
        # operator is always last in column
        operator = parsed_lines[end][column_index]
        # assign to an array so numbers and operator are separate
        push!(problems, (numbers=numbers, operator=operator))
    end

    println("problems: ", problems)

    total = 0
    for problem in problems
        if problem.operator == "+"
            total += sum(problem.numbers)
        elseif problem.operator == "*"
            total += prod(problem.numbers)
        end
    end

    return total
end

input = read(joinpath(@__DIR__, "input.txt"), String)
println(solve(input))
