(require '[clojure.string])

(defn- get-input  []
  (let [file (slurp "day01.txt")
        lines (clojure.string/split file #"\r|\n|\r\n")
        non-empty-lines (filter #(not= "" %) lines)]
    (for [s non-empty-lines] (Integer/parseInt s))))

(defn how-many-increases [heights]
  (get (reduce (fn [[sum prev] curr]
                 [(if (and (not= prev nil)
                           (> curr prev))
                    (+ sum 1)
                    sum)
                  curr])
               [0 nil] heights) 0))

(defn- add-tuples [input n]
  (let [in-count (count input)
        out-count (- in-count n -1)]
    (for [i (range out-count)]
      (apply + (subvec input i (+ i n))))))

(let [input (vec (get-input))]
  (print "Part 1:" (str (how-many-increases input)))
  (print "Part 2:" (str (how-many-increases (add-tuples input 3)))))
