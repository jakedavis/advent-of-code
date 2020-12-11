# 11_seat_selector.rb
#
# Day 11, AoC

class SeatSelector 
  attr_reader :map

  def initialize(input)
    @map = input.split("\n").map {|r| r.split('')}
  end

  def run
    previous_state = nil
    cycles = 0

    while @map != previous_state
      previous_state = @map

      @map.each.with_index do |row, yidx|
        row.each.with_index do |state, xidx|
          @map[xidx][yidx] = resolve_cell(xidx, yidx, state)
        end
      end

      print_map
      cycles += 1
    end

    puts "Reached end of processing after #{cycles} cycles."
  end
  
  def print_map
    puts @map.map { |r| r.join }.join("\n")
  end

  def resolve_cell(x, y, state)
    indices = [
      [x-1, y-1],
      [x, y-1],
      [x+1, y-1],
      [x-1, y],
      [x+1, y],
      [x-1, y+1],
      [x, y+1],
      [x+1, y+1]
    ].select {|n| n.first >= 0 && n.last >= 0}

    surrounding = indices.map do |index|
      @map[index.first][index.last]
    end

    #puts "#{state} #{surrounding}"

    if state == 'L' && surrounding.none? {|s| s == '#'}
      #puts "Occupying empty seat #{x},#{y} ..."
      '#'
    elsif state == '#' && surrounding.count {|s| s == '#'} >= 4
      #puts "Emptying surrounded seat #{x}, #{y} ..."
      'L'
    else
      state
    end
  end
end

input = File.read('./11_seat_selector_input')
SeatSelector.new(input).run