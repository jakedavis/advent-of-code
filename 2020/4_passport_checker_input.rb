# 4_passport_checker.rb
#
# Day 4, AoC

class PassportScanner
  attr_reader :valids

  def initialize(input)
    @valids = nil
    @passports = process_input(input)
  end

  def process_input(input)
    passports = input.split("\n\n")
    passports.map do |p|
      Passport.new(p)
    end
  end

  def verify
    @valids ||= @passports.select do |passport|
      valid?(passport)
    end.length
  end

  def valid?(p)
    byr_valid?(p.byr) && iyr_valid?(p.iyr) && eyr_valid?(p.eyr) && hgt_valid?(p.hgt) &&
      ecl_valid?(p.ecl) && hcl_valid?(p.hcl) && pid_valid?(p.pid) && cid_valid?(p.cid)
  end

  def byr_valid?(byr)
    byr >= 1920 && byr <= 2002
  end

  def iyr_valid?(iyr)
    iyr >= 2010 && iyr <= 2020
  end

  def eyr_valid?(eyr)
    eyr >= 2020 && eyr <= 2030
  end

  def hgt_valid?(hgt)
    value = hgt[0..-2].to_i
    unit = hgt[-2..]

    case unit
    when 'cm'
      value >= 150 && value <= 193
    when 'in'
      value >= 59 && value <= 76
    else
      false
    end
  end

  def ecl_valid?(ecl)
    [
      'amb',
      'blu',
      'brn',
      'gry',
      'grn',
      'hzl',
      'oth'
    ].any? {|c| c == ecl}
  end

  def hcl_valid?(hcl)
    hcl.match(/#[a-f\d]{6}/)
  end

  def pid_valid?(pid)
    pid.length == 9 && pid.match(/\d{9}/)
  end

  def cid_valid?(cid)
    true
  end

  class Passport
    attr_reader :byr, :iyr, :eyr, :ecl, :hcl, :hgt, :pid, :cid

    def initialize(input)
      params = decipher_input(input)

      @byr = params['byr'].to_i || 0
      @iyr = params['iyr'].to_i || 0
      @eyr = params['eyr'].to_i || 0
      @ecl = params['ecl'] || ''
      @hcl = params['hcl'] || ''
      @hgt = params['hgt'] || ''
      @pid = params['pid'] || ''
      @cid = params['cid']
    end

    def decipher_input(input)
      items = input.split(' ')
      items.map do |i|
        i.split(':')
      end.to_h
    end

    def to_s
      "[#{pid}] #{byr} #{iyr}-#{eyr}, eye=#{ecl}, hair=#{hcl}, hgt=#{hgt}"
    end
  end
end

p = PassportScanner.new(File.read('./4_passport_checker_input'))
puts p.verify
