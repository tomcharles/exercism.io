defmodule Words do
  def count(sentence) do
    sentence
    |> String.downcase()
    |> String.split(~r/[\s,_:!&$%^@]+/, trim: true)
    |> List.foldl(%{}, &count_word/2)
  end

  defp count_word(next, acc) do
    Map.update(acc, next, 1, &(&1 + 1))
  end
end
