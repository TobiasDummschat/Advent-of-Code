(require '[clojure.string])

(defn- get-input  []
  (let [file (slurp ".\\day01.txt")
        str-list (clojure.string/split file #"\r|\n|\r\n")]
    (map #(Integer/parseInt %)
         (filter #(not= "" %) str-list))))

(defn how-many-increases [heights]
  (get (reduce (fn [[sum prev] curr]
            [(if (and (not= prev nil) 
                      (> curr prev)) 
               (+ sum 1) 
               sum) 
             curr])
          [0 nil] heights) 0))

(comment
  (get-input))

(printf (str (how-many-increases (get-input))))

