(def use-example false)
(def input-file-name (if use-example "day05_example.txt" "day05.txt"))

(defn- get-input []
  ; transform input into a sequence of lines
  ; each line is a pair of points of the form [[x1 y1] [x2 y2]]
  (->> (slurp input-file-name) ; read file
       (re-seq #"(\d+),(\d+) -> (\d+),(\d+)") ; parse each line
       (map rest) ; remove full match, leaving only number groups
       (map #(for [n %] (Integer/parseInt n))) ; parse all integers
       (map (fn [nums] [[(first nums) (second nums)] ; map to pair of points
                        [(nth nums 2) (nth nums 3)]]))))

(defn- line-to-points [line diagonals?]
  (let [[[x1 y1] [x2 y2]] line
        dx (if (pos? (- x2 x1)) 1 -1)
        dy (if (pos? (- y2 y1)) 1 -1)]
    (cond (= x1 x2) (map vector
                         (repeat x1)
                         (range y1 (+ y2 dy) dy))
          (= y1 y2) (map vector
                         (range x1 (+ x2 dx) dx)
                         (repeat y1))
          diagonals? (map vector
                          (range x1 (+ x2 dx) dx)
                          (range y1 (+ y2 dy) dy))
          :else '())))

(defn- add-line-to-vents [line vents diagonals?]
  ; for each point on line, inc vents at point
  (reduce (fn [vents point]
            (if (vents point)
              (update vents point inc)
              (assoc vents point 1)))
          vents ; initial value
          (line-to-points line diagonals?)))

(defn- compute-vents [lines diagonals?]
  ; loop over lines to compute vent positions
  (loop [vents (hash-map) lines lines]
    (if (empty? lines)
      vents ; return vents, if no more lines
      (recur ; else add line to vents and recur with rest
       (add-line-to-vents (first lines) vents diagonals?)
       (rest lines)))))

(defn- count-dangerous-vents [vents]
  (reduce (fn [count vent-value] (if (> vent-value 1) (inc count) count)) 0 (vals vents)))

(let [lines (get-input)
      part1 (count-dangerous-vents (compute-vents lines false))
      part2 (count-dangerous-vents (compute-vents lines true))]
  (println "Part 1:" part1 "Part 2:" part2))
