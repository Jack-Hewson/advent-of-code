example_array = [
  987654321111111,
    811111111111119,
  234234234234278,
  818181911112111
]

content = File.read("input.txt").split("\n")

total_val = 0

content.each do |val|
  biggest_val = { index: nil, digit: -1 }
  second_biggest_val = { index: nil, digit: -1 }

  val.to_s.each_char.with_index do |char, index|
    digit = char.to_i

   if digit > biggest_val[:digit]
      if index == val.to_s.length - 1
         second_biggest_val = {index: index, digit: digit}
      else
         biggest_val = { index: index, digit: digit }
         second_biggest_val = {index: index, digit: -1}
      end
    elsif digit > second_biggest_val[:digit]
      second_biggest_val = { index: index, digit: digit }
    end
  end

  combined_integer = [biggest_val, second_biggest_val]
    .sort_by { |v| v[:index] }
    .map { |v| v[:digit] }
    .join
    .to_i

  total_val += combined_integer
end

puts total_val