# 17_conway_cube.rb
#
# Day 17 AoC

class ConwayCube
  attr_reader :cube

  def initialize(initial)
    @cycle = 0
    @size  = initial.split("\n").length
    @cube  = [
      [('.' * @size).split('')] * @size,
      initial.split("\n").map {|r| r.split('')},
      [('.' * @size).split('')] * @size
    ]
    @next  = generate_cube(@size+2)
  end

  def neighbors(x, y, z)
    [-1,0,1].product([-1,0,1], [-1,0,1]).map do |offset|
      [
        offset[0] + x,
        offset[1] + y,
        offset[2] + z
      ]
    end - [[x, y, z]]
  end

  def solve(cycles)
    while @cycle < cycles
      print_cube # TODO remove
      @cube.each.with_index do |zs, z|
        zs.each.with_index do |ys, y|
          ys.each.with_index do |xs, x|
            mystate = @cube[z][y][x]
            active_neighbors = neighbors(x, y, z).select do |n|
              nx, ny, nz = n[0], n[1], n[2]
              if nx < 0 || ny < 0
                false
              else
                @cube.fetch(nz, [['.']]).fetch(ny, ['.']).fetch(nx, '.') == '#'
              end
            end

            puts "[#{x},#{y},#{z}] #{mystate} #{active_neighbors.length} #{active_neighbors}"

            nextstate = mystate
            if mystate == '#'
              if active_neighbors.length < 2 || active_neighbors.length > 3
                puts "            transition to . state"
                nexstate = '.'
              end
            else
              if active_neighbors.length == 3
                puts "            transition to # state"
                nextstate = '#'
              end
            end

            @next[z+1][y+1][x+1] = nextstate
          end
        end
      end

      # State transition
      @cube   = @next
      @cycle += 1
      @next   = generate_cube(@cycle+2 + @size)
    end
  end

  def generate_cube(size)
    [[[('.' * size).split('')] * size] * size][0]
  end

  def calculate_actives
    @cube.map do |z|
      z.map do |y|
        y.count do |x|
          x == '#'
        end
      end.reduce(&:+)
    end.reduce(&:+)
  end

  def print_cube
    @cube.each do |zs|
      zs.each do |ys|
        puts ys.join
      end
      puts "\n"
    end
  end
end

input = File.read('../17_debug')
c = ConwayCube.new(input)
c.solve(1)
puts "Had #{c.calculate_actives} active cubes at end of cycle 1"
c.print_cube
