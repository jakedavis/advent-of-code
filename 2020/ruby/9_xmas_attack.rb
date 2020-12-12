# 9_xmas_attack.rb
#
# Day 9, Advent of Code ...

class XmasAttack
  def initialize(stream)
    @lrange = 0
    @hrange = 24
    @stream = stream
    @queue  = stream[@lrange..@hrange]
    
    @invalid = 0
  end

  # Kind of confusing to follow. Array#index returns the first object that matches the criteria. We
  # want the first number that cannot be summed by two numbers in the queue (last 25 numbers).
  # Array#index returns nil if no instance is found. So in otherwords, if we find no instances of a
  # summable number in the queue, it'll return nil. We can ! that to flag to the outer loop when we
  # need to exit. At that point, the incremented high range will be the correct index of the number
  # that satisfies the conditions (since the outerloop starts at hrange+1). We then simply grab the
  # number at that index in the stream.
  def attack
    @stream[(@hrange+1)..].index do |s|
      qidx = !@queue.index do |opt|
        @queue.include?((s-opt).abs)
      end

      @lrange += 1
      @hrange += 1
      @queue = @stream[@lrange..@hrange]

      qidx
    end

    @invalid = @stream[@hrange]
  end

  # Finds the contiguous sum of numbers that equal the invalid number. It also sums the resulting
  # min and max from the array of numbers that compose the sum.
  def contiguous_sum
    sum, @lrange = 0, 0
    
    while @stream[@lrange] < @invalid
      taken = @stream[@lrange..].take_while do |i|
        sum += i
        sum < @invalid
      end
      # Ugly, necessary because Ruby won't take the last element in the above loop if it exactly
      # equals the invalid number. I tried this with <= and it didn't work. I am still somewhat
      # confused by this.
      taken << @stream[@lrange+taken.length]

      return taken.min + taken.max if sum == @invalid

      @lrange += 1
      sum = 0
    end

    return nil
  end
end

x = XmasAttack.new(File.read('./9_xmas_attack_input').split("\n").map(&:to_i))
puts "Invalid number is #{x.attack}"
puts "Min/max sum is #{x.contiguous_sum}"