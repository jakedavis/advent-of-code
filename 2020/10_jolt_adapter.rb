# 10_jolt_adapter.rb
#
# Day 10 AoC

class JoltAdapter
  attr_reader :combos

  def initialize(adapters)
    @adapters = adapters.sort
  end

  def differences
    base = [0] + @adapters
    device = @adapters + [@adapters.max+3]

    diffs = base.zip(device).map do |f, l|
      l - f
    end
  end

  def absolute_difference
    diffs = differences
    diffs.count(1) * diffs.count(3)
  end

  def number_of_paths
    ([0] + @adapters + [@adapters.last+3]).chunk_while do |x, y|
      x + 1 == y
    end.map do |chunk|
      tribonacci(chunk.length+1)
    end.reduce(&:*)
  end

  def tribonacci(of_length)
    seq = [0, 1, 1]
    
    while seq.length < of_length
      seq << (seq[-3] + seq[-2] + seq[-1])
    end
    
    seq.last
  end
end

input = File.read('./10_jolt_adapter_input').split("\n").map(&:to_i)
j = JoltAdapter.new(input)

puts j.absolute_difference
puts j.number_of_paths
