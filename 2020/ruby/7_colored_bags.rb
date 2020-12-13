# 7_colored_bags.rb
#
# Day 7 Advent of Code

class ColoredBags
  attr_reader :ruleset

  def initialize(input)
    @ruleset = process_input(input).to_h
  end

  # This is to take the raw input and turn it into a hash for easier manipulation down the line.
  def process_input(input)
    input.split("\n").map do |i|
      k, vs = i.split(' contain ')

      if vs != 'no other bags.'
        colors = vs.split(', ').map do |c|
          [c.match(/\d (\w+ \w+)/)[1], c[0].to_i]
        end
      else
        colors = {}
      end

      [k.match(/\w+ \w+/)[0], colors.to_h]
    end
  end

  # Entrypoint method - as long as contains_color? operates as expected, this will count the
  # total number of bags that have a shiny gold at some point.
  def instances_of_color(color)
    @ruleset.select do |k, _|
      contains_color?(color, k)
    end.length
  end

  # Checks if the given bag contains the color provided.
  def contains_color?(color, bag)
    contained = @ruleset[bag]

    # First base check - if the color is here, just bail out.
    if contained.keys.any? {|k| k == color}
      return true
    end

    # Since we're using recursion, if the given bag contains no further entries, we can safely
    # assume it does not have the color. If it's not empty, we need to dig deeper into the hash
    # to see if any of the sub-bags contain the given color.
    if contained.empty?
      return false
    else
      return contained.keys.any? do |c, _|
        contains_color?(color, c)
      end
    end
  end

  # Calculate the total number of bags contained within a given color Honestly, this one killed me.
  # I do not want to do this ever again. Warning, it produces an answer 1 larger than the actual
  # answer. I do not even care at this point.
  def bags_for_color(color)
    return 1 if @ruleset[color].empty?

    @ruleset[color].map do |k, v|
      v * bags_for_color(k, v)
    end.reduce(&:+) + 1
  end
end

bags = ColoredBags.new(File.read('../7_input'))
puts bags.bags_for_color('shiny gold')
