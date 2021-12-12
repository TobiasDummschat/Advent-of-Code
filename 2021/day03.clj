(require '[clojure.string])

(defn- get-input-lines  []
  (let [file (slurp "day03.txt")
        lines (clojure.string/split file #"\r|\n|\r\n")]
    lines))

(defn- lines-to-vectors [lines]
  (for [line lines]
    (for [char (clojure.string/split line #"")]
      (Integer/parseInt char 2))))

(defn- vector-sum [vectors]
  (apply (partial mapv +) vectors))

(defn- binary-vector-to-number [v]
  (Integer/parseInt (clojure.string/join "" v) 2))

(defn gamma-rate [vectors]
  (let [expected (/ (count vectors) 2)
        vector_sum (vector-sum vectors)
        gamma-rate-vector (for [entry vector_sum]
                            (if (> entry expected) 1 0))]
    (binary-vector-to-number gamma-rate-vector)))

(defn epsilon-rate [gamma-rate number-of-bits]
  (- (Math/pow 2 number-of-bits) 1 gamma-rate))

(let [lines (get-input-lines)
      vectors (lines-to-vectors lines)
      gamma-rate (gamma-rate vectors)
      epsilon-rate (epsilon-rate gamma-rate (count (first vectors)))]
  (print "Part 1:" (int (* gamma-rate epsilon-rate))
         "Part 2:" nil)
)

