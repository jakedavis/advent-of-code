# 3_ski_free.rb
#
# Really this is Toboggan Free, but you can't win them all.

class Tobagganer
  attr_reader :trees

  def initialize(input, steps)
    @map = Map.new(input)
    @steps = steps
    @x, @y = 0, 0
    @trees = 0
    @results = []
  end

  # Main entrypoint for this class' logic. It'll take an arbitrary number of "steps" that you want
  # completed, to support part 2. Then it iterates over those steps and keeps applying them until
  # you reach the bottom of the map. Once you do, it adds the number of trees hit to the results
  # array, resets everything, and continues on to the next step pair.
  def iterate
    @steps.each do |steps|
      step(steps.first, steps.last) while @y < @map.rows-1
      puts "Reached the bottom! Number of trees you hit was #{@trees}."
      @results << @trees
      @x, @y, @trees = 0, 0, 0
    end
  end

  # An individual step. Adds the given step interval in x/y to your Tobagganer x/y. It then looks at
  # what is at this location in the map and increments the number of trees you hit if necessary.
  def step(x=1, y=1)
    @x += x
    @y += y

    puts "Moving to #{@x}, #{@y} ..."
    found = @map.whats_here(@x, @y)

    if found == '#'
      puts "You hit a tree, ouch!"
      @trees += 1
    end
  end

  # This is for part 2 - just reduce * on the results array to get the answer.
  def multiply_results
    @results.reduce(:*)
  end

  class Map
    attr_reader :rows, :cols

    def initialize(input)
      @map = input.split("\n").map {|r| r.split('')}
      @rows = @map.length
      @cols = @map.first.length
    end

    # You need to modulo the x value to prevent an index out of bounds. Otherwise, this is pretty
    # self-explanatory - look up the value at the given x and y.
    def whats_here(x, y)
      @map[y][x % @cols]
    end
  end
end

steps = [
  [1, 1],
  [3, 1],
  [5, 1],
  [7, 1],
  [1, 2]
]

t = Tobagganer.new(File.read('./3_ski_free_input'), steps)
t.iterate
puts t.multiply_results