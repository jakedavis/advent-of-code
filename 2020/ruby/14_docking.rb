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

        # Convert everything to 36 bits binary
        add_binary = ("%36b" % addr.to_i).gsub(' ', '0')
        val_binary = ("%36b" % value.to_i).gsub(' ', '0')

        case @version
        when 1
          version_1_bootup(add_binary, val_binary, mask)
        when 2
          version_2_bootup(add_binary, val_binary, mask)
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
    @memory[addr] = mask_value(mask, val.chars)
  end

  def version_2_bootup(addr, value, mask)
    mask_addr(mask, addr).each do |a|
      val = value.to_i(base=2)
      @memory[a] = val
    end
  end

  def mask_addr(mask, addr)
    result = mask.chars.map.with_index do |c, idx|
      c == '0' ? addr[idx] : c
    end

    # We need as many addresses as there are 2 to the number of X
    occurrences = result.count {|n| n == 'X'}
    addrs = (1..2**occurrences).map {|n| {n=>[]}}.reduce(&:merge)

    occur = 0
    result.each do |r|
      if r == '0' || r == '1'
        addrs.each {|_, v| v << r}
      elsif r == 'X'
        digit = '0'
        iters = 2**occur
        count = 0
        addrs.each do |_, v|
          v << digit
          count += 1
          if count == iters
            count = 0
            digit = digit == '0' ? '1' : '0'
          end
        end

        occur += 1
      else
        puts "FATAL: Bad character #{r}"
        exit
      end
    end

    addrs.values.map(&:join).map {|a| a.to_i(base=2)}
  end

  def mask_value(mask, value)
    mask.chars.map.with_index do |c, idx|
      c == 'X' ? value[idx] : c
    end.join.to_i(base=2)
  end
end

input = File.read('../14_input').split("\n")
puts Docker.new(input, version: 2).bootup
