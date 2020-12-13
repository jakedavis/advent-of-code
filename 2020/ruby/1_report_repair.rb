# 01_report.rb
#
# Day one of the AoC; report repair exercise

class ReportRepair
  def initialize(vals)
    @vals  = vals
  end

  # This method is mostly the same as before, but accepts n digits you're looking to sum. It
  # delegates most of the work to find_candidates.
  def find_values(n=2)
    @vals.each_with_index do |v, i|
      candidates = find_candidates([v], @vals[i+1..], n)

      if candidates.length == n && candidates.reduce(:+) == 2020
        return candidates.reduce(:*)
      end
    end

    return "None found"
  end

  # To make this more generic, this method now accepts n digits to use for the summation. This
  # should make it extensible for any n. This method manages three queues. Numbers are essentially
  # moving from the pending array to EITHER processed or candidates depending on if they add up to
  # less than or equal to 2020. If the length of the candidates array is less than we expect, we use
  # recursion to find the next digit. Otherwise, we check if it's exactly 2020.
  def find_candidates(candidates, pending, n)
    processed = []

    while pending.length > 0
      candidates.push(pending.pop)
    
      if candidates.length < n 
        if candidates.reduce(:+) <= 2020
          find_candidates(candidates, processed+pending, n)
        end
      elsif candidates.length == n
        if candidates.reduce(:+) == 2020
          return candidates
        end
      end

      processed.push(candidates.pop)
    end

    return []
  end
end

values = File.read('../1_input').split("\n").map(&:to_i)
repair = ReportRepair.new(values)
puts repair.find_values(2)
