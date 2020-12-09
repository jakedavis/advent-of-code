# 7_colored_bags.rb
#
# Day 7 Advent of Code

class ColoredBags
  attr_reader :ruleset

  def initialize(input)
    @ruleset = process_input(input).to_h
  end

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
end

bags = ColoredBags.new(File.read('./7_colored_bags_input'))
require 'json'; puts JSON.pretty_generate(bags.ruleset)
