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

    @positions = determine_positions
  end

  # Part 2 asks us to multiply the values on our ticket that correspond to the position of any
  # field that contains "departure" in its name.
  def departure_product
    @positions.select do |k,v|
      v.include?('departure')
    end.keys.map do |k|
      @mine[k]
    end.reduce(&:*)
  end

  # This will calculate what field is at what position. The way it does this is by process of
  # elimination; essentially, it looks for the index that has the least number of possible values.
  # It assumes this list is of length 1 (not great) and then assigns that list of fields to be
  # "subtracted" by the next index's list of possible fields. This works if there is only one
  # additional possibility per index.
  def determine_positions
    subtract_fields = []
    process_positions.sort_by {|k, v| v.length}.map do |position, fields|
      assigned_field = (fields - subtract_fields)[0]
      subtract_fields = fields

      [position, assigned_field]
    end.to_h
  end

  # This method is WILD. First we need to transpose, as this will organize the ticket values by
  # their index position. This is good. We want to process by position. We then want to select
  # fields for which EVERY value at the given position is within ANY range.
  def process_positions
    @valid.transpose.map.with_index do |t_by_idx, idx|
      fs = @fields.select do |k, f|
        t_by_idx.all? do |v|
          f.any? do |r|
            r.include?(v)
          end
        end
      end.keys

      [idx, fs]
    end.to_h
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
puts t.departure_product
