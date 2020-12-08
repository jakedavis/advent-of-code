# 01_report.rb
#
# Day one of the AoC; report repair exercise

class ReportRepair
  attr_accessor :found

  def initialize(vals)
    @vals  = vals
    @found = nil, nil
  end

  def find_values
    @vals.each do |v|
      @found ||= find_2020(arr)
    end
  end

  def find_2020(arr)
    first = arr.first

    arr.drop(1).each do |v|
      return first, v if first+v == 2020
    end
  end
end

values = [
  1721,
  979,
  366,
  299,
  675,
  1456
]

repair = ReportRepair.new(values)
repair.find_values
puts repair.found
