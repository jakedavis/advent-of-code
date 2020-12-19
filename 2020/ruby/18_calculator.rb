# 18_calculator.rb
#
# Day 18, Advent of Code

class Calculator
  def initialize(precedence_method: 'equal')
    @precedence_method = precedence_method
  end

  # Had a lot of fun with this. This is a shunting yard algorithm for parsing infix notation into
  # reverse polish notation. I did it this way partly for fun and partly because I did not want to
  # deal with parenthesis, which RPN allows you to ignore. You need to convert infix to RPN though,
  # so this accomplishes that. Thanks Djikstra!
  def parse(math)
    # These are QUEUES, so we only want to push and pop to them.
    operators = []
    output    = []

    # This is a little gross but this will split parenthesis from the numbers directly next to
    # them, so it's necessary for the algorithm.
    elements = math.split(' ').map do |element|
      numeric = element.match(/[-]?\d+/)

      if numeric
        lefts   = element.chars.take_while {|e| e == '('}
        rights  = element.chars.reverse.take_while {|e| e == ')'}
        number  = numeric[0].to_i
        [lefts, number, rights]
      else
        element
      end
    end.flatten

    # Main shunting yard algo
    elements.each do |element|
      if element.is_a?(Numeric)
        output.push(element)
      elsif operator?(element)
        # We need to determine if there is greater precedence somewhere to determine what to do
        # down below. In the equal case, everything has equal precedence, so this clause is always
        # true.
        if @precedence_method == 'equal'
          precedence = true
        elsif @precedence_method == 'addition'
          # This isn't the best, and could get more complicated for more operators, but this
          # problem only required + and *, so it's okay. In this scenario, we were supposed to give
          # addition higher precedence than *.
          precedence = operators.last == '+' && element == '*'
        else
          puts "FATAL: Unsure how to resolve precedence ambiguity, have to bail ..."
          exit
        end

        # If there is an operator in the queue, and it isn't a (, and it has greater precedence
        # than the current element, push the operator onto the output queue.
        while !operators.empty? && precedence && operators.last != '('
          output.push(operators.pop)
        end

        # Finish by pushing the operator to the operator queuue
        operators.push(element)
      elsif element == '('
        operators.push(element)
      elsif element == ')'
        while operators.last != '('
          output.push(operators.pop)
        end

        # This should discard (
        operators.pop
      else
        puts "FATAL: Unrecognized character #{element}, have to bail ..."
        exit
      end
    end

    # Last step is to push any remaining operators onto the output stack.
    while !operators.empty?
      output.push(operators.pop)
    end

    output
  end

  def calculate(rpn)
    temp = []
    rpn.each do |elem|
      if operator?(elem)
        temp.push(temp.pop.send(elem.to_sym, temp.pop))
      else
        temp.push(elem)
      end
    end

    temp.last
  end

  def solve(line)
    calculate(parse(line))
  end

  def operator?(elem)
    ['+', '-', '*', '/', '^'].any? {|o| o == elem}
  end
end

input = File.read('../18_input').split("\n")
c = Calculator.new(precedence_method: 'addition')

result = input.map do |l|
  c.solve(l)
end.reduce(&:+)

puts result
