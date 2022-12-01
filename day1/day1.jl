function main()
  input = split(read("input", String), "\n\n") .|> (x -> split(x, "\n"))

  # MAP phase
  data = Vector{Int64}[]
  for x ∈ input
    push!(data, begin
      l = Int64[]
      for s ∈ x
        k = tryparse(Int64, s)
        k ≠ nothing && push!(l, k)
      end
      l
    end)
  end

  # REDUCE phase
  amount_calories = sum.(data)

  # Finding maximum
  n = length(amount_calories)
  max_ind = 1
  for i = 1:n
    if amount_calories[i] > amount_calories[max_ind]
      max_ind = i
    end
  end

  println(amount_calories[max_ind])

  # Finding second maximum
  snd_max_ind = 1
  for i = 1:n
    if amount_calories[i] > amount_calories[snd_max_ind] && i ≠ max_ind
      snd_max_ind = i
    end
  end

  # Finding last (third) maximum
  trd_max_ind = 1
  for i = 1:n
    if amount_calories[i] > amount_calories[trd_max_ind] && i ≠ max_ind && i ≠ snd_max_ind
      trd_max_ind = i
    end
  end

  println(sum(amount_calories[[max_ind, snd_max_ind, trd_max_ind]]))
end

main()
