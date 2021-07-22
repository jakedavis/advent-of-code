# 20_satellite.rb
#
# Day 20, AoC

class SatelliteMapper
  def initialize(tiles)
    @borders = get_borders(tiles)
  end

  def get_borders(tiles)
    tiles.map do |tile|
      id = tile.split("\n").first.match(/(\d+)/)[1].to_i
      rows = tile.split("\n").drop(1).map {|r| r.split('')}
      columns = rows.transpose

      {
        id: id,
        top: rows[0],
        bottom: rows[-1],
        left: columns[0],
        right: columns[-1]
      }
    end
  end

  def find_corners
    tops    = @borders.map { |b| b[:top] }
    bottoms = @borders.map { |b| b[:bottom] }
    lefts   = @borders.map { |b| b[:left] }
    rights  = @borders.map { |b| b[:right] }
    product = 1

    @borders.map do |element|
      puts "#{element[:top]}"
      exit
      if !bottoms.include?(element[:top])
        if !lefts.include?(element[:right]) || !rights.include?(element[:left])
          #puts "Found top corner #{element[:id]}!"
          product *= element[:id]
        end
      elsif !tops.include?(element[:bottom])
        if !lefts.include?(element[:right]) || !rights.include?(element[:left])
          #puts "Found bottom corner #{element[:id]}!"
          product *= element[:id]
        end
      end
    end

    product
  end
end

input = File.read('../20_input').split("\n\n")
puts SatelliteMapper.new(input).find_corners
