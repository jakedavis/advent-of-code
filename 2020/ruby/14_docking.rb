# 14_docking.rb
#
# Day 14 Advent of Code

class Docker
  def initialize(sequence, version: 1)
    @sequence = sequence
    @memory = {}
    @version = version
  end

  def bootup
    mask = nil

    @sequence.each do |s|
      cmd, value = s.split(' = ')

      if cmd == 'mask'
        mask = value
      elsif cmd[0..2] == 'mem'
        addr = cmd.match(/(\d+)/)[1]

        case @version
        when 1
          version_1_bootup(cmd, value, mask)

        when 2
          version_2_bootup(cmd, value, mask)
        else
          puts "FATAL: Incompatible version #{@version}"
          exit
        end
      else
        puts "Fatal bootup command #{cmd} = #{value}"
        exit
      end
    end

    @memory.values.reduce(&:+)
  end

  def version_1_bootup(addr, value, mask)
    val = ("%36b" % value.to_i).gsub(' ', '0')
    @memory[addr] = mask_value(mask, val.chars)
  end

  def version_2_bootup(addr, value, mask)
    addrs = mask_addr(mask, addr)
    addrs.each do |a|
      @memory[a] = value
    end
  end

  def mask_addr(mask, addr)
    value = mask.chars.map.with_index do |c, idx|
      c == '0' ? addr[idx] : c
    end

    addrs = []
    # recur on value, check for X, gen addrs

    # Calculate number of possibilities based on X
    # Construct all possibles at once
    #   At each X, half 0/half 1?
    #   Need to count X's or something
    #   At first X, one 0 first
    #   At second X, two 0 first
    #   At third X, four? 0 first
    #   etc, so 2^X zeroes per X encountered
    # 00000011001
    # 01000011001
    # 00001011001
    # 01001011001
    #
    # Take numbers until X

    addrs
  end

  def mask_value(mask, value)
    mask.chars.map.with_index do |c, idx|
      c == 'X' ? value[idx] : c
    end.join.to_i(base=2)
  end
end

input = File.read('../14_input').split("\n")
puts Docker.new(input, version: 2).bootup
