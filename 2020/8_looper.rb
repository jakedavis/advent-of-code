# 8_looper.rb
#
# Day 8 Advent of Code

require 'set'

class Looper 
  def initialize(set, debug_inst: -1)
    @instruction_set = set.split("\n")
    @acc = 0
    @inst = 0
    @visited = Set.new
    @last = @instruction_set.length-1
    
    # Debugging purposes
    @debug_inst = debug_inst
  end

  # Per the instructions, nop simply increments the instruction counter by 1.
  def nop(_)
    @inst += 1
  end

  # Jumps val positions relative to current position. No effect on accumulator. There is special
  # code for debugging a jmp that should be a nop as well, which by default will not run.
  def jmp(val)
    if @inst == @debug_inst
      puts "Swapping jmp for nop on instruction #{@inst} ..."
      nop(@inst)
      return
    end
    
    @inst += val
  end

  # Adds the val to the accumulator and increments the instruction counter by 1.
  def acc(val)
    @acc += val
    @inst += 1
  end

  # Run forever unless we hit exit conditions. Those are:
  #   1. An instruction has already been visited (not very practical but it's the exercise)
  #   2. The current instruction is greater than the last instruction in the set
  #
  # Otherwise, uses Object#send to convert the plaintext instruction to a Ruby method call, and adds
  # the current instruction to the visited set.
  def execute
    while true
      if @visited.include?(@inst)
        puts "Already visited #{@inst}. Accumulator value is #{@acc}."
        return @inst
      end

      if @inst > @last
        puts "Reached end of boot sequence, terminating!"
        puts "Accumulator value is #{@acc}."
        return @inst
      end

      cmd, arg = @instruction_set[@inst].split(' ')
      
      @visited << @inst
      self.send(cmd.to_sym, arg.to_i)
    end
  end
end

# This is for finding the instruction that needs to become a nop
require 'json'; 
debug = JSON.parse(File.read('./8_looper_debug_list'))
debug.reject do |d|
  l = Looper.new(File.read('./8_looper_input'), debug_inst: d)
  l.execute
end

# This is the 'normal' flow
l = Looper.new(File.read('./8_looper_input'))
puts l.execute
