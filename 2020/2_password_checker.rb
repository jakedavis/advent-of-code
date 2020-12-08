# 2_password_checker.rb
#
# Advent of Code Day 2!

class PasswordChecker
  attr_reader :valid 

  def initialize(list)
    @list = list
    @valid = 0
  end

  def resolve    
    @list.each do |l|
      policy, pw = l.split(': ')

      p = Policy.new(policy)
      @valid += 1 if p.valid_by_position?(pw)
    end
  end

  class Policy
    def initialize(raw='')
      range, sym = raw.split(' ')
      lower, upper  = range.split('-').map(&:to_i)

      @lower = lower
      @upper = upper
      @sym   = sym
    end

    def valid_by_count?(password)
      len = password.split('').filter {|p| p == @sym}.length

      len >= @lower && len <= @upper
    end

    def valid_by_position?(password)
      check_lower = password.split('')[@lower-1]
      check_upper = password.split('')[@upper-1]

      return false if check_lower == check_upper
      
      @sym == check_lower || @sym == check_upper
    end
  end
end

pc = PasswordChecker.new(File.read('./2_password_checker_input').split("\n"))
pc.resolve

puts pc.valid