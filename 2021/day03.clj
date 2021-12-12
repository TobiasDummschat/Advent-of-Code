(require '[clojure.string])

(defn- get-input-lines  []
  (let [file (slurp "day03.txt")
        lines (clojure.string/split file #"\r|\n|\r\n")]
    lines))

(defn- lines-to-vectors [lines]
  (apply vector
         (for [line lines]
           (apply vector
                  (for [char (clojure.string/split line #"")]
                    (Integer/parseInt char 2))))))

(defn- vector-sum [vectors]
  (apply (partial mapv +) vectors))

(defn- binary-vector-to-number [v]
  (Integer/parseInt (clojure.string/join "" v) 2))

(defn gamma-rate [vectors]
  (let [half (/ (count vectors) 2)
        vector_sum (vector-sum vectors)
        gamma-rate-vector (for [entry vector_sum]
                            (if (> entry half) 1 0))]
    (binary-vector-to-number gamma-rate-vector)))

(defn epsilon-rate [gamma-rate number-of-bits]
  (- (Math/pow 2 number-of-bits) 1 gamma-rate))

(defn- get-common-bit [vectors bit common]
  (let [half (/ (count vectors) 2)
        found (apply + (for [v vectors] (get v bit)))
        most-common (cond (> found half) 1 
                          (< found half) 0
                          :else 1 )]
    (case common
      :most most-common
      :least (- 1 most-common))))

(defn- filter-on-bit [vectors bit common]
  (let [wanted (get-common-bit vectors bit common)]
    (filter #(= wanted (get % bit)) vectors)))

(defn- find-rating [vectors rating]
  (let [common (case rating :oxygen :most :co2 :least)]
    (loop [bit 0 vectors vectors]
      (if (= 1 (count vectors))
        (binary-vector-to-number (first vectors))
        (recur (inc bit) (filter-on-bit vectors bit common))))))

(let [lines (get-input-lines)
      vectors (lines-to-vectors lines)
      gamma-rate (gamma-rate vectors)
      epsilon-rate (epsilon-rate gamma-rate (count (first vectors)))
      oxygen-rating (find-rating vectors :oxygen)
      co2-rating (find-rating vectors :co2)]
  (print "Part 1:" (int (* gamma-rate epsilon-rate))
         "Part 2:" (* oxygen-rating co2-rating)))
