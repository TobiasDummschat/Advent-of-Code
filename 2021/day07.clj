(require '[clojure.string :as str])

(def day "07")
(def use-example true)
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

(defn- best-pos [crabs part] 
  (case part 
    1 (median crabs)
    2 0 ; TODO
  ))

(let [crabs (get-input)
      part1 (fuel-usage crabs (best-pos crabs 1) 1)
      part2 (fuel-usage crabs (best-pos crabs 2) 2)]
  (println "Part 1:" part1 "Part 2:" part2))


