# 13_bus.rb
#
# Day 13, Advent of Code

class BusRider
  def initialize(timestamp, lines)
    @timestamp = timestamp.to_i
    @lines = lines.split(',').map(&:to_i)
  end

  def next_bus
    choice = @lines.reject {|l| l == 0}.map do |line|
      { line => (@timestamp.to_f/line).ceil*line }
    end.reduce(&:merge).sort_by {|_, v| v}.first

    choice[0] * (choice[1] - @timestamp)
  end

  def contest
    iterations = 1
    sequential = false
    timestamp  = 0

    lines = @lines.map.with_index do |n, offset|
      { n => offset }
    end.reduce(&:merge).select {|k,_| k != 0}
    multiplier = lines.max
    offset = multiplier[1]

    until sequential
      timestamp = multiplier[0] * iterations

      sequential = lines.map do |l, idx|
        (timestamp + (idx - offset)) % l
      end.all? {|l| l == 0}

      if iterations % 1000000 == 0
        puts iterations
      end
      iterations += 1
    end

    timestamp
  end
end

timestamp, lines = File.read('./13_input').split("\n")
#puts BusRider.new(timestamp, lines).next_bus
puts BusRider.new(timestamp, lines).contest
