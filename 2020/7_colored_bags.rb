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
  def find_shiny_golds
    @ruleset.select do |k, _|
      contains_shiny_gold?(k)
    end.length
  end

  # 
  def contains_shiny_gold?(bag)
    contained = @ruleset[bag]

    # First base check - if "shiny gold" is here, just bail out.
    if contained.keys.any? {|k| k == 'shiny gold'}
      return true
    end

    # Since we're using recursion, if the given bag contains no further entries, we can safely
    # assume it does not have a shiny gold. If it's not empty, we need to dig deeper into the hash
    # to see if any of the sub-bags contain shiny gold.
    if contained.empty?
      return false
    else
      return contained.keys.any? do |c, _|
        contains_shiny_gold?(c)
      end
    end

    # I'm not sure this is strictly necessary but D E F E N S I V E  P R O G R A M M I N G
    return false
  end
end

bags = ColoredBags.new(File.read('./7_colored_bags_input'))
puts bags.find_shiny_golds
