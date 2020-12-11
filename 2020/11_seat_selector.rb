# 11_seat_selector.rb
#
# Day 11, AoC

class SeatSelector
  attr_reader :map

  def initialize(input, occupied_threshold: 4, strategy: 'adjacent')
    @current = input.split("\n").map {|r| r.split('')}
    @previous = []
    @cycles = 0
    @changes = 0

    # Options
    @occupied_threshold = occupied_threshold
    @strategy = strategy
  end

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

  def look(y, x, ydelta, xdelta)
    # This is an out-of-bounds that will wrap around - we want to avoid this. It's a Ruby quirk!
    if y.negative? || x.negative?
      return @previous[y-ydelta][x-xdelta]
    end

    state = @previous.fetch(y, [])[x]
    if state == 'L' || state == '#'
      return state
    elsif state == '.'
      look(y+ydelta, x+xdelta, ydelta, xdelta)
    else
      # This is an out-of-bounds I think
      return '.'
    end
  end

  def resolve_cell(y, x, state)
    surrounding = []

    if @strategy == 'line of sight'
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
      end.select {|s| !s.nil?}
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
SeatSelector.new(input, occupied_threshold: 5, strategy: 'line of sight').run
