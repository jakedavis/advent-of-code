defmodule ReportRepair do

  def repair do
    {:ok, f} = File.read("../1_input")
    f
  end

  def printsomething(some) do
    IO.puts "Hi " <> some
  end
end

ReportRepair.printsomething("lol")
ReportRepair.repair