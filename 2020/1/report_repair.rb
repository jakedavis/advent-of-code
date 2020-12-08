# 01_report.rb
#
# Day one of the AoC; report repair exercise

class ReportRepair
  attr_accessor :found

  def initialize(vals)
    @vals  = vals
    @found = nil
  end

  # The main trick here is slicing out the first entry in the array with i+1, which means we don't
  # check extraneous values later on (in otherwords, the number of checks will decrease each
  # iteration, versus checking the first value in the array which has already been verified).
  def find_values
    @vals.each_with_index do |v, i|
      @found ||= find_2020(v, @vals[i+1..])
      break if found
    end
  end

  # Pretty simple - return the product of two numbers whose sum is 2020. Return nil at the end is
  # required because Array#each will return the array itself, and we're using ||= above. Probably a
  # fancier way to do this.
  def find_2020(v, arr)
    arr.each do |u|
      return v*u if v+u == 2020
    end

    return nil
  end
end

values = File.read('./report_repair_input').split("\n").map(&:to_i)
repair = ReportRepair.new(values)
repair.find_values
puts repair.found
