# 10_jolt_adapter.rb
#
# Day 10 AoC

class JoltAdapter
  attr_reader :combos

  def initialize(adapters)
    @adapters = adapters.sort
    @combos   = tree_of_possibilities
  end

  def differences
    base = [0] + @adapters
    device = @adapters + [@adapters.max+3]

    diffs = base.zip(device).map do |f, l|
      l - f
    end

    diffs.count(1) * diffs.count(3)
  end

  def tree_of_possibilities
    @adapters.map do |a|
      { a => possible_for(a) }
    end.reduce(&:merge)
  end

  def possible_for(k)
    possibles = [ k+1, k+2, k+3 ]
    possibles.select do |p|
      @adapters.include?(p)
    end
  end

  def derive_tree(k, vs)
    if vs.empty?
      [k+3]
    else
      vs.map do |v|
        derive_tree(v, @combos[v])
      end
    end
  end
end

input = File.read('./10_jolt_adapter_input').split("\n").map(&:to_i)
j = JoltAdapter.new(input)

puts j.differences
puts j.derive_tree(1, j.combos[1])
