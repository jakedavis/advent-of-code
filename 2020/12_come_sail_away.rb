# 12_come_sail_away.rb
#
# Day 12 Advent of Code

class Navigator
  def initialize(input, strategy: 'simple')
    @steps = input.split("\n")
    @x = 0
    @y = 0
    @facing = 1

    # Waypoint location
    @wp = [1, 10]

    # options
    @strategy = strategy
  end

  def sail
    @steps.each.with_index do |step|
      method = 'process_step_' + @strategy
      self.send(method.to_sym, step[0], step[1..].to_i)
    end

    puts "No more instructions: #{@x.abs}+#{@y.abs}=#{@x.abs+@y.abs}, facing #{direction(@facing)}"
  end

  def process_step_simple(command, arg)
    case command
    when 'N'
      @y += arg
    when 'S'
      @y -= arg
    when 'E'
      @x += arg
    when 'W'
      @x -= arg
    when 'L'
      @facing = (@facing - arg/90) % 4
    when 'R'
      @facing = (@facing + arg/90) % 4
    when 'F'
      process_step_simple(direction(@facing), arg)
    else
      puts "Invalid step: #{command}#{arg}"
    end

    #puts "[#{direction(@facing)}@#{@x},#{@y}] #{command}#{arg}"
  end

  def process_step_actual(command, arg)
    case command
    when 'N'
      @wp = [@wp.first+arg, @wp.last]
    when 'S'
      @wp = [@wp.first-arg, @wp.last]
    when 'E'
      @wp = [@wp.first, @wp.last+arg]
    when 'W'
      @wp = [@wp.first, @wp.last-arg]
    when 'L'
      modulo = arg / 90 % 4
      if modulo == 1
        @wp = [@wp.last, -@wp.first]
      elsif modulo == 2
        @wp = [-@wp.first, -@wp.last]
      elsif modulo == 3
        @wp = [-@wp.last, @wp.first]
      end
    when 'R'
      modulo = arg / 90 % 4
      if modulo == 1
        @wp = [-@wp.last, @wp.first]
      elsif modulo == 2
        @wp = [-@wp.first, -@wp.last]
      elsif modulo == 3
        @wp = [@wp.last, -@wp.first]
      end
    when 'F'
      @x += @wp.first*arg
      @y += @wp.last*arg
    else
      puts "Invalid step: #{command}#{arg}"
    end

    #puts "[#{direction(@facing)}@#{@x},#{@y}] #{command}#{arg}"
  end

  def direction(num)
    ['N', 'E', 'S', 'W'][num]
  end
end

input = File.read('./12_come_sail_away_input')
Navigator.new(input, strategy: 'actual').sail
