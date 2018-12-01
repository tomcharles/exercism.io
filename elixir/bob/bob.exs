defmodule Bob do
  def hey(input) do
    cond do
      empty?(input) ->
        "Fine. Be that way!"

      shouting_questions?(input) ->
        "Calm down, I know what I'm doing!"

      question?(input) ->
        "Sure."

      shouting?(input) ->
        "Whoa, chill out!"

      true ->
        "Whatever."
    end
  end

  defp empty?(input) do
    String.trim(input) == ""
  end

  defp question?(input) do
    ends_with?("?", input)
  end

  defp shouting?(input) do
    String.upcase(input) == input && String.downcase(input) != input
  end

  def shouting_questions?(input) do
    shouting?(input) && question?(input)
  end

  defp ends_with?(char, input) do
    String.last(input) == char
  end
end
