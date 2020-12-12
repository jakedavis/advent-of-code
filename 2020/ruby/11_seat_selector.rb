# 11_seat_selector.rb
#
# Day 11, AoC

class SeatSelector
  attr_reader :map

  def initialize(input, occupied_threshold: 4, strategy: 'adjacent')
    @current = input.split("\n").map {|r| r.split('')}
    @previous = []
    @column_size = @current.length
    @row_size = @current.first.length

    @cycles = 0
    @changes = 0

    # Options
    @occupied_threshold = occupied_threshold
    @strategy = strategy
  end

  # This is the entrypoint. As long as you fed in a valid map to the constructor, this will "just
  # work". The Marshal::load line is something I did not know about Ruby: Object#dup will create a
  # shallow copy of an object, so you still get any side effects from changing the original. You
  # have to use this code to actually create a separate copy in memory. There might be a way to do
  # this without that (you can "commit" a row once you're processing two rows ahead), but I don't
  # know, this seems fine enough.
  def run
    while @current - @previous != []
      @previous = Marshal::load(Marshal.dump(@current))

      @previous.each.with_index do |row, yidx|
        row.each.with_index do |state, xidx|
          @current[yidx][xidx] = resolve_cell(yidx, xidx, state)
        end
      end

      puts "Cycle #{@cycles} with #{@changes} changes"
      @cycles += 1
      @changes = 0
    end

    puts "Reached end of processing after #{@cycles} cycles, #{occupied_seats} seats are occupied."
  end

  def occupied_seats
    @current.map do |row|
      row.count {|c| c == '#'}
    end.reduce(&:+)
  end

  def print_map(state)
    puts state.map { |r| r.join }.join("\n")
  end

  # This is a recursive lookup for the first cell that isn't empty. We need to bail out if the cell
  # we want to look at is outside of the actual size of the map. Ruby will let you use negative
  # numbers but we actually don't want that here. Returning nil is fine because the code that calls
  # this will specifically check for L or #.
  def look(y, x, ydelta, xdelta)
    return nil if y.negative? || x.negative?
    return nil if y > @column_size || x > @row_size

    state = @previous.fetch(y, [])[x]
    if state == 'L' || state == '#'
      return state
    else
      return look(y+ydelta, x+xdelta, ydelta, xdelta)
    end
  end

  # This is the main logic for processing one cell on a given run. It supports two strategies, which
  # correspond to the two parts of the problem. The main thing is checking the "surrounding" cells
  # based on the strategy, and then running that plus the current cell's value through a very small
  # if tree to determine what its new state should be.
  def resolve_cell(y, x, state)
    surrounding = []

    if @strategy == 'los'
      surrounding = [
        look(y-1, x-1, -1, -1),
        look(y, x-1, 0, -1),
        look(y+1, x-1, 1, -1),
        look(y-1, x, -1, 0),
        look(y+1, x, 1, 0),
        look(y-1, x+1, -1, 1),
        look(y, x+1, 0, 1),
        look(y+1, x+1, 1, 1)
      ]
    elsif @strategy == 'adjacent'
      indices = [
        [y-1, x-1],
        [y, x-1],
        [y+1, x-1],
        [y-1, x],
        [y+1, x],
        [y-1, x+1],
        [y, x+1],
        [y+1, x+1]
      ].select {|n| n.first >= 0 && n.last >= 0}

      surrounding = indices.map do |index|
        @previous.fetch(index.first, [])[index.last]
      end
    else
      puts "Invalid strategy, bailing."
      exit
    end

    if state == 'L' && surrounding.none? {|s| s == '#'}
      @changes += 1
      '#'
    elsif state == '#' && surrounding.count {|s| s == '#'} >= @occupied_threshold
      @changes += 1
      'L'
    else
      state
    end
  end
end

input = File.read('./11_seat_selector_input')
SeatSelector.new(input, occupied_threshold: 5, strategy: 'los').run
