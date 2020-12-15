# 15_number_game.rb
#
# Day 15, Advent of Code

class NumberGame
  def initialize(intro, target=2020)
    @numbers = {}
    @target  = target
    @turn    = intro.length+1
    @spoken  = intro.last
    @intro   = intro

    intro.each_with_index do |n, idx|
      @numbers[n] = [idx+1]
    end
  end

  def play
    puts "Starting game at turn #{@turn} with intro #{@intro.join(', ')}! Our goal is #{@target} turns."

    while @turn != @target+1
      if @numbers[@spoken].length == 1
        @spoken = 0
      else
        @spoken = @numbers[@spoken][-1] - @numbers[@spoken][-2]
      end

      @numbers[@spoken] ||= []
      @numbers[@spoken] << @turn
      @turn += 1
    end

    @spoken
  end
end

input = [ 0, 6, 1, 7, 2, 19, 20 ]
puts NumberGame.new(input, 30000000).play
