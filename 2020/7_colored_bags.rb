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

  # Entrypoint method - as long as contains_shiny_gold? operates as expected, this will count the
  # total number of bags that have a shiny gold at some point.
  def instances_of_color(color)
    @ruleset.select do |k, _|
      contains_color?(color, k)
    end.length
  end

  # Another entrypoint for summing the total number of bags required for a given color.
  def total_bags_for_color(color)
    bags_for_color(color)
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

  # 
  def bags_for_color(color, layer=1)
    bags = @ruleset[color].values.reduce(&:+) || 0
    subbags = @ruleset[color].keys.map do |b|
      bags_for_color(b, layer+1)
    end.reduce(&:+) || 0

    puts "#{"  " * layer}[#{color}] bags=#{bags} subbags=#{subbags}"

    if color == 'shiny gold'
      require 'pry'; binding.pry
    end

    return bags + subbags
  end
end

bags = ColoredBags.new(File.read('./7_colored_bags_input'))
puts bags.total_bags_for_color('shiny gold')
