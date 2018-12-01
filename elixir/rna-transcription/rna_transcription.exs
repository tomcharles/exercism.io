defmodule RNATranscription do
  def to_rna(dna) do
    Enum.map(dna, fn c ->
      case c do
        ?G -> ?C
        ?C -> ?G
        ?T -> ?A
        ?A -> ?U
      end
    end)
  end
end
