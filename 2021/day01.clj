(require '[clojure.string])

; parse input file to integer vectors
(defn- get-input  []
  (let [
        ; read file
        file (slurp "day01.txt") 

        ; split lines
        lines (clojure.string/split file #"\r|\n|\r\n") 

        ; remove empty lines
        non-empty-lines (filter #(not= "" %) lines)] 

    ; parse to integers
    (for [s non-empty-lines] (Integer/parseInt s))))

; count the number of height increases
(defn how-many-increases [heights]
  ; reduction with pair of number of increases and previous value
  (get (reduce (fn [[sum prev] curr]
                 [(if (and (not= prev nil)
                           (> curr prev))
                    (+ sum 1)
                    sum)
                  curr])
               [0 nil] heights) 0))

; add all subvectors of length n and get resulting vector
(defn- add-tuples [input n]
  (let [in-count (count input)
        out-count (- in-count n -1)]
    (for [i (range out-count)]
      (apply + (subvec input i (+ i n))))))

(let [input (vec (get-input))]
  (print "Part 1:" (str (how-many-increases input)))
  (print "Part 2:" (str (how-many-increases (add-tuples input 3)))))
