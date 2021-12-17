(require '[clojure.string :as str])

(def day "07")
(def use-example false)
(def input-file-name (str/join ["day" day (if use-example "_example" "") ".txt"]))

(defn- get-input []
  (->> (slurp input-file-name)
       (re-seq #"\d+")
       (map #(Integer. %))))

(defn- median [nums]
  (let [sorted-nums (sort nums)
        length (count nums)]
    (nth sorted-nums (/ (dec length) 2))))

(defn- distance [x y]
  (Math/abs (- x y)))

(defn- fuel [distance part]
  (case part
    1 distance
    2 (int (* distance (inc distance) (/ 2))))) ; gauß sum

(defn- fuel-usage [crabs target part]
  (apply + (for [crab crabs] (fuel (distance crab target) part))))

(defn- mean [nums]
  (float (/ (apply + nums) 
            (count nums))))

(defn- best-pos [crabs part]
  (case part
    1 (median crabs) ; minimizes number of crabs on each side
    2 (let [mean (mean crabs) ; minimizes error (distances) on each side
            floor (Math/floor mean) ; but we have to deal with floats
            fuel-floor (fuel-usage crabs floor 2)
            fuel-ceil (fuel-usage crabs (inc floor) 2)
            ] 
        (if (< fuel-floor fuel-ceil) floor (inc floor))
        )  
    ))

(let [crabs (get-input)
      part1 (fuel-usage crabs (best-pos crabs 1) 1)
      part2 (fuel-usage crabs (best-pos crabs 2) 2)]
  (println "Part 1:" part1 "Part 2:" part2))


