(require '[clojure.string])

(defn- get-input-lines  []
  (let [file (slurp "day02.txt")
        lines (clojure.string/split file #"\r|\n|\r\n")
        split-lines (for [line lines] (clojure.string/split line #" "))
        parsed-lines (for [[direction amount] split-lines]
                       {:direction direction :distance (Integer/parseInt amount)})]
    parsed-lines))

(defn get-position-part-1 [instructions]
  (reduce (fn [[h-pos depth] instruction]
            (let [d (instruction :distance)]
              (case (instruction :direction)
                "forward" [(+ h-pos d) depth]
                "down" [h-pos (+ depth d)]
                "up" [h-pos (- depth d)])))
          [0 0]
          instructions))

(defn get-position-part-2 [instructions]
  (reduce (fn [[hpos depth aim] instruction]
            (let [d (instruction :distance)]
              (case (instruction :direction)
                "down" [hpos depth (+ aim d)]
                "up" [hpos depth (- aim d)]
                "forward" [(+ hpos d) (+ depth (* aim d)) aim])))
          [0 0 0]
          instructions))

(get-input-lines)
(let [input (get-input-lines)
      pos1 (get-position-part-1 input)
      pos2 (get-position-part-2 input)]
  (print "Part 1:" (* (pos1 0) (pos1 1))
         "Part 2:" (* (pos2 0) (pos2 1))))
