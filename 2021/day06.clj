(require '[clojure.string :as str])

(def day "06")
(def use-example false)
(def input-file-name (str/join ["day" day (if use-example "_example" "") ".txt"]))

(defn- ages-to-fish-by-age
  "converts a list of ages to a map of ages and number of fish that age"
  [ages]
  (reduce
   (fn [fish-by-age age]
     (update fish-by-age age inc))
   (zipmap (range 9) (repeat 0)) ; initial value
   ages))

(defn- get-input []
  (->> (slurp input-file-name)
       (re-seq #"\d+")
       (map #(Integer. %))
       (ages-to-fish-by-age)))

(defn- next-day [fish-by-age]
  {0 (fish-by-age 1)
   1 (fish-by-age 2)
   2 (fish-by-age 3)
   3 (fish-by-age 4)
   4 (fish-by-age 5)
   5 (fish-by-age 6)
   6 (+ (fish-by-age 0) (fish-by-age 7))
   7 (fish-by-age 8)
   8 (fish-by-age 0)})

(defn- n-days [fish-by-age n]
  (if (= n 0) fish-by-age
      (recur (next-day fish-by-age) (dec n))))

(defn- count-vals [some-map] (apply + (vals some-map)))

(let [initial-fish-by-age (get-input)
      part1 (count-vals (n-days initial-fish-by-age 80))
      part2 (count-vals (n-days initial-fish-by-age 256))]
  (println "Part 1:" part1 "Part 2:" part2))
