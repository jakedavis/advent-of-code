# 19_message_matcher.rb
#
# Day 19, Advent of Code

class MessageMatcher
  def initialize(rules, messages)
    @rules = process_rules(rules)
    @messages = messages
    puts @rules
  end

  def process_rules(rules)
    rules.map do |rule|
      number, remainder = rule.split(': ')
      rule1, rule2 = remainder.split(' | ')

      if rule1[1] == 'a' || rule1[1] == 'b'
        { number.to_i => rule1[1] }
      else
        subrules = []
        subrules << rule1.split(' ').map(&:to_i)
        subrules << rule2.split(' ').map(&:to_i) if rule2

        { number.to_i => subrules }
      end
    end.reduce(&:merge)
  end

  # 115: [
  #   [20, 70],
  #   [30, 93]
  # ]
  def resolve_rule(num)
    this_rule = @rules[num]

    if this_rule == "a" || this_rule == "b"
      return this_rule
    else
      return this_rule.map do |subrule|
        subrule.map {|r| resolve_rule(r)}
      end
    end
  end

  # 0: a {1} b
  # 1: aaab | aaba | bbab | bbba | abaa | abbb | baaa | babb
  # 2: aa | bb
  # 3: ab | ba
  # 4: a
  # 5: b
  # a aaab b
  # a aaba b
  # a bbab b
  # a bbba b
  # a abaa b
  # a abbb b
  # a baaa b
  # a babb b
  def verify_message(message)

  end
end

rules, messages = File.read('../19_debug').split("\n\n")
m = MessageMatcher.new(rules.split("\n"), messages.split("\n"))
puts "#{m.resolve_rule(0)[0]}"
