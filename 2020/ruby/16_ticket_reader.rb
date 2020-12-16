# 16_ticket_reader.rb
#
# Day 16 Advent of Code

class TicketScanner
  attr_reader :tickets

  def initialize(fields, tickets, myticket: [])
    @fields  = fields.map {|f| process_field(f)}.to_h
    @tickets = tickets.map {|t| process_ticket(t)}
    @mine    = process_ticket(myticket)

    @valid, @invalid = scan(@tickets)
  end

  def process_ticket(ticket)
    ticket.split(',').map(&:to_i)
  end

  def process_field(field)
    m = /([\w\s]+): (\d+)-(\d+) or (\d+)-(\d+)/.match(field)
    [ m[1], [ m[2].to_i..m[3].to_i, m[4].to_i..m[5].to_i ] ]
  end

  # This is a specific variation of scanning tickets. We want the sum of the invalid values within
  # any ticket. Only invalid tickets have invalid values. We need to select those values from each
  # invalid ticket and then sum it all up.
  def invalid_sum
    @invalid.map do |t|
      t.reject do |v|
        value_valid?(v)
      end
    end.flatten.reduce(&:+)
  end

  # A value is valid if it is a member of any field.
  def value_valid?(v)
    @fields.values.reduce(&:+).any? do |f|
      f.include?(v)
    end
  end

  # A ticket is valid if every value within the ticket is valid.
  def ticket_valid?(t)
    t.all? do |v|
      value_valid?(v)
    end
  end

  # Separate out tickets into valid or invalid based on whether they are valid.
  def scan(tickets)
    tickets.partition do |t|
      ticket_valid?(t)
    end
  end
end

fields, mine, tickets = File.read('../16_input').split("\n\n")
fields  = fields.split("\n")
tickets = tickets.split("\n").drop(1)
mine    = mine.split("\n").last

t = TicketScanner.new(fields, tickets, myticket: mine)
puts t.invalid_sum
