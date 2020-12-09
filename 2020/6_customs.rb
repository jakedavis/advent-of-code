# 6_customs_input.rb
#
# Day 6 AoC

class Customs
  def initialize(input)
    @groups = input.split("\n\n").map {|g| Group.new(g)}
  end

  def score
    @groups.map(&:score_unan).reduce(:+)
  end

  class Group
    def initialize(forms)
      @forms = forms.split("\n")
    end

    def score
      @forms.join.chars.uniq.length
    end

    def score_unan
      @forms.map(&:chars).reduce(&:intersection).length
    end
  end

  class Form
    attr_reader :answers

    def initialize(answers)
      @answers = answers
    end

    def score
      @answers.length
    end
  end
end

c = Customs.new(File.read('./6_customs_input'))
puts c.score
