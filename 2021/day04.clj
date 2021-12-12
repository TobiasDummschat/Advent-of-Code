(require '[clojure.string :as str])

(def use-example false)
(def input-file-name (if use-example "day04_example.txt" "day04.txt"))

(defn- parse-draws [draws-string]
  (-> draws-string
      (str/split #",")
      (#(for [numstr %] (Integer/parseInt numstr)))))

(defn- parse-board [board-string]
  (-> board-string
      ; split board lines
      (str/split #"\r|\r?\n")
      ; split board numbers (and trim possible leading spaces)
      (#(for [line %]
          (str/split (str/trim line) #" +")))
      ; parse to integers
      (#(for [line %]
          (for [numstr line]
            (Integer/parseInt numstr))))))

(defn- get-input []
  (let [file (slurp input-file-name)
        blocks (str/split file #"(\r|\r?\n){2}")
        draws (parse-draws (first blocks))
        boards (for [board (rest blocks)] (parse-board board))]
    [draws boards]))

(defn- transpose [matrix]
  (apply map list matrix))

(defn- check-bingo-rows [board]
  ; check if for any row every entry is x
  (some true? (for [row board]
                (every? #(= :x %) row))))

(defn- check-bingo-cols [board]
  (check-bingo-rows (transpose board)))

(defn- check-bingo [board]
  (or (check-bingo-rows board)
      (check-bingo-cols board)))

(defn- mark-draw [boards draw]
  ; go through boards->rows->cells and mark ones that equal draw
  (for [board boards]
    (for [row board]
      (for [cell row]
        (if (= cell draw) :x cell)))))

(defn- play [boards draws is-done]
  (if (empty? draws)
    "no more draws"
    (-> boards
        (mark-draw (first draws))
        (#(let [result (is-done % draws)]
            (if result
              result
              (play % (rest draws) is-done)))))))

(defn- get-puzzle-answer [board prev-draw]
  (->> board
       (flatten) ; flatten board to list of numbers and x-es
       (filter #(not (= :x %))) ; filter out x-es
       (apply +) ; add remaining numbers
       (* prev-draw) ; multiply by last draw
       ))

(defn- part-1 [boards draws]
  (let [boards-with-bingo (filter check-bingo boards)]
    (if (seq boards-with-bingo)
      (get-puzzle-answer (first boards-with-bingo) (first draws))
      false)))

(defn- finish-board [board draws]
  (let [draw (first draws)
        next-board (first (mark-draw [board] draw))]
    (if (check-bingo next-board)
      [next-board draw]
      (recur next-board (rest draws)))))

(defn- part-2 [boards draws]
  (let [boards-without-bingo (filter #(not (check-bingo %)) boards)]
    ; play until only one board left
    (if (= 1 (count boards-without-bingo))
      ; play last board until win
      (let [last-board (first boards-without-bingo)
            [completed-board last-draw] (finish-board last-board (rest draws))]
        (get-puzzle-answer completed-board last-draw))
      false)))

(let [[draws boards] (get-input)]
  {:part1 (play boards draws part-1)
   :part2 (play boards draws part-2)})

(comment
  (let [[draws boards] (get-input)
        board (first boards)]
    (flatten board)))
