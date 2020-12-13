# 5_binary_boarding.rb
#
# Day 5, Advent of Code

class BinaryBoarding
  def initialize(seats)
    @seats = seats.map {|s| Seat.new(s)}
  end

  # Easy as pie.
  def highest
    @seats.map(&:id).max
  end

  # Kind of a weird one with some assumptions baked in. I knew I could grab the index, and I
  # examinated the data a bit to figure out that the seat IDs are monotonically increasing. So the
  # idea then was to calculate the offset of the first element of the sorted list, and use that to
  # determine when the index didn't match what it shouldu have.
  def missing
    sorted = @seats.sort_by {|s| s.id}
    offset = sorted.first.id

    sorted.each_with_index do |s, i|
      if s.id != i+offset
        puts "ID #{i+offset} missing from manifest"
        offset += 1
      end
    end
  end

  class Seat
    attr_accessor :id, :row, :seat

    # Ruby is so great
    def initialize(pattern)
      @row  = pattern[0..6].gsub('F', '0').gsub('B', '1').to_i(2)
      @seat = pattern[7..9].gsub('L', '0').gsub('R', '1').to_i(2)
      @id   = @row * 8 + @seat
    end

    def to_s
      "[#{@id}] Row=#{@row} Seat=#{@seat}"
    end
  end
end

b = BinaryBoarding.new(File.read('../5_input').split("\n"))
puts b.highest
b.missing
