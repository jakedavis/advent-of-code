# 3_ski_free.rb
#
# Really this is Toboggan Free, but you can't win them all.

class Tobagganer
  attr_reader :trees

  def initialize(input)
    @map = Map.new(input)
    @x, @y = 0, 0
    @trees = 0
  end

  def iterate
    step while @y < @map.rows-1
    puts "Reached the bottom! Number of trees you hit was #{@trees}."
  end

  def step(x=3, y=1)
    @x += x
    @y += y

    puts "Moving to #{@x}, #{@y} ..."
    found = @map.whats_here(@x, @y)

    if found == '#'
      puts "You hit a tree, ouch!"
      @trees += 1
    end
  end

  class Map
    attr_reader :rows, :cols

    def initialize(input)
      @map = input.split("\n").map {|r| r.split('')}
      @rows = @map.length
      @cols = @map.first.length
    end

    def whats_here(x, y)
      @map[y][x % @cols]
    end
  end
end

t = Tobagganer.new(File.read('./3_ski_free_input'))
t.iterate